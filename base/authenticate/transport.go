package authenticate

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
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	new     grpctransport.Handler
	check   grpctransport.Handler
	refresh grpctransport.Handler
}

func (g *GRPCServer) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	_, resp, err := g.new.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.NewResponse), nil
}

func (g *GRPCServer) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base.Response, error) {
	_, resp, err := g.check.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	_, resp, err := g.refresh.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.RefreshResponse), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(format.Metadata()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle("/refresh", httptransport.NewServer(
		endpoints.RefreshEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Sum", logger)))...,
	))

	return m
}
func decodeHTPPUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_authenticate.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_authenticate.AuthenticateServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		new: grpctransport.NewServer(
			endpoints.NewEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "New", logger)))...),
		check: grpctransport.NewServer(
			endpoints.CheckEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Check", logger)))...),
		refresh: grpctransport.NewServer(
			endpoints.RefreshEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Refresh", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(format.GRPCMetadata()),
	}

	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"New",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_authenticate.NewResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		newEndpoint = limiter(newEndpoint)
		newEndpoint = opentracing.TraceClient(otTracer, "New")(newEndpoint)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "New",
			Timeout: 5 * time.Second,
		}))(newEndpoint)
	}

	var checkEndpoint endpoint.Endpoint
	{
		checkEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"Check",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		checkEndpoint = limiter(checkEndpoint)
		checkEndpoint = opentracing.TraceClient(otTracer, "Check")(checkEndpoint)
		checkEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Check",
			Timeout: 5 * time.Second,
		}))(checkEndpoint)
	}

	var refreshEndpoint endpoint.Endpoint
	{
		refreshEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"Refresh",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		refreshEndpoint = limiter(refreshEndpoint)
		refreshEndpoint = opentracing.TraceClient(otTracer, "Refresh")(refreshEndpoint)
		refreshEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Refresh",
			Timeout: 5 * time.Second,
		}))(refreshEndpoint)
	}

	return Endpoints{
		NewEndpoint:     newEndpoint,
		CheckEndpoint:   checkEndpoint,
		RefreshEndpoint: refreshEndpoint,
	}
}
