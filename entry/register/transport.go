package register

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
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/transport"
)

type GRPCServer struct {
	fromap    grpctransport.Handler
	fromoauth grpctransport.Handler
	admin     grpctransport.Handler
}

func (g *GRPCServer) FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error) {
	_, resp, err := g.fromap.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {
	_, resp, err := g.fromoauth.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Admin(ctx context.Context, in *fs_entry_register.AdminRequest) (*fs_base.Response, error) {
	_, resp, err := g.admin.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
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
	m.Handle(GetFromAPFunc().Infix, httptransport.NewServer(
		endpoints.FromAPEndpoint,
		decodeFromAP,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "FromAP", logger)))...,
	))

	m.Handle(GetFromOAuthFunc().Infix, httptransport.NewServer(
		endpoints.FromAPEndpoint,
		decodeFromOAuth,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "FromOAuth", logger)))...,
	))

	m.Handle(GetAdminFunc().Infix, httptransport.NewServer(
		endpoints.AdminEndpoint,
		decodeAdmin,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Admin", logger)))...,
	))

	return m
}

func decodeFromAP(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_register.FromAPRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeFromOAuth(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_register.FromOAuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeAdmin(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_entry_register.AdminRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_entry_register.RegisterServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}

	return &GRPCServer{
		fromap: grpctransport.NewServer(
			endpoints.FromAPEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FromAP", logger)))...),
		fromoauth: grpctransport.NewServer(
			endpoints.FromOAuthEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FromOAuth", logger)))...),
		admin: grpctransport.NewServer(
			endpoints.AdminEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Admin", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
	}

	var fromAPEndpoint endpoint.Endpoint
	{
		fromAPEndpoint = grpctransport.NewClient(conn,
			"fs.entry.register.Register",
			"FromAP",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_strategy.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		fromAPEndpoint = limiter(fromAPEndpoint)
		fromAPEndpoint = opentracing.TraceClient(otTracer, "FromAP")(fromAPEndpoint)
		fromAPEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FromAP",
			Timeout: 5 * time.Second,
		}))(fromAPEndpoint)
	}

	var fromOAuthEndpoint endpoint.Endpoint
	{
		fromOAuthEndpoint = grpctransport.NewClient(conn,
			"fs.entry.register.Register",
			"FromOAuth",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_strategy.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		fromOAuthEndpoint = limiter(fromOAuthEndpoint)
		fromOAuthEndpoint = opentracing.TraceClient(otTracer, "FromOAuth")(fromOAuthEndpoint)
		fromOAuthEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FromOAuth",
			Timeout: 5 * time.Second,
		}))(fromOAuthEndpoint)
	}

	var adminEndpoint endpoint.Endpoint
	{
		adminEndpoint = grpctransport.NewClient(conn,
			"fs.entry.register.Register",
			"Admin",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_strategy.GetResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		adminEndpoint = limiter(adminEndpoint)
		adminEndpoint = opentracing.TraceClient(otTracer, "Admin")(adminEndpoint)
		adminEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Admin",
			Timeout: 5 * time.Second,
		}))(adminEndpoint)
	}

	return Endpoints{
		FromOAuthEndpoint: fromOAuthEndpoint,
		FromAPEndpoint:    fromAPEndpoint,
	}
}
