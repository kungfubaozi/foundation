package authorization

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/authorization/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/userinfo"
)

type Service interface {
	Sync(ctx context.Context, in *fs_base_authorization.SyncRequest) (*fs_base.Response, error)

	Check(ctx context.Context, in *fs_base_authorization.SyncRequest) (*fs_base.Response, error)
}

type authorizationService struct {
	session     *mgo.Session
	usercli     user.Service
	userinfocli userinfo.Service
}

func (svc *authorizationService) Sync(ctx context.Context, in *fs_base_authorization.SyncRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *authorizationService) Check(ctx context.Context, in *fs_base_authorization.SyncRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func NewService(session *mgo.Session, usercli user.Service) Service {
	var svc Service
	{
		svc = &authorizationService{}
	}
	return svc
}
