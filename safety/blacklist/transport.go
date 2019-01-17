package blacklist

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
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type GRPCServer struct {
	checkMeta    grpctransport.Handler
	checkAccount grpctransport.Handler
	add          grpctransport.Handler
}

func (g *GRPCServer) Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error) {
	_, resp, err := g.add.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) CheckAccount(ctx context.Context, in *fs_safety_blacklist.CheckAccountRequest) (*fs_base.Response, error) {
	_, resp, err := g.checkAccount.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) CheckMeta(ctx context.Context, in *fs_safety_blacklist.CheckMetaRequest) (*fs_base.Response, error) {
	_, resp, err := g.checkMeta.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle(fs_functions.GetAddBlacklistFunc().Infix, httptransport.NewServer(
		endpoints.AddEndpoint,
		decodeHTTPAdd,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "FromAP", logger)))...,
	))

	return m
}

func decodeHTTPAdd(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_safety_blacklist.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_safety_blacklist.BlacklistServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}

	return &GRPCServer{
		checkAccount: grpctransport.NewServer(
			endpoints.CheckAccountEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "CheckAccount", logger)))...),
		checkMeta: grpctransport.NewServer(
			endpoints.CheckMetaEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "CheckMeta", logger)))...),
		add: grpctransport.NewServer(
			endpoints.AddEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Add", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
	}

	var checkMetaEndpoint endpoint.Endpoint
	{
		checkMetaEndpoint = grpctransport.NewClient(conn,
			"fs.safety.blacklist.Blacklist",
			"CheckMeta",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		checkMetaEndpoint = limiter(checkMetaEndpoint)
		checkMetaEndpoint = opentracing.TraceClient(otTracer, "CheckMeta")(checkMetaEndpoint)
		checkMetaEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "CheckMeta",
			Timeout: 5 * time.Second,
		}))(checkMetaEndpoint)
	}

	var checkAccountEndpoint endpoint.Endpoint
	{
		checkAccountEndpoint = grpctransport.NewClient(conn,
			"fs.safety.blacklist.Blacklist",
			"CheckAccount",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		checkAccountEndpoint = limiter(checkAccountEndpoint)
		checkAccountEndpoint = opentracing.TraceClient(otTracer, "CheckAccount")(checkAccountEndpoint)
		checkAccountEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "CheckAccount",
			Timeout: 5 * time.Second,
		}))(checkAccountEndpoint)
	}

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = grpctransport.NewClient(conn,
			"fs.safety.blacklist.Blacklist",
			"Add",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		addEndpoint = limiter(addEndpoint)
		addEndpoint = opentracing.TraceClient(otTracer, "Add")(addEndpoint)
		addEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Add",
			Timeout: 5 * time.Second,
		}))(addEndpoint)
	}

	return Endpoints{
		CheckMetaEndpoint:    checkMetaEndpoint,
		CheckAccountEndpoint: checkAccountEndpoint,
		AddEndpoint:          addEndpoint,
	}
}
