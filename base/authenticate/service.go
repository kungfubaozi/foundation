package authenticate

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/go-kit/kit/log"
	"os"
	"sync"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Service interface {
	New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error)

	Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error)

	OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error)

	Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error)
}

//只检查用户、状态，以及策略等鉴权问题
type authenticateService struct {
	statecli     state.Service
	usercli      user.Service
	reportercli  reportercli.Channel
	messsagecli  messagecli.Channel
	pool         *redis.Pool
	logger       log.Logger
	redisync     *fs_redisync.Redisync
	blacklistcli fs_safety_blacklist.BlacklistServer
}

func (svc *authenticateService) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	resp := func(state *fs_base.State) (*fs_base_authenticate.RefreshResponse, error) {
		return &fs_base_authenticate.RefreshResponse{State: state}, nil
	}
	if len(in.RefreshToken) < 32 {
		svc.logger.Log("step")
		return resp(errno.ErrRequest)
	}
	refreshTokenClaims, err := DecodeToken(in.RefreshToken)
	svc.logger.Log(refreshTokenClaims.ExpiresAt)
	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrToken)
	}
	if refreshTokenClaims.Valid() != nil {
		svc.logger.Log("step")
		return resp(errno.ErrTokenExpired)
	}
	token := refreshTokenClaims.Token
	if token.Access {
		svc.logger.Log("step")
		return resp(errno.ErrToken)
	}

	//锁定30s
	if s := svc.redisync.Lock(fs_function_tags.GetAuthenticateRefreshTag(), token.Relation, 30); s.Ok {
		return resp(s)
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)

	repo := svc.GetRepo()
	defer repo.Close()

	auth, err := repo.Get(token.UserId, token.ClientId, token.Relation)
	if err != nil {
		return resp(errno.ErrTokenExpired)
	}

	node := utils.NodeGenerate()
	auth.AccessTokenUUID = node.Generate().Base64()
	var accessToken string

	//!!!验证操作已经在拦截器中设置完成（状态检测，黑名单，用户依存等）

	//可能是跳转的token
	if auth.ClientId != meta.ClientId {
		svc.logger.Log("step")
		//需要通过web端才可跨项目访问
		if auth.Platform == fs_constants.PLATFORM_WEB && meta.Platform == fs_constants.PLATFORM_WEB {
			accessToken, err = EncodeAccessToken(&fs_base_authenticate.Authorize{
				ClientId:        meta.ClientId,
				UserId:          auth.UserId,
				AccessTokenUUID: auth.AccessTokenUUID,
				Relation:        auth.Relation,
			})
			if err != nil {
				svc.logger.Log("err")
				return resp(errno.ErrSystem)
			}

			//由于在缓存里取出的是遵循 userId,clientId,relationId
			//当出现使用refreshToken刷新别的web client时，会新建一个以当前用户的基本信息+client生成token
			//公用一个refreshToken的Relation，
			auth.ClientId = meta.ClientId
		} else { //除了web端其他的客户端不支持token公用
			svc.logger.Log("step")
			return resp(errno.ErrSupport)
		}
	} else {
		svc.logger.Log("step")
		accessToken, err = EncodeAccessToken(auth)
		if err != nil {
			svc.logger.Log("step")
			return resp(errno.ErrSystem)
		}
	}

	err = repo.Add(auth)
	if err != nil {
		return resp(errno.ErrSystem)
	}

	return &fs_base_authenticate.RefreshResponse{
		State:       errno.Ok,
		AccessToken: accessToken,
	}, nil
}

func (svc *authenticateService) OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.DelAll(in.UserId)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *authenticateService) GetRepo() repository {
	return &authenticateRepository{conn: svc.pool.Get()}
}

