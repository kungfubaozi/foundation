package validate

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/validate/pb"
)

type Endpoints struct {
	VerificationEndpoint endpoint.Endpoint
	CreateEndpoint       endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var verificationEndpoint endpoint.Endpoint
	{
		verificationEndpoint = MakeVerificationEndpoint(svc)
		verificationEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(verificationEndpoint)
		verificationEndpoint = zipkin.TraceEndpoint(trace, "Verification")(verificationEndpoint)
	}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = MakeCreateEndpoint(svc)
		createEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(createEndpoint)
		createEndpoint = zipkin.TraceEndpoint(trace, "Create")(createEndpoint)
	}

	return Endpoints{
		VerificationEndpoint: verificationEndpoint,
		CreateEndpoint:       createEndpoint,
	}

}

func (g *Endpoints) Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base_validate.VerificationResponse, error) {
	resp, err := g.VerificationEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_validate.VerificationResponse), nil
}

func (g *Endpoints) Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error) {
	resp, err := g.CreateEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_validate.CreateResponse), nil
}

func MakeVerificationEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Verification(ctx, request.(*fs_base_validate.VerificationRequest))
	}
}

func MakeCreateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Create(ctx, request.(*fs_base_validate.CreateRequest))
	}
}
