package user

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/user/pb"
)

type Endpoints struct {
	AddEndpoint              endpoint.Endpoint
	FindByUserIdEndpoint     endpoint.Endpoint
	FindByPhoneEndpoint      endpoint.Endpoint
	FindByEmailEndpoint      endpoint.Endpoint
	FindByEnterpriseEndpoint endpoint.Endpoint
	UpdatePhoneEndpoint      endpoint.Endpoint
	UpdateEmailEndpoint      endpoint.Endpoint
	UpdateEnterpriseEndpoint endpoint.Endpoint
	UpdatePasswordEndpoint   endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = MakeAddEndpoint(svc)
		addEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(addEndpoint)
		addEndpoint = zipkin.TraceEndpoint(trace, "Add")(addEndpoint)
	}

	var findByUserIdEndpoint endpoint.Endpoint
	{
		findByUserIdEndpoint = MakeFindByUserIdEndpoint(svc)
		findByUserIdEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(findByUserIdEndpoint)
		findByUserIdEndpoint = zipkin.TraceEndpoint(trace, "FindByUserId")(findByUserIdEndpoint)
	}

	var findByPhoneEndpoint endpoint.Endpoint
	{
		findByPhoneEndpoint = MakeFindByPhoneEndpoint(svc)
		findByPhoneEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(findByPhoneEndpoint)
		findByPhoneEndpoint = zipkin.TraceEndpoint(trace, "FindByPhone")(findByPhoneEndpoint)
	}

	var findByEmailEndpoint endpoint.Endpoint
	{
		findByEmailEndpoint = MakeFindByEmailEndpoint(svc)
		findByEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(findByEmailEndpoint)
		findByEmailEndpoint = zipkin.TraceEndpoint(trace, "FindByEmail")(findByEmailEndpoint)
	}

	var findByEnterpriseEndpoint endpoint.Endpoint
	{
		findByEnterpriseEndpoint = MakeFindByEnterpriseEndpoint(svc)
		findByEnterpriseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(findByEnterpriseEndpoint)
		findByEnterpriseEndpoint = zipkin.TraceEndpoint(trace, "FindByEnterprise")(findByEnterpriseEndpoint)
	}

	var updatePhoneEndpoint endpoint.Endpoint
	{
		updatePhoneEndpoint = MakeUpdatePhoneEndpoint(svc)
		updatePhoneEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePhoneEndpoint)
		updatePhoneEndpoint = zipkin.TraceEndpoint(trace, "UpdatePhone")(updatePhoneEndpoint)
	}

	var updateEmailEndpoint endpoint.Endpoint
	{
		updateEmailEndpoint = MakeUpdateEmailEndpoint(svc)
		updateEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEmailEndpoint)
		updateEmailEndpoint = zipkin.TraceEndpoint(trace, "UpdateEmail")(updateEmailEndpoint)
	}

	var updateEnterpriseEndpoint endpoint.Endpoint
	{
		updateEnterpriseEndpoint = MakeUpdateEnterpriseEndpoint(svc)
		updateEnterpriseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updateEnterpriseEndpoint)
		updateEnterpriseEndpoint = zipkin.TraceEndpoint(trace, "UpdateEnterprise")(updateEnterpriseEndpoint)
	}

	var updatePasswordEndpoint endpoint.Endpoint
	{
		updatePasswordEndpoint = MakeUpdatePasswordEndpoint(svc)
		updatePasswordEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(updatePasswordEndpoint)
		updatePasswordEndpoint = zipkin.TraceEndpoint(trace, "UpdatePassword")(updatePasswordEndpoint)
	}

	return Endpoints{
		AddEndpoint:              addEndpoint,
		FindByEmailEndpoint:      findByEmailEndpoint,
		FindByEnterpriseEndpoint: findByEnterpriseEndpoint,
		FindByPhoneEndpoint:      findByPhoneEndpoint,
		FindByUserIdEndpoint:     findByUserIdEndpoint,
		UpdateEmailEndpoint:      updateEmailEndpoint,
		UpdateEnterpriseEndpoint: updateEnterpriseEndpoint,
		UpdatePasswordEndpoint:   updatePasswordEndpoint,
		UpdatePhoneEndpoint:      updatePhoneEndpoint,
	}

}

func (g Endpoints) Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error) {
	resp, err := g.AddEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) FindByUserId(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	resp, err := g.FindByUserIdEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g Endpoints) FindByPhone(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	resp, err := g.FindByPhoneEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g Endpoints) FindByEnterprise(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	resp, err := g.FindByEnterpriseEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g Endpoints) FindByEmail(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	resp, err := g.FindByEmailEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g Endpoints) UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdatePhoneEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdateEnterpriseEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdateEmailEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	resp, err := g.UpdatePasswordEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Add(ctx, request.(*fs_base_user.AddRequest))
	}
}

func MakeFindByEmailEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FindByEmail(ctx, request.(*fs_base_user.FindRequest))
	}
}

func MakeFindByEnterpriseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FindByEnterprise(ctx, request.(*fs_base_user.FindRequest))
	}
}

func MakeFindByPhoneEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FindByPhone(ctx, request.(*fs_base_user.FindRequest))
	}
}

func MakeFindByUserIdEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FindByUserId(ctx, request.(*fs_base_user.FindRequest))
	}
}

func MakeUpdatePhoneEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdatePhone(ctx, request.(*fs_base_user.UpdateRequest))
	}
}

func MakeUpdateEmailEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdateEmail(ctx, request.(*fs_base_user.UpdateRequest))
	}
}

func MakeUpdateEnterpriseEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdateEnterprise(ctx, request.(*fs_base_user.UpdateRequest))
	}
}

func MakeUpdatePasswordEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.UpdatePassword(ctx, request.(*fs_base_user.UpdateRequest))
	}
}
