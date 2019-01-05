package update

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/safety/update/pb"
)

type Endpoints struct {
	UpdatePhoneEndpoint      endpoint.Endpoint
	UpdateEmailEndpoint      endpoint.Endpoint
	UpdateEnterpriseEndpoint endpoint.Endpoint
	UpdatePasswordEndpoint   endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var updatePhone endpoint.Endpoint
	{
		updatePhone = MakeUpdatePhoneEndpoint(svc)
		updatePhone = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePhone)
		updatePhone = zipkin.TraceEndpoint(trace, "UpdatePhone")(updatePhone)

	}

	var updateEmail endpoint.Endpoint
	{
		updateEmail = MakeUpdateEmailEndpoint(svc)
		updateEmail = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEmail)
		updateEmail = zipkin.TraceEndpoint(trace, "UpdateEmail")(updateEmail)
	}

	var updateEnterprise endpoint.Endpoint
	{
		updateEnterprise = MakeUpdateEnterpriseEndpoint(svc)
		updateEnterprise = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEnterprise)
		updateEnterprise = zipkin.TraceEndpoint(trace, "UpdateEnterprise")(updateEnterprise)
	}

	var updatePassword endpoint.Endpoint
	{
		updatePassword = MakeUpdatePasswordEndpoint(svc)
		updatePassword = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePassword)
		updatePassword = zipkin.TraceEndpoint(trace, "UpdatePassword")(updatePassword)
	}

	return Endpoints{
		UpdatePhoneEndpoint:      updatePhone,
		UpdateEnterpriseEndpoint: updateEnterprise,
		UpdatePasswordEndpoint:   updatePassword,
		UpdateEmailEndpoint:      updateEmail,
	}
}

func (g Endpoints) UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdatePhoneEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdateEnterpriseEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdatePasswordEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdateEmailEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeUpdatePhoneEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdatePhone(ctx, request.(*fs_safety_update.UpdateRequest))
	}
}

func MakeUpdateEnterpriseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdateEnterprise(ctx, request.(*fs_safety_update.UpdateRequest))
	}
}

func MakeUpdatePasswordEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdatePassword(ctx, request.(*fs_safety_update.UpdateRequest))
	}
}

func MakeUpdateEmailEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdateEmail(ctx, request.(*fs_safety_update.UpdateRequest))
	}
}
