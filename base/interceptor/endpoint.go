package interceptor

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/interceptor/pb"
)

type Endpoints struct {
	AuthEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client *functionmw.MWServices) Endpoints {

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = MakeAuthEndpoint(svc)
		authEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(authEndpoint)
		authEndpoint = zipkin.TraceEndpoint(trace, "Auth")(authEndpoint)

		authEndpoint = functionmw.WithMeta(logger, client)(authEndpoint)
	}

	return Endpoints{
		AuthEndpoint: authEndpoint,
	}

}

func (g Endpoints) Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error) {
	resp, err := g.AuthEndpoint(ctx, in)
	if err != nil {
		fmt.Println("auth")
		return nil, err
	}
	return resp.(*fs_base_interceptor.AuthResponse), nil
}

func MakeAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Auth(ctx, request.(*fs_base_interceptor.AuthRequest))
	}
}
