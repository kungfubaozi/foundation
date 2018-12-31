package project

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/strategy"
)

type Service interface {
	New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error)
}

type projectService struct {
	session     *mgo.Session
	strategycli strategy.Service
	reportercli reporter.Service
}
