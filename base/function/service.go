package function

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
)

type Service interface {
	Add(ctx context.Context, in *fs_base_function.UpsertResponse) (*fs_base_function.UpsertResponse, error)

	Remove(ctx context.Context, in *fs_base_function.RemoveRequest) (*fs_base.Response, error)

	Update(ctx context.Context, in *fs_base_function.UpsertResponse) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error)
}

type functionService struct {
	session *mgo.Session
}
