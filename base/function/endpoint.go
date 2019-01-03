package function

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/cmd/authenticatemw"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/project/cmd/projectmw"
	"zskparker.com/foundation/base/validate/cmd/validatemw"
)

type Endpoints struct {
	AddEndpoint    endpoint.Endpoint
	RemoveEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger,
	authenticatecli authenticate.Service, projectcli project.Service, validatecli validatemw.Client) Endpoints {

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = MakeAddEndpoint(svc)
		addEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addEndpoint)
		addEndpoint = zipkin.TraceEndpoint(trace, "Add")(addEndpoint)

		addEndpoint = validatemw.Middleware(validatecli)(addEndpoint)
		addEndpoint = authenticatemw.Middleware(authenticatecli)(addEndpoint)
		addEndpoint = projectmw.Middleware(projectcli)(addEndpoint)
	}

	var removeEndpoint endpoint.Endpoint
	{
		removeEndpoint = MakeRemoveEndpoint(svc)
		removeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(removeEndpoint)
		removeEndpoint = zipkin.TraceEndpoint(trace, "Remove")(removeEndpoint)

		removeEndpoint = validatemw.Middleware(validatecli)(removeEndpoint)
		removeEndpoint = authenticatemw.Middleware(authenticatecli)(removeEndpoint)
		removeEndpoint = projectmw.Middleware(projectcli)(removeEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = MakeUpdateEndpoint(svc)
		updateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEndpoint)
		updateEndpoint = zipkin.TraceEndpoint(trace, "Update")(updateEndpoint)

		updateEndpoint = validatemw.Middleware(validatecli)(updateEndpoint)
		updateEndpoint = authenticatemw.Middleware(authenticatecli)(updateEndpoint)
		updateEndpoint = projectmw.Middleware(projectcli)(updateEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)

		getEndpoint = validatemw.Middleware(validatecli)(getEndpoint)
		getEndpoint = authenticatemw.Middleware(authenticatecli)(getEndpoint)
		getEndpoint = projectmw.Middleware(projectcli)(getEndpoint)
	}

	return Endpoints{
		AddEndpoint:    addEndpoint,
		RemoveEndpoint: removeEndpoint,
		UpdateEndpoint: updateEndpoint,
		GetEndpoint:    getEndpoint,
	}

}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Add(ctx, request.(*fs_base_function.UpsertRequest))
	}
}

func MakeRemoveEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Remove(ctx, request.(*fs_base_function.RemoveRequest))
	}
}

func MakeUpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Update(ctx, request.(*fs_base_function.UpsertRequest))
	}
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_function.GetRequest))
	}
}
