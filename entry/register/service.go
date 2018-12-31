package register

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/safety/verification"
)

type Service interface {
	FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error)

	FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error)
}

type registerService struct {
	usercli     user.Service
	validatecli validate.Service
}

func (svc *registerService) FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error) {
	resp := verification.FromRequestMeta(svc.validatecli, in.Meta, names.F_DO_REGISTER)
	if !resp.Ok {
		return errno.ErrResponse(resp)
	}
}

//从第三方注册不需要验证码
func (svc *registerService) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {

}

func NewService(usercli user.Service, validatecli validate.Service) Service {
	var svc Service
	{
		svc = &registerService{usercli: usercli, validatecli: validatecli}
	}
	return svc
}
