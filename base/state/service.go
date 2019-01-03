package state

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
)

type Service interface {
	Upsert(ctx context.Context, in *fs_base_state.UpsertRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_state.GetRequest) (*fs_base_state.GetResponse, error)
}

type stateService struct {
	pool    *redis.Pool
	session *mgo.Session
}

func (svc *stateService) GetRepo() repository {
	return &stateRepository{conn: svc.pool.Get(), session: svc.session.Clone()}
}

func (svc *stateService) Upsert(ctx context.Context, in *fs_base_state.UpsertRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	e := repo.Upset(in.Key, in.Status)
	if e != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *stateService) Get(ctx context.Context, in *fs_base_state.GetRequest) (*fs_base_state.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	resp := &fs_base_state.GetResponse{}
	i, e := repo.Get(in.Key)
	if e != nil {
		return &fs_base_state.GetResponse{State: errno.ErrSystem}, nil
	}
	if i == names.F_USER_STATE_FROZE {
		resp.State = errno.ErrFroze
		return resp, nil
	}
	return &fs_base_state.GetResponse{
		State:  resp.State,
		Status: i,
	}, nil
}

func NewService(pool *redis.Pool, session *mgo.Session) Service {
	var svc Service
	{
		svc = &stateService{pool: pool, session: session}
	}
	return svc
}
