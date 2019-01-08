package function

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	get grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_function.FunctionServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Get", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(conn,
			"fs.base.function.Function",
			"Get",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_function.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		getEndpoint = limiter(getEndpoint)
		getEndpoint = opentracing.TraceClient(otTracer, "Get")(getEndpoint)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Get",
			Timeout: 5 * time.Second,
		}))(getEndpoint)
	}

	return Endpoints{
		GetEndpoint: getEndpoint,
	}

}

func (g *GRPCServer) Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error) {
	_, resp, err := g.get.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_function.GetResponse), nil
}
