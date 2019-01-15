package authenticate

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
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	new         grpctransport.Handler
	check       grpctransport.Handler
	get         grpctransport.Handler
	replaceAuth grpctransport.Handler
	offlineUser grpctransport.Handler
}

func (g *GRPCServer) New(ctx context.Context, in *fs_base_authenticate.NewRequest) (*fs_base_authenticate.NewResponse, error) {
	_, resp, err := g.new.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.NewResponse), nil
}

func (g *GRPCServer) Check(ctx context.Context, in *fs_base_authenticate.CheckRequest) (*fs_base_authenticate.CheckResponse, error) {
	_, resp, err := g.check.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.CheckResponse), nil
}

func (g *GRPCServer) Get(ctx context.Context, in *fs_base_authenticate.GetRequest) (*fs_base_authenticate.GetResponse, error) {
	_, resp, err := g.get.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_authenticate.GetResponse), nil
}

func (g *GRPCServer) OfflineUser(ctx context.Context, in *fs_base_authenticate.OfflineUserRequest) (*fs_base.Response, error) {
	_, resp, err := g.offlineUser.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) ReplaceAuth(ctx context.Context, in *fs_base_authenticate.ReplaceAuthRequest) (*fs_base.Response, error) {
	_, resp, err := g.replaceAuth.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
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
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Get", logger)))...),
		replaceAuth: grpctransport.NewServer(
			endpoints.ReplaceAuthEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "ReplaceAuth", logger)))...),
		offlineUser: grpctransport.NewServer(
			endpoints.OfflineUserEndpoint,
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

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"Get",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_authenticate.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		getEndpoint = limiter(getEndpoint)
		getEndpoint = opentracing.TraceClient(otTracer, "Get")(getEndpoint)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Get",
			Timeout: 5 * time.Second,
		}))(getEndpoint)
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

	var replaceAuthEndpoint endpoint.Endpoint
	{
		replaceAuthEndpoint = grpctransport.NewClient(conn,
			"fs.base.authenticate.Authenticate",
			"ReplaceAuth",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		replaceAuthEndpoint = limiter(replaceAuthEndpoint)
		replaceAuthEndpoint = opentracing.TraceClient(otTracer, "ReplaceAuth")(replaceAuthEndpoint)
		replaceAuthEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "ReplaceAuth",
			Timeout: 5 * time.Second,
		}))(replaceAuthEndpoint)
	}

	return Endpoints{
		NewEndpoint:         newEndpoint,
		CheckEndpoint:       checkEndpoint,
		ReplaceAuthEndpoint: replaceAuthEndpoint,
		OfflineUserEndpoint: offlineUserEndpoint,
		GetEndpoint:         getEndpoint,
	}
}
