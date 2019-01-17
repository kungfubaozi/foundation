package update

import (
	"context"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/pkg/errno"
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

func (svc *updateService) update(ctx context.Context, in *fs_safety_update.UpdateRequest, c int64) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	var err error
	resp := &fs_base.Response{State: errno.Ok}
	switch c {
	case 1:
		resp, err = svc.usercli.UpdatePhone(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case 2:
		resp, err = svc.usercli.UpdateEmail(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case 3:
		resp, err = svc.usercli.UpdatePassword(ctx, &fs_base_user.UpdateRequest{
			Value: in.Value,
		})
		break
	case 4:
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
	return svc.update(ctx, in, 1)
}

func (svc *updateService) UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, 4)
}

func (svc *updateService) UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, 2)
}

func (svc *updateService) UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	return svc.update(ctx, in, 3)
}

func NewService(userlci user.Service, validatecli validate.Service) Service {
	var svc Service
	{
		svc = &updateService{usercli: userlci, validatecli: validatecli}
	}
	return svc
}
