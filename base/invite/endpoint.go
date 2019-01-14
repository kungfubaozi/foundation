package invite

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/pb"
)

type Endpoints struct {
	AddEndpoint        endpoint.Endpoint
	UpdateEndpoint     endpoint.Endpoint
	GetEndpoint        endpoint.Endpoint
	GetInvitesEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, clients *functionmw.MWServices) Endpoints {

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = MakeAddEndpoint(svc)
		addEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addEndpoint)
		addEndpoint = zipkin.TraceEndpoint(trace, "Add")(addEndpoint)

		addEndpoint = functionmw.WithMeta(logger, clients)(addEndpoint)
	}

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = MakeUpdateEndpoint(svc)
		updateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEndpoint)
		updateEndpoint = zipkin.TraceEndpoint(trace, "Update")(updateEndpoint)
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = MakeGetEndpoint(svc)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getEndpoint)
		getEndpoint = zipkin.TraceEndpoint(trace, "Get")(getEndpoint)
	}

	var getInvitesEndpoint endpoint.Endpoint
	{
		getInvitesEndpoint = MakeGetInvitesEndpoint(svc)
		getInvitesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getInvitesEndpoint)
		getInvitesEndpoint = zipkin.TraceEndpoint(trace, "GetInvites")(getInvitesEndpoint)

		getInvitesEndpoint = functionmw.WithMeta(logger, clients)(getInvitesEndpoint)
	}

	return Endpoints{
		AddEndpoint:        addEndpoint,
		GetEndpoint:        getEndpoint,
		UpdateEndpoint:     updateEndpoint,
		GetInvitesEndpoint: getInvitesEndpoint,
	}
}

func (g Endpoints) Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error) {
	resp, err := g.AddEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) Get(ctx context.Context, in *fs_base_invite.GetRequest) (*fs_base_invite.GetResponse, error) {
	resp, err := g.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_invite.GetResponse), nil
}

func (g Endpoints) Update(ctx context.Context, in *fs_base_invite.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdateEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) GetInvites(ctx context.Context, in *fs_base_invite.GetInvitesRequest) (*fs_base_invite.GetInvitesResponse, error) {
	resp, err := g.GetInvitesEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_invite.GetInvitesResponse), nil
}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Add(ctx, request.(*fs_base_invite.AddRequest))
	}
}

func MakeUpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Update(ctx, request.(*fs_base_invite.UpdateRequest))
	}
}

func MakeGetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Get(ctx, request.(*fs_base_invite.GetRequest))
	}
}

func MakeGetInvitesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetInvites(ctx, request.(*fs_base_invite.GetInvitesRequest))
	}
}
