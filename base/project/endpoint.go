package project

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
)

type Endpoints struct {
	NewEndpoint endpoint.Endpoint
	GetEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service) Endpoints {
	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = MakeNewEndpoint(svc)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
	}

	return Endpoints{
		NewEndpoint: newEndpoint,
		GetEndpoint: getEndpoint,
	}
}

func (g Endpoints) New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error) {
	resp, err := g.NewEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_project.GetResponse), nil
}

func MakeNewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.New(ctx, request.(*fs_base_project.NewRequest))
	}
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_project.GetRequest))
	}
}
