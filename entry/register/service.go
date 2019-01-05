package register

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/errno"
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
	panic(errno.ERROR)
}

//从第三方注册不需要验证码
func (svc *registerService) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {
	panic(errno.ERROR)
}

func NewService(usercli user.Service, validatecli validate.Service) Service {
	var svc Service
	{
		svc = &registerService{usercli: usercli, validatecli: validatecli}
	}
	return svc
}
