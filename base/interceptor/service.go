package interceptor

import (
	"context"
	"zskparker.com/foundation/base/interceptor/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/validate"
)

type Service interface {
	Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error)
}

//其余的会在拦截器做处理
type interceptorService struct {
	validatecli validate.Service
}

func (svc *interceptorService) Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error) {
	meta := ctx.Value("meta").(*fs_base.Metadata)
}

func NewService(validatecli validate.Service) Service {
	var svc Service
	{
		svc = &interceptorService{validatecli: validatecli}
	}
	return svc
}
