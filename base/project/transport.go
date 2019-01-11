package project

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
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	new            grpctransport.Handler
	get            grpctransport.Handler
	enablePlatform grpctransport.Handler
	init           grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_project.ProjectServer {
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
		new: grpctransport.NewServer(
			endpoints.NewEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "New", logger)))...),
		enablePlatform: grpctransport.NewServer(
			endpoints.EnablePlatformEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "EnablePlatform", logger)))...),
		init: grpctransport.NewServer(
			endpoints.InitEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Init", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1000))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
	}

	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(conn,
			"fs.base.project.Project",
			"Get",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_project.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		getEndpoint = limiter(getEndpoint)
		getEndpoint = opentracing.TraceClient(otTracer, "Get")(getEndpoint)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Get",
			Timeout: 5 * time.Second,
		}))(getEndpoint)
	}

	var newEndpoint endpoint.Endpoint
	{
		newEndpoint = grpctransport.NewClient(conn,
			"fs.base.project.Project",
			"New",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		newEndpoint = limiter(newEndpoint)
		newEndpoint = opentracing.TraceClient(otTracer, "New")(newEndpoint)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "New",
			Timeout: 5 * time.Second,
		}))(newEndpoint)
	}

	var enablePlatformEndpoint endpoint.Endpoint
	{
		enablePlatformEndpoint = grpctransport.NewClient(conn,
			"fs.base.project.Project",
			"EnablePlatform",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		enablePlatformEndpoint = limiter(enablePlatformEndpoint)
		enablePlatformEndpoint = opentracing.TraceClient(otTracer, "EnablePlatform")(enablePlatformEndpoint)
		enablePlatformEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "EnablePlatform",
			Timeout: 5 * time.Second,
		}))(enablePlatformEndpoint)
	}

	var initEndpoint endpoint.Endpoint
	{
		initEndpoint = grpctransport.NewClient(conn,
			"fs.base.project.Project",
			"Init",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_project.InitResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		initEndpoint = limiter(initEndpoint)
		initEndpoint = opentracing.TraceClient(otTracer, "Init")(initEndpoint)
		initEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Init",
			Timeout: 5 * time.Second,
		}))(initEndpoint)
	}

	return Endpoints{
		GetEndpoint:            getEndpoint,
		NewEndpoint:            newEndpoint,
		EnablePlatformEndpoint: enablePlatformEndpoint,
		InitEndpoint:           initEndpoint,
	}
}

func (g *GRPCServer) Init(ctx context.Context, in *fs_base_project.InitRequest) (*fs_base_project.InitResponse, error) {
	_, resp, err := g.init.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_project.InitResponse), nil
}

func (g *GRPCServer) New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error) {
	_, resp, err := g.new.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error) {
	_, resp, err := g.enablePlatform.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error) {
	_, resp, err := g.get.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_project.GetResponse), nil
}