func (svc *authenticateService) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error) {

	ps := errno.Ok
	var level int64

	resp := func(state *fs_base.State) (*fs_base_authenticate.CheckResponse, error) {
		return &fs_base_authenticate.CheckResponse{State: state}, nil
	}

	accessTokenClaims, err := DecodeToken(in.Metadata.Token)
	if err != nil {
		svc.logger.Log("err", err)
		return resp(errno.ErrToken)
	}
	token := accessTokenClaims.Token
	if !token.Access {
		svc.logger.Log("state", "err")
		return resp(errno.ErrToken)
	}
	//检查过期
	if accessTokenClaims.Valid() != nil {
		svc.logger.Log("state", "err")
		return resp(errno.ErrTokenExpired)
	}
	repo := svc.GetRepo()
	defer repo.Close()
	auth, err := repo.Get(token.UserId, token.ClientId, token.Relation)
	if err != nil {
		svc.logger.Log("state", "err")
		return resp(errno.ErrTokenExpired)
	}

	wg := sync.WaitGroup{}

	//这里不需要检查在线数量
	errc := func(st *fs_base.State) {
		if !st.Ok {
			ps = st
		}
		wg.Done()
	}

	//检查用户状态
	wg.Add(3)
	go func() {
		a, e := svc.statecli.Get(context.Background(), &fs_base_state.GetRequest{
			Key: token.UserId,
		})
		if e != nil {
			errc(errno.ErrSystem)
			svc.logger.Log("err", e)
			return
		}
		if !a.State.Ok {
			errc(a.State)
			svc.logger.Log("state", a.State)
			return
		}
		if a.Status != fs_constants.STATE_OK {
			errc(errno.ErrRequest)
			return
		}
		wg.Done()
	}()

	//检查是否存在用户
	go func() {
		a, e := svc.usercli.FindByUserId(context.Background(), &fs_base_user.FindRequest{
			Value: auth.UserId,
		})
		if e != nil {
			errc(errno.ErrSystem)
			svc.logger.Log("err", e)
			return
		}
		if !a.State.Ok {
			errc(a.State)
			svc.logger.Log("state", a.State)
			return
		}
		level = a.Level
		wg.Done()
	}()

	//检查黑名单
	go func() {
		b, e := svc.blacklistcli.CheckMeta(context.Background(), &fs_safety_blacklist.CheckMetaRequest{
			UserId: in.Metadata.UserId,
		})
		if e != nil {
			errc(errno.ErrSystem)
			svc.logger.Log("err", e)
			return
		}
		if !b.State.Ok {
			errc(b.State)
			svc.logger.Log("state", b.State)
			return
		}
		wg.Done()
	}()

	st := fs_metadata_transport.ContextToStrategy(ctx)
	if st.Events.OnUserEntry.OpenReview == 2 { //是否开始审核
		wg.Add(1)
		go func() {

		}()
	}

	wg.Wait()

	if !ps.Ok {
		return resp(ps)
	}
	if in.Metadata.UserAgent == auth.UserAgent && in.Metadata.Platform == auth.Platform &&
		in.Metadata.Device == auth.Device && token.Relation == auth.Relation &&
		token.UUID == auth.AccessTokenUUID && token.ClientId == in.Metadata.ClientId { //token的ClientId必须和Meta ClientId一致

		//只有web端可以跨项目跳转
		if auth.Platform != fs_constants.PLATFORM_WEB {
			svc.logger.Log("platform", "web")
			if auth.ClientId != in.Metadata.ClientId {
				svc.logger.Log("state", "err")
				return resp(errno.ErrToken)
			}
		}

		svc.logger.Log("state", "ok")

		return &fs_base_authenticate.CheckResponse{
			State:     errno.Ok,
			UserId:    auth.UserId,
			ProjectId: auth.ProjectId,
			ClientId:  auth.ClientId,
			Relation:  auth.Relation,
			Level:     auth.Level,
		}, nil
	}
	svc.logger.Log("check", "over")
	return resp(errno.ErrToken)
}

