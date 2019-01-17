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
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	new         grpctransport.Handler
	check       grpctransport.Handler
	refresh     grpctransport.Handler
	offlineUser grpctransport.Handler
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
		endpoints.RefreshEndpoint,
		decodeHTTPRefresh,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Refresh", logger)))...,
	))

	return m
}

func decodeHTTPRefresh(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_authenticate.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func (g *GRPCServer) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	_, resp, err := g.new.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_authenticate.NewResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_authenticate.NewResponse), nil
}

func (g *GRPCServer) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error) {
	_, resp, err := g.check.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_authenticate.CheckResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_authenticate.CheckResponse), nil
}

func (g *GRPCServer) OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error) {
	_, resp, err := g.offlineUser.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Refresh(ctx context.Context, in *fs_base_authenticate.RefreshRequest) (*fs_base_authenticate.RefreshResponse, error) {
	_, resp, err := g.refresh.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_authenticate.RefreshResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_authenticate.RefreshResponse), nil
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_authenticate.AuthenticateServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
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
		offlineUser: grpctransport.NewServer(
			endpoints.OfflineUserEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "OfflineUser", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
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
			fs_base_authenticate.CheckResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		checkEndpoint = limiter(checkEndpoint)
		checkEndpoint = opentracing.TraceClient(otTracer, "Check")(checkEndpoint)
		checkEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Check",
			Timeout: 5 * time.Second,
		}))(checkEndpoint)
	}

	var offlineUserEndpoint endpoint.Endpoint
	{
		offlineUserEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"OfflineUser",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		offlineUserEndpoint = limiter(offlineUserEndpoint)
		offlineUserEndpoint = opentracing.TraceClient(otTracer, "OfflineUser")(offlineUserEndpoint)
		offlineUserEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "OfflineUser",
			Timeout: 5 * time.Second,
		}))(offlineUserEndpoint)
	}

	var refreshEndpoint endpoint.Endpoint
	{
		refreshEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"Refresh",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_authenticate.RefreshResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		refreshEndpoint = limiter(refreshEndpoint)
		refreshEndpoint = opentracing.TraceClient(otTracer, "Refresh")(refreshEndpoint)
		refreshEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Refresh",
			Timeout: 5 * time.Second,
		}))(refreshEndpoint)
	}

	return Endpoints{
		NewEndpoint:         newEndpoint,
		CheckEndpoint:       checkEndpoint,
		OfflineUserEndpoint: offlineUserEndpoint,
		RefreshEndpoint:     refreshEndpoint,
	}
}
