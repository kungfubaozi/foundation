package refresh

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/refresh/pb"
)

type Endpoints struct {
	AuthEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, clients *functionmw.MWServices) Endpoints {

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = MakeAuthEndpoint(svc)
		authEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(authEndpoint)
		authEndpoint = zipkin.TraceEndpoint(trace, "Auth")(authEndpoint)

		authEndpoint = functionmw.WithMeta(logger, clients)(authEndpoint)
	}

	return Endpoints{
		AuthEndpoint: authEndpoint,
	}

}

func (g Endpoints) Auth(ctx context.Context, in *fs_base_refresh.AuthRequest) (*fs_base_refresh.AuthResponse, error) {
	resp, err := g.AuthEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_refresh.AuthResponse), nil
}

func MakeAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Auth(ctx, request.(*fs_base_refresh.AuthRequest))
	}
}
