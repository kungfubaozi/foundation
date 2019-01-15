package refresh

import (
	"context"
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
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"net/http"
	"time"
	"zskparker.com/foundation/base/refresh/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	auth grpctransport.Handler
}

func (g *GRPCServer) Auth(ctx context.Context, in *fs_base_refresh.AuthRequest) (*fs_base_refresh.AuthResponse, error) {
	_, resp, err := g.auth.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_refresh.AuthResponse), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle(fs_functions.GetRefreshFunc().Infix, httptransport.NewServer(
		endpoints.AuthEndpoint,
		decodeHTTPAuth,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Auth", logger)))...,
	))

	return m
}

func decodeHTTPAuth(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_refresh.AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_refresh.RefreshServer {
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
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
		zipkinClient,
	}
	var authEndpoint endpoint.Endpoint
	{
		authEndpoint = grpctransport.NewClient(conn,
			"fs.base.refresh.Refresh",
			"Auth",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_refresh.AuthResponse{},
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
