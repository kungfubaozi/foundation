package update

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
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/update/pb"
)

type GRPCServer struct {
	updatePhone      grpctransport.Handler
	updateEmail      grpctransport.Handler
	updateEnterprise grpctransport.Handler
	updatePassword   grpctransport.Handler
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle("/update/password", httptransport.NewServer(
		endpoints.UpdatePasswordEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Sum", logger)))...,
	))
	m.Handle("/update/phone", httptransport.NewServer(
		endpoints.UpdatePhoneEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Concat", logger)))...,
	))
	m.Handle("/update/enterprise", httptransport.NewServer(
		endpoints.UpdateEnterpriseEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Concat", logger)))...,
	))
	m.Handle("/update/email", httptransport.NewServer(
		endpoints.UpdateEmailEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "Concat", logger)))...,
	))
	return m
}

func decodeHTPPUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_safety_update.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_safety_update.UpdateServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(fs_metadata_transport.GRPCToContext()),
		zipkinServer,
	}
	return &GRPCServer{
		updatePhone: grpctransport.NewServer(
			endpoints.UpdatePhoneEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdatePhone", logger)))...),
		updateEmail: grpctransport.NewServer(
			endpoints.UpdateEmailEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdateEmail", logger)))...),
		updateEnterprise: grpctransport.NewServer(
			endpoints.UpdateEnterpriseEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdateEnterprise", logger)))...),
		updatePassword: grpctransport.NewServer(
			endpoints.UpdatePasswordEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdatePassword", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		grpctransport.ClientBefore(fs_metadata_transport.ContextToGRPC()),
		zipkinClient,
	}
	var updatePhoneEndpoint endpoint.Endpoint
	{
		updatePhoneEndpoint = grpctransport.NewClient(conn,
			"fs.safety.update.Update",
			"UpdatePhone",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		updatePhoneEndpoint = limiter(updatePhoneEndpoint)
		updatePhoneEndpoint = opentracing.TraceClient(otTracer, "UpdatePhone")(updatePhoneEndpoint)
		updatePhoneEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "UpdatePhone",
			Timeout: 5 * time.Second,
		}))(updatePhoneEndpoint)
	}
	var updateEnterpriseEndpoint endpoint.Endpoint
	{
		updateEnterpriseEndpoint = grpctransport.NewClient(conn,
			"fs.safety.update.Update",
			"UpdateEnterprise",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		updateEnterpriseEndpoint = limiter(updateEnterpriseEndpoint)
		updateEnterpriseEndpoint = opentracing.TraceClient(otTracer, "UpdateEnterprise")(updateEnterpriseEndpoint)
		updateEnterpriseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "UpdateEnterprise",
			Timeout: 5 * time.Second,
		}))(updateEnterpriseEndpoint)
	}
	var updateEmailEndpoint endpoint.Endpoint
	{
		updateEmailEndpoint = grpctransport.NewClient(conn,
			"fs.safety.update.Update",
			"UpdateEmail",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		updateEmailEndpoint = limiter(updateEmailEndpoint)
		updateEmailEndpoint = opentracing.TraceClient(otTracer, "UpdateEmail")(updateEmailEndpoint)
		updateEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "UpdateEmail",
			Timeout: 5 * time.Second,
		}))(updateEmailEndpoint)
	}
	var updatePasswordEndpoint endpoint.Endpoint
	{
		updatePasswordEndpoint = grpctransport.NewClient(conn,
			"fs.safety.update.Update",
			"UpdatePassword",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		updatePasswordEndpoint = limiter(updatePasswordEndpoint)
		updatePasswordEndpoint = opentracing.TraceClient(otTracer, "UpdatePassword")(updatePasswordEndpoint)
		updatePasswordEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "UpdatePassword",
			Timeout: 5 * time.Second,
		}))(updatePasswordEndpoint)
	}
	return &Endpoints{
		UpdatePhoneEndpoint:      updatePhoneEndpoint,
		UpdatePasswordEndpoint:   updatePasswordEndpoint,
		UpdateEnterpriseEndpoint: updateEnterpriseEndpoint,
		UpdateEmailEndpoint:      updateEmailEndpoint,
	}
}

func (g *GRPCServer) UpdatePhone(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updatePhone.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdateEnterprise(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updateEnterprise.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdatePassword(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updatePassword.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdateEmail(ctx context.Context, in *fs_safety_update.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updateEmail.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}
