package function

import (
	"context"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	Add(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base_function.UpsertResponse, error)

	Remove(ctx context.Context, in *fs_base_function.RemoveRequest) (*fs_base.Response, error)

	Update(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error)
}

type functionService struct {
	session *mgo.Session
}

func (svc *functionService) Add(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base_function.UpsertResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	f, err := repo.Get(in.Api)
	if err == mgo.ErrNotFound {
		err = nil
	}
	if err != nil {
		return &fs_base_function.UpsertResponse{
			State: errno.ErrSystem,
		}, nil
	}
	if f != nil {
		return &fs_base_function.UpsertResponse{
			State: errno.ErrRequest,
		}, nil
	}
	f = &function{
		Func:         uuid.NewV4().String()[24:],
		ZH:           in.Zh,
		API:          in.Api,
		Level:        in.Level,
		EN:           in.En,
		Fcv:          in.Fcv,
		Verification: in.Verification,
	}
	err = repo.Add(f)
	if err != nil {
		return &fs_base_function.UpsertResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_function.UpsertResponse{
		State: errno.Ok,
		Func:  f.Func,
	}, nil
}

func (svc *functionService) Remove(ctx context.Context, in *fs_base_function.RemoveRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.Remove(in.Func, "func")
	if err == mgo.ErrNotFound {
		return errno.ErrResponse(errno.ErrInvalid)
	}
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *functionService) Update(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *functionService) Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	var function *function
	var err error
	if len(in.Func) == 0 {
		function, err = repo.FindByFunc(in.Func)
	} else {
		function, err = repo.Get(in.Api)
	}
	if err != nil {
		return &fs_base_function.GetResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_function.GetResponse{
		State: errno.Ok,
		Func: &fs_base_function.Func{
			Zh:           function.ZH,
			Api:          function.API,
			En:           function.EN,
			Fcv:          function.Fcv,
			Verification: function.Verification,
			Func:         function.Func,
		},
	}, nil
}

func NewService(session *mgo.Session) Service {
	var svc Service
	{
		svc = &functionService{session: session}
	}
	return svc
}

func (svc *functionService) GetRepo() repository {
	return &functionRepository{session: svc.session.Clone()}
}
