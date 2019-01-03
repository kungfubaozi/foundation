package authenticate

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/pborman/uuid"
	"time"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/authorize"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error)

	Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error)

	Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error)
}

//只检查用户、状态，以及策略等鉴权问题
type authenticateService struct {
	statecli    state.Service
	usercli     user.Service
	reportercli reporter.Service
	messsagecli messagecli.MessageChannel
	pool        *redis.Pool
}

func (svc *authenticateService) GetRepo() repository {
	return &authenticateRepository{conn: svc.pool.Get()}
}

func (svc *authenticateService) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	panic("implement me")
}

func (svc *authenticateService) sizeCheck() {

}

func (svc *authenticateService) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error) {
	accessTokenClaims, err := authorize.DecodeToken(in.Metadata.Token)
	if err != nil {
		return errno.ErrResponse(errno.ErrToken)
	}
	token := accessTokenClaims.Token
	resp := &fs_base.Response{}
	if !token.Access {
		resp.State = errno.ErrToken
		return resp, nil
	}
	repo := svc.GetRepo()
	defer repo.Close()
	auth, err := repo.Get(token.UserId, token.ClientId, token.UUID)
	if err != nil {
		return errno.ErrResponse(errno.ErrTokenExpired)
	}

	//这里不需要检查在线数量
	errc := make(chan error, 3)

	//检查用户状态
	go func(r *fs_base.Response) {
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
	go func(r *fs_base.Response) {
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
		errc <- err
	}(resp)

	if e := <-errc; e != nil {
		return resp, nil
	}
	//检查是否过期
	if accessTokenClaims.VerifyExpiresAt(time.Now().UnixNano(), true) {
		if in.Metadata.UserAgent == auth.UserAgent && in.Metadata.Platform == auth.Platform &&
			in.Metadata.Device == auth.Device && token.Relation == auth.Relation &&
			token.UUID == auth.AccessTokenUUID {
			resp.State = errno.Ok
			return resp, nil
		}
	}
	resp.State = errno.ErrTokenExpired
	return resp, nil
}

func (svc *authenticateService) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	in.Authorize.AccessTokenUUID = uuid.New()
	in.Authorize.RefreshTokenUUID = uuid.New()
	//关联id，用于关联两个token(accessToken.refreshToken)
	in.Authorize.Relation = uuid.New()
	var accessToken, refreshToken string
	errc := make(chan error, 2)
	go func() {
		a, err := authorize.EncodeAccessToken(in.Authorize)
		accessToken = a
		errc <- err
	}()
	go func() {
		a, err := authorize.EncodeRefreshToken(in.Authorize)
		refreshToken = a
		errc <- err
	}()
	if err := <-errc; err != nil {
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	repo := svc.GetRepo()
	defer repo.Close()
	//检查在线数量
	v, err := repo.SizeOf(in.Authorize.UserId)
	if err != nil {
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	if v != nil && len(v) > 0 {
		//offline
		c := 0
		for _, k := range v {
			key := k.(string)
			if key[0:32] == in.Authorize.ClientId {
				c++
				if c >= int(in.MaxOnlineCount) {
					repo.Del(in.Authorize.UserId, key) //这里不作错误处理
					//send offline message
					svc.messsagecli.SendOffline(&fs_base.DirectMessage{
						To:      key[33:],
						Content: "offline",
					})
					break
				}
			}
		}
	}
	err = repo.Add(in.Authorize)
	if err != nil {
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_authenticate.NewResponse{
		State:        errno.Ok,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		Session:      in.Authorize.Relation,
	}, nil
}

func NewService(statecli state.Service, usercli user.Service, reportercli reporter.Service,
	pool *redis.Pool, messsagecli messagecli.MessageChannel) Service {
	var svc Service
	{
		svc = &authenticateService{statecli: statecli, usercli: usercli, reportercli: reportercli,
			pool: pool, messsagecli: messsagecli}
	}
	return svc
}
