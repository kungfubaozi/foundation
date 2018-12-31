package state

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state/pb"
)

type Endpoints struct {
	UpsertEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var upsertEndpoint endpoint.Endpoint
	{
		upsertEndpoint = MakeUpsertEndpoint(svc)
		upsertEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(upsertEndpoint)
		upsertEndpoint = zipkin.TraceEndpoint(trace, "Upsert")(upsertEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)
	}

	return Endpoints{
		UpsertEndpoint: upsertEndpoint,
		GetEndpoint:    getEndpoint,
	}

}

func (g *Endpoints) Upsert(ctx context.Context, in *fs_base_state.UpsertRequest) (*fs_base.Response, error) {
	resp, err := g.UpsertEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *Endpoints) Get(ctx context.Context, in *fs_base_state.GetRequest) (*fs_base_state.GetResponse, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_state.GetResponse), nil
}

func MakeUpsertEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Upsert(ctx, request.(*fs_base_state.UpsertRequest))
	}
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_state.GetRequest))
	}
}
