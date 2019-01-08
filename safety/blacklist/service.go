package blacklist

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Service interface {
	Check(ctx context.Context, in *fs_safety_blacklist.CheckRequest) (*fs_base.Response, error)

	Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error)
}

type blacklistService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
}

func (svc *blacklistService) Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *blacklistService) Check(ctx context.Context, in *fs_safety_blacklist.CheckRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func NewService(session *mgo.Session, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &blacklistService{session: session, reportercli: reportercli}
	}
	return svc
}
