package invite

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/pb"
)

type Service interface {
	Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error)
}

type inviteService struct {
	session *mgo.Session
}
