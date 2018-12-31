package unblock

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/safety/unblock/pb"
)

type Service interface {
	Unlock(ctx context.Context, in *fs_safety_unblock.UnlockRequest) (*fs_base.Response, error)
}

type unblockService struct {
	session     *mgo.Session
	usercli     user.Service
	reportercli reporter.Service
}
