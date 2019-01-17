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
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/middlewares"
)

type Endpoints struct {
	NewEndpoint         endpoint.Endpoint
	CheckEndpoint       endpoint.Endpoint
	OfflineUserEndpoint endpoint.Endpoint
	RefreshEndpoint     endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client fs_endpoint_middlewares.Endpoint) Endpoints {

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

	var offlineUserEndpoint endpoint.Endpoint
	{
		offlineUserEndpoint = MakeOfflineUserEndpoint(svc)
		offlineUserEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(offlineUserEndpoint)
		offlineUserEndpoint = zipkin.TraceEndpoint(trace, "OfflineUser")(offlineUserEndpoint)

	}

	var refreshEndpoint endpoint.Endpoint
	{
		refreshEndpoint = MakeRefreshEndpoint(svc)
		refreshEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(refreshEndpoint)
		refreshEndpoint = zipkin.TraceEndpoint(trace, "Refresh")(refreshEndpoint)

		refreshEndpoint = client.WithMeta()(refreshEndpoint)
	}

	return Endpoints{
		NewEndpoint:         newEndpoint,
		CheckEndpoint:       checkEndpoint,
		OfflineUserEndpoint: offlineUserEndpoint,
		RefreshEndpoint:     refreshEndpoint,
	}

}

func MakeRefreshEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Refresh(ctx, request.(*fs_base_authenticate.RefreshRequest))
	}
}

func MakeOfflineUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.OfflineUser(ctx, request.(*fs_base_authenticate.OfflineUserRequest))
	}
}

func MakeNewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.New(ctx, request.(*fs_base_authenticate.NewRequest))
	}
}

func MakeCheckEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Check(ctx, request.(*fs_base_authenticate.CheckRequest))
	}
}

func (g Endpoints) OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error) {
	resp, err := g.OfflineUserEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	resp, err := g.NewEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.NewResponse), nil
}

func (g Endpoints) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error) {
	resp, err := g.CheckEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.CheckResponse), nil
}

func (g Endpoints) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	resp, err := g.RefreshEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.RefreshResponse), nil
}
