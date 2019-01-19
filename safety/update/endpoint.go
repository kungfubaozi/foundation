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
	"zskparker.com/foundation/pkg/middlewares"
	"zskparker.com/foundation/safety/update/pb"
)

type Endpoints struct {
	UpdatePhoneEndpoint    endpoint.Endpoint
	UpdateEmailEndpoint    endpoint.Endpoint
	ResetPasswordEndpoint  endpoint.Endpoint
	UpdatePasswordEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, clients fs_endpoint_middlewares.Endpoint) Endpoints {

	var updatePhone endpoint.Endpoint
	{
		updatePhone = MakeUpdatePhoneEndpoint(svc)
		updatePhone = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePhone)
		updatePhone = zipkin.TraceEndpoint(trace, "UpdatePhone")(updatePhone)

		updatePhone = clients.WithMeta()(updatePhone)
	}

	var updateEmail endpoint.Endpoint
	{
		updateEmail = MakeUpdateEmailEndpoint(svc)
		updateEmail = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEmail)
		updateEmail = zipkin.TraceEndpoint(trace, "UpdateEmail")(updateEmail)

		updateEmail = clients.WithMeta()(updateEmail)
	}

	var resetPasswordEndpoint endpoint.Endpoint
	{
		resetPasswordEndpoint = MakeResetPasswordEndpoint(svc)
		resetPasswordEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(resetPasswordEndpoint)
		resetPasswordEndpoint = zipkin.TraceEndpoint(trace, "ResetPassword")(resetPasswordEndpoint)

		resetPasswordEndpoint = clients.WithMeta()(resetPasswordEndpoint)
	}

	var updatePassword endpoint.Endpoint
	{
		updatePassword = MakeUpdatePasswordEndpoint(svc)
		updatePassword = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePassword)
		updatePassword = zipkin.TraceEndpoint(trace, "UpdatePassword")(updatePassword)

		updatePassword = clients.WithMeta()(updatePassword)
	}

	return Endpoints{
		UpdatePhoneEndpoint:    updatePhone,
		ResetPasswordEndpoint:  resetPasswordEndpoint,
		UpdatePasswordEndpoint: updatePassword,
		UpdateEmailEndpoint:    updateEmail,
	}
}

func (g Endpoints) UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdatePhoneEndpoint(ctx, in)
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

func (g Endpoints) ResetPassword(ctx context.Context, in *fs_safety_update.ResetPasswordRequest) (*fs_base.Response, error) {
	resp, err := g.ResetPasswordEndpoint(ctx, in)
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

func MakeResetPasswordEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.ResetPassword(ctx, request.(*fs_safety_update.ResetPasswordRequest))
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
