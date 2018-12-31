package password

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/safety/password/pb"
)

type Service interface {
	//重置密码
	Reset(ctx context.Context, in *fs_safety_password.ResetRequest) (*fs_base.Response, error)
}

type passwordService struct {
	userlci     user.Service
	reportercli reporter.Service
}
