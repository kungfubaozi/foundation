package invite

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
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	add        grpctransport.Handler
	update     grpctransport.Handler
	get        grpctransport.Handler
	getInvites grpctransport.Handler
}

func (g *GRPCServer) Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error) {
	_, resp, err := g.add.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Get(ctx context.Context, in *fs_base_invite.GetRequest) (*fs_base_invite.GetResponse, error) {
	_, resp, err := g.get.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_invite.GetResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_invite.GetResponse), nil
}

func (g *GRPCServer) Update(ctx context.Context, in *fs_base_invite.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.update.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base.Response{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) GetInvites(ctx context.Context, in *fs_base_invite.GetInvitesRequest) (*fs_base_invite.GetInvitesResponse, error) {
	_, resp, err := g.getInvites.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_base_invite.GetInvitesResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_base_invite.GetInvitesResponse), nil
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle(fs_functions.GetInviteUserFunc().Infix, httptransport.NewServer(
		endpoints.AddEndpoint,
		decodeHTTPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Add", logger)))...,
	))

	return m
}

func decodeHTTPUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_invite.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_invite.InviteServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}
	return &GRPCServer{
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Get", logger)))...),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Update", logger)))...),
		getInvites: grpctransport.NewServer(
			endpoints.GetInvitesEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "GetInvites", logger)))...),
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
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
		zipkinClient,
	}
	var getEndpoint endpoint.Endpoint
	{
		getEndpoint = grpctransport.NewClient(conn,
			"fs.base.invite.Invite",
			"Get",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_invite.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		getEndpoint = limiter(getEndpoint)
		getEndpoint = opentracing.TraceClient(otTracer, "Get")(getEndpoint)
		getEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Get",
			Timeout: 5 * time.Second,
		}))(getEndpoint)
	}

	var getInvitesEndpoint endpoint.Endpoint
	{
		getInvitesEndpoint = grpctransport.NewClient(conn,
			"fs.base.invite.Invite",
			"GetInvites",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_invite.GetInvitesResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		getInvitesEndpoint = limiter(getInvitesEndpoint)
		getInvitesEndpoint = opentracing.TraceClient(otTracer, "GetInvites")(getInvitesEndpoint)
		getInvitesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "GetInvites",
			Timeout: 5 * time.Second,
		}))(getInvitesEndpoint)
	}

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = grpctransport.NewClient(conn,
			"fs.base.invite.Invite",
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

	var updateEndpoint endpoint.Endpoint
	{
		updateEndpoint = grpctransport.NewClient(conn,
			"fs.base.invite.Invite",
			"Update",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		updateEndpoint = limiter(updateEndpoint)
		updateEndpoint = opentracing.TraceClient(otTracer, "Update")(updateEndpoint)
		updateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Update",
			Timeout: 5 * time.Second,
		}))(updateEndpoint)
	}

	return &Endpoints{
		GetEndpoint:        getEndpoint,
		UpdateEndpoint:     updateEndpoint,
		GetInvitesEndpoint: getInvitesEndpoint,
		AddEndpoint:        addEndpoint,
	}
}
