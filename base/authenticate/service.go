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
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error)

	Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error)

	Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error)

	OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error)
}

//只检查用户、状态，以及策略等鉴权问题
type authenticateService struct {
	statecli    state.Service
	usercli     user.Service
	reportercli reportercli.Channel
	messsagecli messagecli.Channel
	pool        *redis.Pool
	logger      log.Logger
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

//获取新的AccessToken
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
	repo := svc.GetRepo()
	defer repo.Close()

	meta := fs_metadata_transport.ContextToMeta(ctx)

	auth, err := repo.Get(token.UserId, token.ClientId, token.Relation)
	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrTokenExpired)
	}
	node := utils.NodeGenerate()
	auth.AccessTokenUUID = node.Generate().Base64()
	var accessToken string
	var wr *fs_base.State

	wg := sync.WaitGroup{}

	errc := func(e error) {
		if err == nil {
			err = e
		}
		wg.Done()
	}

	//检查用户状态
	wg.Add(1)
	go func() {
		a, e := svc.statecli.Get(context.Background(), &fs_base_state.GetRequest{
			Key: auth.UserId,
		})
		if e != nil {
			errc(e)
			return
		}
		if !a.State.Ok {
			wr = a.State
			errc(errno.ERROR)
			return
		}
		wg.Done()
	}()

	//检查是否存在用户
	wg.Add(1)
	go func() {
		a, e := svc.usercli.FindByUserId(context.Background(), &fs_base_user.FindRequest{
			Value: auth.UserId,
		})
		if e != nil {
			errc(e)
			return
		}
		if !a.State.Ok {
			wr = a.State
			errc(errno.ERROR)
			return
		}
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrSystem)
	}

	if wr != nil {
		svc.logger.Log("step")
		return resp(wr)
	}

	//可能是跳转的token
	if auth.ClientId != meta.ClientId {
		svc.logger.Log("step")
		//需要通过web端才可跨项目访问
		if auth.Platform == fs_constants.PLATFORM_WEB && meta.Platform == fs_constants.PLATFORM_WEB {
			accessToken, err = encodeAccessToken(&fs_base_authenticate.Authorize{
				ClientId:        meta.ClientId,
				UserId:          auth.UserId,
				AccessTokenUUID: auth.AccessTokenUUID,
				Relation:        auth.Relation,
			})
			if err != nil {
				svc.logger.Log("err")
				return resp(errno.ErrSystem)
			}
		} else {
			svc.logger.Log("step")
			return resp(errno.ErrSupport)
		}
	} else {
		svc.logger.Log("step")
		accessToken, err = encodeAccessToken(auth)
		if err != nil {
			svc.logger.Log("step")
			return resp(errno.ErrSystem)
		}
	}

	//覆盖原有的token
	err = repo.Add(auth)
	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrSystem)
	}

	return &fs_base_authenticate.RefreshResponse{
		State:       errno.Ok,
		AccessToken: accessToken,
	}, nil

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
	wg.Add(1)
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
		wg.Done()
	}()

	//检查是否存在用户
	wg.Add(1)
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
	errc := func(e error) {
		if err == nil {
			err = e
		}
		wg.Done()
	}
	wg.Add(1)
	go func() {
		a, err := encodeAccessToken(in.Authorize)
		auth.Access = a
		errc(err)
	}()
	wg.Add(1)
	go func() {
		a, err := encodeRefreshToken(in.Authorize)
		auth.Refresh = a
		errc(err)
	}()
	wg.Add(1)
	go func() {
		a, e := svc.statecli.Get(context.Background(), &fs_base_state.GetRequest{
			Key: in.Authorize.UserId,
		})
		if e != nil {
			errc(e)
			return
		}
		if !a.State.Ok {
			cs = a.State
			errc(errno.ERROR)
			return
		}
		wg.Done()
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
	pool *redis.Pool, messsagecli messagecli.Channel) Service {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	var svc Service
	{

		svc = &authenticateService{statecli: statecli, usercli: usercli, reportercli: reportercli,
			pool: pool, messsagecli: messsagecli, logger: logger}
	}
	return svc
}
