package update

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/update/pb"
)

//http
type Service interface {
	UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error)

	ResetPassword(ctx context.Context, in *fs_safety_update.ResetPasswordRequest) (*fs_base.Response, error)
}

type updateService struct {
	usercli user.Service
}

func (svc *updateService) ResetPassword(ctx context.Context, in *fs_safety_update.ResetPasswordRequest) (*fs_base.Response, error) {
	if len(in.Account) == 0 || len(in.New) < 6 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	var u *fs_base_user.FindResponse
	var err error

	//检查对应的验证账号是否是当前传过来的
	if s := fs_metadata_transport.CheckValidateAccount(ctx, in.Account); !s.Ok {
		return errno.ErrResponse(s)
	}

	if fs_regx_match.Phone(in.Account) {
		u, err = svc.usercli.FindByPhone(context.Background(), &fs_base_user.FindRequest{
			Value: in.Account,
		})
	} else if fs_regx_match.Email(in.Account) {
		u, err = svc.usercli.FindByEmail(context.Background(), &fs_base_user.FindRequest{
			Value: in.Account,
		})
	} else {
		return errno.ErrResponse(errno.ErrRequest)
	}

	if err != nil || u == nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !u.State.Ok {
		return errno.ErrResponse(u.State)
	}

	return svc.update(ctx, &fs_safety_update.UpdateRequest{Value: in.New}, u.UserId, 3)
}

func (svc *updateService) update(ctx context.Context, in *fs_safety_update.UpdateRequest, userId string, c int64) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	var err error
	resp := &fs_base.Response{State: errno.Ok}
	switch c {
	case 1:
		resp, err = svc.usercli.UpdatePhone(ctx, &fs_base_user.UpdateRequest{
			Value:  in.Value,
			UserId: userId,
		})
		break
	case 2:
		resp, err = svc.usercli.UpdateEmail(ctx, &fs_base_user.UpdateRequest{
			Value:  in.Value,
			UserId: userId,
		})
		break
	case 3:
		resp, err = svc.usercli.UpdatePassword(ctx, &fs_base_user.UpdateRequest{
			Value:  in.Value,
			UserId: userId,
		})
		break
	}
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return resp, nil
}

func (svc *updateService) UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	meta := fs_metadata_transport.ContextToMeta(ctx)
	return svc.update(ctx, in, meta.UserId, 1)
}

func (svc *updateService) UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	panic(errno.ERROR)
}

func (svc *updateService) UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	meta := fs_metadata_transport.ContextToMeta(ctx)
	return svc.update(ctx, in, meta.UserId, 2)
}

func (svc *updateService) UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	meta := fs_metadata_transport.ContextToMeta(ctx)
	return svc.update(ctx, in, meta.UserId, 3)
}

func NewService(userlci user.Service) Service {
	var svc Service
	{
		svc = &updateService{usercli: userlci}
	}
	return svc
}
