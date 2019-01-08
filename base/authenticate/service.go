package authenticate

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/errno"
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

func (svc *authenticateService) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	panic("implement me")
}

func (svc *authenticateService) sizeCheck() {

}

func (svc *authenticateService) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error) {
	resp := &fs_base_authenticate.CheckResponse{}
	accessTokenClaims, err := decodeToken(in.Metadata.Token)
	if err != nil {
		resp.State = errno.ErrToken
		return resp, nil
	}
	token := accessTokenClaims.Token
	if !token.Access {
		resp.State = errno.ErrToken
		return resp, nil
	}
	repo := svc.GetRepo()
	defer repo.Close()
	auth, err := repo.Get(token.UserId, token.ClientId, token.UUID)
	if err != nil {
		resp.State = errno.ErrTokenExpired
		return resp, nil
	}

	//这里不需要检查在线数量
	errc := make(chan error, 2)

	//检查用户状态
	go func(r *fs_base_authenticate.CheckResponse) {
		a, e := svc.statecli.Get(context.Background(), &fs_base_state.GetRequest{})
		if e != nil {
			errc <- e
		}
		if !a.State.Ok {
			r.State = a.State
			errc <- errno.ERROR
			return
		}
		errc <- nil
	}(resp)

	//检查是否存在用户
	go func(r *fs_base_authenticate.CheckResponse) {
		a, e := svc.usercli.FindByUserId(context.Background(), &fs_base_user.FindRequest{
			Value: auth.UserId,
		})
		if e != nil {
			errc <- e
		}
		if !a.State.Ok {
			r.State = a.State
			errc <- errno.ERROR
			return
		}
		//if !in.AllowOtherProjectUserToLogin && a.FromProjectId != auth.ProjectId {
		//	r.State = errno.ErrProjectAccess
		//	errc <- errno.ERROR
		//}
		resp.Level = a.Level
		errc <- err
	}(resp)

	if e := <-errc; e != nil {
		return resp, nil
	}
	//检查是否过期
	if accessTokenClaims.VerifyExpiresAt(time.Now().UnixNano(), true) {
		if in.Metadata.UserAgent == auth.UserAgent && in.Metadata.Platform == auth.Platform &&
			in.Metadata.Device == auth.Device && token.Relation == auth.Relation &&
			token.UUID == auth.AccessTokenUUID && token.ClientId == auth.ClientId {
			resp.State = errno.Ok

			resp.UserId = auth.UserId
			resp.ProjectId = auth.ProjectId
			resp.ClientId = auth.ClientId
			resp.Relation = auth.Relation
			resp.Level = auth.Level

			return resp, nil
		}
	}
	resp.State = errno.ErrTokenExpired
	return resp, nil
}

func (svc *authenticateService) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	node := utils.NodeGenerate()
	in.Authorize.AccessTokenUUID = node.Generate().Base64()
	in.Authorize.RefreshTokenUUID = node.Generate().Base64()
	//关联id，用于关联两个token(accessToken.refreshToken)
	in.Authorize.Relation = node.Generate().Base64()
	var accessToken, refreshToken string
	var err error
	wg := sync.WaitGroup{}
	errc := func(e error) {
		if err == nil {
			err = e
		}
	}
	wg.Add(1)
	go func() {
		a, err := encodeAccessToken(in.Authorize)
		accessToken = a
		errc(err)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		a, err := encodeRefreshToken(in.Authorize)
		refreshToken = a
		errc(err)
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		fmt.Println("encode err", err)
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	repo := svc.GetRepo()
	defer repo.Close()
	//检查在线数量
	v, err := repo.SizeOf(in.Authorize.UserId)
	if err != nil {
		fmt.Println("size of", err)
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	fmt.Println("size ", len(v))
	if v != nil && len(v) > 0 {
		//offline
		c := 0
		i := len(in.Authorize.UserId)
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
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_authenticate.NewResponse{
		State:        errno.Ok,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Session:      in.Authorize.Relation,
	}, nil
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
	var svc Service
	{
		svc = &authenticateService{statecli: statecli, usercli: usercli, reportercli: reportercli,
			pool: pool, messsagecli: messsagecli}
	}
	return svc
}
