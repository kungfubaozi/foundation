package froze

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/safety/froze/pb"
)

type Service interface {
	Check(ctx context.Context, in *fs_safety_froze.CheckRequest) (*fs_base.Response, error)

	Add(ctx context.Context, in *fs_safety_froze.AddRequest) (*fs_base.Response, error)

	Remove(ctx context.Context, in *fs_safety_froze.RemoveRequest) (*fs_base.Response, error)
}

type frozeService struct {
	session     *mgo.Session
	reportercli reporter.Service
}
