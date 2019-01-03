package update

import (
	"context"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/safety/update/pb"
)

//http
type Service interface {
	UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)
}

type updateService struct {
	usercli         user.Service
	validatecli     validate.Service
	authenticatecli authenticate.Service
}

func (svc *updateService) update(ctx context.Context, in *fs_safety_update.UpdateRequest, c string) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	var err error
	resp := &fs_base.Response{State: errno.Ok}
	switch c {
	case names.F_FUNC_UPDATE_PHONE:
		resp, err = svc.usercli.UpdatePhone(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case names.F_FUNC_UPDATE_EMAIL:
		resp, err = svc.usercli.UpdateEmail(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case names.F_FUNC_UPDATE_PASSWORD:
		resp, err = svc.usercli.UpdatePassword(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case names.F_FUNC_UPDATE_ENTERPRISE:
		resp, err = svc.usercli.UpdateEnterprise(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	}
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return resp, nil
}

func (svc *updateService) UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, names.F_FUNC_UPDATE_PHONE)
}

func (svc *updateService) UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, names.F_FUNC_UPDATE_ENTERPRISE)
}

func (svc *updateService) UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, names.F_FUNC_UPDATE_EMAIL)
}

func (svc *updateService) UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, names.F_FUNC_UPDATE_PASSWORD)
}

func NewService(userlci user.Service, validatecli validate.Service) Service {
	var svc Service
	{
		svc = &updateService{usercli: userlci, validatecli: validatecli}
	}
	return svc
}
