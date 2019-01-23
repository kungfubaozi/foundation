package userinfo

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/userinfo/pb"
)

type Service interface {
	Upsert(ctx context.Context, in *fs_base_userinfo.UpsertRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_userinfo.GetRequest) (*fs_base_userinfo.GetResponse, error)
}
