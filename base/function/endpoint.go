package function

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/middlewares"
)

type Endpoints struct {
	GetEndpoint  endpoint.Endpoint
	InitEndpoint endpoint.Endpoint
}

func (g Endpoints) Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_function.GetResponse), nil
}

func (g Endpoints) Init(ctx context.Context, in *fs_base_function.InitRequest) (*fs_base.Response, error) {
	resp, err := g.InitEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, clients fs_endpoint_middlewares.Endpoint) Endpoints {

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)

		getEndpoint = clients.WithMeta()(getEndpoint)
	}

	var initEndpoint endpoint.Endpoint
	{
		initEndpoint = MakeInitEndpoint(svc)
		initEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(initEndpoint)
		initEndpoint = zipkin.TraceEndpoint(trace, "Init")(initEndpoint)
	}

	return Endpoints{
		GetEndpoint:  getEndpoint,
		InitEndpoint: initEndpoint,
	}

}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_function.GetRequest))
	}
}

func MakeInitEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Init(ctx, request.(*fs_base_function.InitRequest))
	}
}
