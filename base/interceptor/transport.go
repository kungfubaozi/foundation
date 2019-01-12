package interceptor

import (
	"encoding/json"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net/http"
	"time"
	"zskparker.com/foundation/base/interceptor/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	auth grpctransport.Handler
}

func (g *GRPCServer) Auth(ctx context.Context, in *fs_base_interceptor.AuthRequest) (*fs_base_interceptor.AuthResponse, error) {
	_, resp, err := g.auth.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_interceptor.AuthResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_interceptor.AuthResponse), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle(fs_functions.GetInterceptFunc().Infix, httptransport.NewServer(
		endpoints.AuthEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Auth", logger)))...,
	))

	return m
}

func decodeHTPPUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_interceptor.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_interceptor.InterceptorServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}

	return &GRPCServer{
		auth: grpctransport.NewServer(
			endpoints.AuthEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Auth", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
	}

	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = grpctransport.NewClient(conn,
			"fs.base.interceptor.Interceptor",
			"Auth",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_interceptor.AuthResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		authEndpoint = limiter(authEndpoint)
		authEndpoint = opentracing.TraceClient(otTracer, "Auth")(authEndpoint)
		authEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Auth",
			Timeout: 5 * time.Second,
		}))(authEndpoint)
	}

	return Endpoints{
		AuthEndpoint: authEndpoint,
	}
}
