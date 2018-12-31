package strategy

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/strategy/pb"
)

type Service interface {
	Get(ctx context.Context, in *fs_base_strategy.GetRequest) (*fs_base_strategy.GetResponse, error)

	Upsert(ctx context.Context, in *fs_base_strategy.UpsertRequest) (*fs_base.Response, error)
}

type strategyService struct {
	session *mgo.Session
}
