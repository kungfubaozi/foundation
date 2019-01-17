package blacklist

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/middlewares"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Endpoints struct {
	CheckAccountEndpoint endpoint.Endpoint
	CheckMetaEndpoint    endpoint.Endpoint
	AddEndpoint          endpoint.Endpoint
}

func (g Endpoints) CheckMeta(ctx context.Context, in *fs_safety_blacklist.CheckMetaRequest) (*fs_base.Response, error) {
	resp, err := g.CheckMetaEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) CheckAccount(ctx context.Context, in *fs_safety_blacklist.CheckAccountRequest) (*fs_base.Response, error) {
	resp, err := g.CheckAccountEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error) {
	resp, err := g.AddEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client fs_endpoint_middlewares.Endpoint) Endpoints {

	var checkAccountEndpoint endpoint.Endpoint
	{
		checkAccountEndpoint = MakeCheckAccountEndpoint(svc)
		checkAccountEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(checkAccountEndpoint)
		checkAccountEndpoint = zipkin.TraceEndpoint(trace, "CheckAccount")(checkAccountEndpoint)
	}

	var checkMetaEndpoint endpoint.Endpoint
	{
		checkMetaEndpoint = MakeCheckMetaEndpoint(svc)
		checkMetaEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(checkMetaEndpoint)
		checkMetaEndpoint = zipkin.TraceEndpoint(trace, "CheckMeta")(checkMetaEndpoint)
	}

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = MakeAddEndpoint(svc)
		addEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addEndpoint)
		addEndpoint = zipkin.TraceEndpoint(trace, "Add")(addEndpoint)

		addEndpoint = client.WithMeta()(addEndpoint)
	}

	return Endpoints{
		CheckAccountEndpoint: checkAccountEndpoint,
		CheckMetaEndpoint:    checkMetaEndpoint,
		AddEndpoint:          addEndpoint,
	}

}

func MakeCheckAccountEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.CheckAccount(ctx, request.(*fs_safety_blacklist.CheckAccountRequest))
	}
}

func MakeCheckMetaEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.CheckMeta(ctx, request.(*fs_safety_blacklist.CheckMetaRequest))
	}
}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Add(ctx, request.(*fs_safety_blacklist.AddRequest))
	}
}
