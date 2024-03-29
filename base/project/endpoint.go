package project

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/pkg/middlewares"
)

type Endpoints struct {
	NewEndpoint            endpoint.Endpoint
	GetEndpoint            endpoint.Endpoint
	EnablePlatformEndpoint endpoint.Endpoint
	InitEndpoint           endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client fs_endpoint_middlewares.Endpoint) Endpoints {
	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = MakeNewEndpoint(svc)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(newEndpoint)
		newEndpoint = zipkin.TraceEndpoint(trace, "New")(newEndpoint)

		newEndpoint = client.WithMeta()(newEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)
	}

	var enablePlatform endpoint.Endpoint
	{
		enablePlatform = MakeEnablePlatformEndpoint(svc)
		enablePlatform = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(enablePlatform)
		enablePlatform = zipkin.TraceEndpoint(trace, "EnablePlatform")(enablePlatform)

		enablePlatform = client.WithMeta()(enablePlatform)
	}

	var initEndpoint endpoint.Endpoint
	{
		initEndpoint = MakeInitEndpoint(svc)
		initEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(initEndpoint)
		initEndpoint = zipkin.TraceEndpoint(trace, "Init")(initEndpoint)
	}

	return Endpoints{
		NewEndpoint:            newEndpoint,
		GetEndpoint:            getEndpoint,
		EnablePlatformEndpoint: enablePlatform,
		InitEndpoint:           initEndpoint,
	}
}

func (g Endpoints) Init(ctx context.Context, in *fs_base_project.InitRequest) (*fs_base_project.InitResponse, error) {
	resp, err := g.InitEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_project.InitResponse), nil
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

func (g Endpoints) EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeInitEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Init(ctx, request.(*fs_base_project.InitRequest))
	}
}

func MakeNewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.New(ctx, request.(*fs_base_project.NewRequest))
	}
}

func MakeEnablePlatformEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EnablePlatform(ctx, request.(*fs_base_project.EnablePlatformRequest))
	}
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_project.GetRequest))
	}
}
