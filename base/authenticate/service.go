package authenticate

import (
	"context"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/user"
)

type Service interface {
	Offline(ctx context.Context, in *fs_base_authenticate.OfflineRequest) (*fs_base.Response, error)

	New(ctx context.Context, in *fs_base_authenticate.Authorize) (*fs_base_authenticate.NewResponse, error)

	Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error)
}

//只检查用户、状态，以及策略等鉴权问题
type authenticateService struct {
	statecli state.Service
	usercli  user.Service
}
