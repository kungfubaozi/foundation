package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/safety/verification"
)

type Service interface {
	//grpc
	Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error)

	//grpc
	Get(ctx context.Context, in *fs_base_user.GetRequest) (*fs_base_user.GetResponse, error)

	//http
	UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	//http
	UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	//http
	UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	//http
	UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)
}

type userService struct {
	session         *mgo.Session
	validatecli     validate.Service
	authenticatecli authenticate.Service
}

func (svc *userService) GetRepo() repository {
	return &userRepository{session: svc.session.Clone()}
}

func (svc *userService) UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	resp := verification.FromRequestMeta(svc.validatecli, in.Meta, names.F_DO_UPDATE_PHONE)
	if !resp.Ok {
		return errno.ErrResponse(resp)
	}

	meta := ctx.Value("meta").(*fs_base.Metadata)

	//offline user
	r, err := svc.authenticatecli.Offline(ctx, &fs_base_authenticate.OfflineRequest{
		UserId: meta.UserId,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !r.State.Ok {
		return errno.ErrResponse(r.State)
	}

	//update
	repo := svc.GetRepo()
	defer repo.Close()

	p, err := bcrypt.GenerateFromPassword([]byte(in.Value), bcrypt.DefaultCost)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	err = repo.UpdatePhone(meta.UserId, string(p))
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	resp := verification.FromRequestMeta(svc.validatecli, in.Meta, names.F_DO_UPDATE_EMAIL)
	if !resp.Ok {
		return errno.ErrResponse(resp)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//offline user
	r, err := svc.authenticatecli.Offline(ctx, &fs_base_authenticate.OfflineRequest{
		UserId: meta.UserId,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !r.State.Ok {
		return errno.ErrResponse(r.State)
	}
	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err = repo.UpdateEmail(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	resp := verification.FromRequestMeta(svc.validatecli, in.Meta, names.F_DO_UPDATE_ENTERPRISE)
	if !resp.Ok {
		return errno.ErrResponse(resp)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//offline user
	r, err := svc.authenticatecli.Offline(ctx, &fs_base_authenticate.OfflineRequest{
		UserId: meta.UserId,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !r.State.Ok {
		return errno.ErrResponse(r.State)
	}
	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err = repo.UpdateEnterprise(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	resp := verification.FromRequestMeta(svc.validatecli, in.Meta, names.F_DO_UPDATE_PASSWORD)
	if !resp.Ok {
		return errno.ErrResponse(resp)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//offline user
	r, err := svc.authenticatecli.Offline(ctx, &fs_base_authenticate.OfflineRequest{
		UserId: meta.UserId,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !r.State.Ok {
		return errno.ErrResponse(r.State)
	}
	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err = repo.UpdatePassword(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error) {

}

func (svc *userService) Get(ctx context.Context, in *fs_base_user.GetRequest) (*fs_base_user.GetResponse, error) {

}

func NewService(session *mgo.Session) Service {
	var svc Service
	{
		svc = &userService{session: session}
	}
	return svc
}
