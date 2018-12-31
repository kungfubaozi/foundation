package interceptor

import (
	"context"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/interceptor/pb"
	"zskparker.com/foundation/base/validate"
)

type Service interface {
	Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error)
}

type interceptorService struct {
	validatecli validate.Service
	functioncli function.Service //功能检查
}
