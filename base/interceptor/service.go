package interceptor

import (
	"context"
	"zskparker.com/foundation/base/interceptor/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/transport"
)

type Service interface {
	Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error)
}

//其余的会在拦截器做处理
type interceptorService struct {
	reportercli reportercli.Channel
	messagecli  messagecli.Channel
}

func (svc *interceptorService) Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error) {
	meta := fs_metadata_transport.ContextToMeta(ctx)
	return &fs_base_interceptor.AuthResponse{
		State:  errno.Ok,
		UserId: meta.UserId,
	}, nil
}

func NewService(reportercli reportercli.Channel, messagecli messagecli.Channel) Service {
	var svc Service
	{
		svc = &interceptorService{reportercli: reportercli, messagecli: messagecli}
	}
	return svc
}
