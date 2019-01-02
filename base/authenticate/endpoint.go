package authenticate

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project"
)

type Endpoints struct {
	OfflineEndpoint endpoint.Endpoint
	NewEndpoint     endpoint.Endpoint
	CheckEndpoint   endpoint.Endpoint
	RefreshEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, functioncli function.Service,
	projectcli project.Service) Endpoints {

	var offlineEndpoint endpoint.Endpoint
	{
		offlineEndpoint = MakeOfflineEndpoint(svc)
		offlineEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(offlineEndpoint)
		offlineEndpoint = zipkin.TraceEndpoint(trace, "Offline")(offlineEndpoint)
	}

	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = MakeNewEndpoint(svc)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(newEndpoint)
		newEndpoint = zipkin.TraceEndpoint(trace, "New")(newEndpoint)
	}

	var checkEndpoint endpoint.Endpoint
	{
		checkEndpoint = MakeCheckEndpoint(svc)
		checkEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(checkEndpoint)
		checkEndpoint = zipkin.TraceEndpoint(trace, "Check")(checkEndpoint)

	}

	return Endpoints{
		OfflineEndpoint: offlineEndpoint,
		NewEndpoint:     newEndpoint,
		CheckEndpoint:   checkEndpoint,
	}

}

func MakeOfflineEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Offline(ctx, request.(*fs_base_authenticate.OfflineRequest))
	}
}

func MakeNewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.New(ctx, request.(*fs_base_authenticate.Authorize))
	}
}

func MakeCheckEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Check(ctx, request.(*fs_base_authenticate.CheckRequest))
	}
}

func (g Endpoints) Offline(ctx context.Context, in *fs_base_authenticate.OfflineRequest) (*fs_base.Response, error) {
	resp, err := g.OfflineEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) New(ctx context.Context, in *fs_base_authenticate.Authorize) (*fs_base_authenticate.NewResponse, error) {
	resp, err := g.NewEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.NewResponse), nil
}

func (g Endpoints) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error) {
	resp, err := g.CheckEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}
