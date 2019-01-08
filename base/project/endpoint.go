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
)

type Endpoints struct {
	NewEndpoint            endpoint.Endpoint
	GetEndpoint            endpoint.Endpoint
	EnablePlatformEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {
	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = MakeNewEndpoint(svc)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(newEndpoint)
		newEndpoint = zipkin.TraceEndpoint(trace, "New")(newEndpoint)
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
	}

	return Endpoints{
		NewEndpoint:            newEndpoint,
		GetEndpoint:            getEndpoint,
		EnablePlatformEndpoint: enablePlatform,
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

func (g Endpoints) EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
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
