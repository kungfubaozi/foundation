package strategy

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/middlewares"
)

type Endpoints struct {
	GetEndpoint    endpoint.Endpoint
	UpsertEndpoint endpoint.Endpoint
	InitEndpoint   endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client fs_endpoint_middlewares.Endpoint) Endpoints {

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)
	}

	var upsertEndpoint endpoint.Endpoint
	{
		upsertEndpoint = MakeUpsertEndpoint(svc)
		upsertEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(upsertEndpoint)
		upsertEndpoint = zipkin.TraceEndpoint(trace, "Upsert")(upsertEndpoint)

		upsertEndpoint = client.WithMeta()(upsertEndpoint)
	}

	var initEndpoint endpoint.Endpoint
	{
		initEndpoint = MakeInitEndpoint(svc)
		initEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(initEndpoint)
		initEndpoint = zipkin.TraceEndpoint(trace, "Init")(initEndpoint)

	}

	return Endpoints{
		GetEndpoint:    getEndpoint,
		UpsertEndpoint: upsertEndpoint,
		InitEndpoint:   initEndpoint,
	}

}

func (g Endpoints) Init(ctx context.Context, in *fs_base_strategy.InitRequest) (*fs_base.Response, error) {
	resp, err := g.InitEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) Get(ctx context.Context, in *fs_base_strategy.GetRequest) (*fs_base_strategy.GetResponse, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_strategy.GetResponse), nil
}

func (g Endpoints) Upsert(ctx context.Context, in *fs_base_strategy.UpsertRequest) (*fs_base.Response, error) {
	resp, err := g.UpsertEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_strategy.GetRequest))
	}
}

func MakeInitEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Init(ctx, request.(*fs_base_strategy.InitRequest))
	}
}

func MakeUpsertEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Upsert(ctx, request.(*fs_base_strategy.UpsertRequest))
	}
}
