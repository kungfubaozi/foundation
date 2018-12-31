package blacklist

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Service interface {
	Check(ctx context.Context, in *fs_safety_blacklist.CheckRequest) (*fs_base.Response, error)

	Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error)
}

type blacklistService struct {
	session     *mgo.Session
	reportercli reporter.Service
}