func (svc *authenticateService) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	resp := func(state *fs_base.State) (*fs_base_authenticate.NewResponse, error) {
		return &fs_base_authenticate.NewResponse{State: state}, nil
	}
	if in.MaxOnlineCount == -1 {
		return resp(errno.ErrProject)
	}
	node := utils.NodeGenerate()
	in.Authorize.AccessTokenUUID = node.Generate().Base64()
	in.Authorize.RefreshTokenUUID = node.Generate().Base64()
	//关联id，用于关联两个token(accessToken.refreshToken)
	in.Authorize.Relation = node.Generate().Base64()
	auth := &auth{}
	var err error
	var cs *fs_base.State

	wg := sync.WaitGroup{}
	errc := func(e *fs_base.State) {
		if cs.Ok {
			cs = e
		}
		wg.Done()
	}
	wg.Add(4)
	go func() {
		a, err := EncodeAccessToken(in.Authorize)
		if err != nil {
			errc(errno.ErrSystem)
			return
		}
		auth.Access = a
		errc(errno.Ok)
	}()
	go func() {
		a, err := EncodeRefreshToken(in.Authorize)
		if err != nil {
			errc(errno.ErrSystem)
			return
		}
		auth.Refresh = a
		errc(errno.Ok)
	}()
	go func() {
		a, e := svc.statecli.Get(context.Background(), &fs_base_state.GetRequest{
			Key: in.Authorize.UserId,
		})
		if e != nil {
			errc(errno.ErrSystem)
			return
		}
		if !a.State.Ok {
			errc(a.State)
			return
		}
		errc(errno.Ok)
	}()
	go func() {
		b, e := svc.blacklistcli.CheckMeta(context.Background(), &fs_safety_blacklist.CheckMetaRequest{
			UserId: in.Authorize.UserId,
		})
		if e != nil {
			errc(errno.ErrSystem)
			svc.logger.Log("err", e)
			return
		}
		if !b.State.Ok {
			errc(b.State)
			svc.logger.Log("state", b.State)
			return
		}
		errc(errno.Ok)
	}()
	wg.Wait()

	if err != nil {
		fmt.Println("encode err", err)
		return resp(errno.ErrSystem)
	}
	if cs != nil {
		return resp(cs)
	}
	repo := svc.GetRepo()
	defer repo.Close()
	//检查在线数量
	v, err := repo.SizeOf(in.Authorize.UserId)
	if err != nil {
		fmt.Println("size of", err)
		return resp(errno.ErrSystem)
	}
	fmt.Println("size ", len(v))
	if len(v) > 0 {
		//offline
		c := 0
		i := len(in.Authorize.ClientId)
		for _, k := range v {
			key := b2s(k.([]uint8))
			if key[0:i] == in.Authorize.ClientId {
				c++
				if c >= int(in.MaxOnlineCount) {
					repo.Del(in.Authorize.UserId, key) //这里不作错误处理
					//send offline message
					svc.messsagecli.SendOffline(&fs_base.DirectMessage{
						To:      key[i+1:],
						Content: "offline",
					})
				}
			}
		}
	}
	err = repo.Add(in.Authorize)
	if err != nil {
		fmt.Println("add", err)
		return resp(errno.ErrSystem)
	}
	return &fs_base_authenticate.NewResponse{
		State:        errno.Ok,
		RefreshToken: auth.Refresh,
		AccessToken:  auth.Access,
		Session:      in.Authorize.Relation,
	}, nil
}

type auth struct {
	Refresh string
	Access  string
}

func b2s(bs []uint8) string {
	var ba []byte
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func NewService(statecli state.Service, usercli user.Service, reportercli reportercli.Channel,
	pool *redis.Pool, messsagecli messagecli.Channel, redisync *fs_redisync.Redisync, blacklistcli blacklist.Service) Service {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var svc Service
	{

		svc = &authenticateService{statecli: statecli, usercli: usercli, reportercli: reportercli,
			pool: pool, messsagecli: messsagecli, logger: logger, redisync: redisync, blacklistcli: blacklistcli}
	}
	return svc
}
