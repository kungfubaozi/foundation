package authenticate

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/pkg/authorize"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	Offline(ctx context.Context, in *fs_base_authenticate.OfflineRequest) (*fs_base.Response, error)

	New(ctx context.Context, in *fs_base_authenticate.Authorize) (*fs_base_authenticate.NewResponse, error)

	Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error)

	Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error)
}

//只检查用户、状态，以及策略等鉴权问题
type authenticateService struct {
	statecli    state.Service
	usercli     user.Service
	reportercli reporter.Service
	functioncli function.Service
	pool        *redis.Pool
}

func (svc *authenticateService) GetRepo() repository {
	return &authenticateRepository{conn: svc.pool.Get()}
}

func (svc *authenticateService) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	panic("implement me")
}

func (svc *authenticateService) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *authenticateService) New(ctx context.Context, in *fs_base_authenticate.Authorize) (*fs_base_authenticate.NewResponse, error) {
	var accessToken, refreshToken string
	errc := make(chan error, 2)
	go func() {
		a, err := authorize.EncodeAccessToken(in)
		accessToken = a
		errc <- err
	}()
	go func() {
		a, err := authorize.EncodeRefreshToken(in)
		refreshToken = a
		errc <- err
	}()
	if err := <-errc; err != nil {
		return &fs_base_authenticate.NewResponse{State: errno.ErrSystem}, nil
	}
	repo := svc.GetRepo()
	defer repo.Close()

}

func (svc *authenticateService) Offline(ctx context.Context, in *fs_base_authenticate.OfflineRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func NewService(statecli state.Service, usercli user.Service, reportercli reporter.Service, functioncli function.Service, pool *redis.Pool) Service {
	var svc Service
	{
		svc = &authenticateService{statecli: statecli, usercli: usercli, reportercli: reportercli,
			functioncli: functioncli, pool: pool}
	}
	return svc
}
