package verification

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/safety/verification/pb"
)

type Endpoints struct {
	NewEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = MakeNewEndpoint(svc)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(newEndpoint)
		newEndpoint = zipkin.TraceEndpoint(trace, "New")(newEndpoint)
	}

	return Endpoints{
		NewEndpoint: newEndpoint,
	}

}

func (g Endpoints) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	resp, err := g.NewEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_safety_verification.NewResponse), nil
}

func MakeNewEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.New(ctx, request.(*fs_safety_verification.NewRequest))
	}
}
