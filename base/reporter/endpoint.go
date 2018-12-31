package reporter

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/pb"
)

type Endpoints struct {
	WriteEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service) Endpoints {

	return Endpoints{
		WriteEndpoint: MakeWriteEndpoint(svc),
	}

}

func (g Endpoints) Write(ctx context.Context, in *fs_base_reporter.WriteRequest) (*fs_base.Response, error) {
	resp, err := g.WriteEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeWriteEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Write(ctx, request.(*fs_base_reporter.WriteRequest))
	}
}
