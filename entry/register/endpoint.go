package register

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/tags"
)

type Endpoints struct {
	FromAPEndpoint    endpoint.Endpoint
	FromOAuthEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client *functionmw.MWServices) Endpoints {

	var fromAPEndpoint endpoint.Endpoint
	{
		fromAPEndpoint = MakeFromAPEndpoint(svc)
		fromAPEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(fromAPEndpoint)
		fromAPEndpoint = zipkin.TraceEndpoint(trace, "FromAP")(fromAPEndpoint)

		fromAPEndpoint = functionmw.WithExpress(logger, client, fs_function_tags.GetFromAPFuncTag())(fromAPEndpoint)
	}

	var fromOAuthEndpoint endpoint.Endpoint
	{
		fromOAuthEndpoint = MakeFromOAuthEndpoint(svc)
		fromOAuthEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(fromOAuthEndpoint)
		fromOAuthEndpoint = zipkin.TraceEndpoint(trace, "FromOAuth")(fromOAuthEndpoint)

		fromOAuthEndpoint = functionmw.WithExpress(logger, client, fs_function_tags.GetFromOAuthFuncTag())(fromOAuthEndpoint)
	}

	return Endpoints{
		FromAPEndpoint:    fromAPEndpoint,
		FromOAuthEndpoint: fromOAuthEndpoint,
	}

}

func (g Endpoints) FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error) {
	resp, err := g.FromAPEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {
	resp, err := g.FromOAuthEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeFromAPEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FromAP(ctx, request.(*fs_entry_register.FromAPRequest))
	}
}

func MakeFromOAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.FromOAuth(ctx, request.(*fs_entry_register.FromOAuthRequest))
	}
}
