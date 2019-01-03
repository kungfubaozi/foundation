package user

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
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	add              grpctransport.Handler
	findByUserId     grpctransport.Handler
	findByPhone      grpctransport.Handler
	findByEmail      grpctransport.Handler
	findByEnterprise grpctransport.Handler
	updatePhone      grpctransport.Handler
	updateEmail      grpctransport.Handler
	updateEnterprise grpctransport.Handler
	updatePassword   grpctransport.Handler
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(format.Metadata()),
		zipkinServer,
	}

	m := http.NewServeMux()
	m.Handle("/update/password", httptransport.NewServer(
		endpoints.UpdatePasswordEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "UpdatePassword", logger)))...,
	))

	m.Handle("/update/enterprise", httptransport.NewServer(
		endpoints.UpdateEnterpriseEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "UpdateEnterprise", logger)))...,
	))

	m.Handle("/update/phone", httptransport.NewServer(
		endpoints.FindByPhoneEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "UpdatePhone", logger)))...,
	))

	m.Handle("/update/email", httptransport.NewServer(
		endpoints.UpdateEmailEndpoint,
		decodeHTPPUpdate,
		format.EncodeHTTPGenericResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "UpdateEmail", logger)))...,
	))

	return m

}

func decodeHTPPUpdate(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_base_user.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func (g *GRPCServer) Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error) {
	_, resp, err := g.add.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) FindByUserId(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	_, resp, err := g.findByUserId.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g *GRPCServer) FindByEmail(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	_, resp, err := g.findByEmail.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g *GRPCServer) FindByPhone(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	_, resp, err := g.findByPhone.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g *GRPCServer) FindByEnterprise(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	_, resp, err := g.findByEnterprise.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_user.FindResponse), nil
}

func (g *GRPCServer) UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updatePhone.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updateEnterprise.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updatePassword.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	_, resp, err := g.updateEmail.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_user.UserServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		add: grpctransport.NewServer(
			endpoints.AddEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Add", logger)))...),
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
		updateEmail: grpctransport.NewServer(
			endpoints.UpdateEmailEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdateEmail", logger)))...),
		updatePhone: grpctransport.NewServer(
			endpoints.UpdatePhoneEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "UpdatePhone", logger)))...),
		findByEmail: grpctransport.NewServer(
			endpoints.FindByEmailEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FindByEmail", logger)))...),
		findByPhone: grpctransport.NewServer(
			endpoints.FindByPhoneEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FindByPhone", logger)))...),
		findByEnterprise: grpctransport.NewServer(
			endpoints.FindByEnterpriseEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FindByEnterprise", logger)))...),
		findByUserId: grpctransport.NewServer(
			endpoints.FindByUserIdEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "FindByUserId", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(format.GRPCMetadata()),
	}

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
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

	var updatePhoneEndpoint endpoint.Endpoint
	{
		updatePhoneEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
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

	var updateEmailEndpoint endpoint.Endpoint
	{
		updateEmailEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
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

	var updateEnterpriseEndpoint endpoint.Endpoint
	{
		updateEnterpriseEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
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

	var updatePasswordEndpoint endpoint.Endpoint
	{
		updatePasswordEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
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

	var findByUserIdEndpoint endpoint.Endpoint
	{
		findByUserIdEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
			"FindByUserId",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_user.FindResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		findByUserIdEndpoint = limiter(findByUserIdEndpoint)
		findByUserIdEndpoint = opentracing.TraceClient(otTracer, "FindByUserId")(findByUserIdEndpoint)
		findByUserIdEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FindByUserId",
			Timeout: 5 * time.Second,
		}))(findByUserIdEndpoint)
	}

	var findByEmailEndpoint endpoint.Endpoint
	{
		findByEmailEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
			"FindByEmail",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_user.FindResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		findByEmailEndpoint = limiter(findByEmailEndpoint)
		findByEmailEndpoint = opentracing.TraceClient(otTracer, "FindByEmail")(findByEmailEndpoint)
		findByEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FindByEmail",
			Timeout: 5 * time.Second,
		}))(findByEmailEndpoint)
	}

	var findByPhoneEndpoint endpoint.Endpoint
	{
		findByPhoneEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
			"FindByPhone",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_user.FindResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		findByPhoneEndpoint = limiter(findByPhoneEndpoint)
		findByPhoneEndpoint = opentracing.TraceClient(otTracer, "FindByPhone")(findByPhoneEndpoint)
		findByPhoneEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FindByPhone",
			Timeout: 5 * time.Second,
		}))(findByPhoneEndpoint)
	}

	var findByEnterpriseEndpoint endpoint.Endpoint
	{
		findByEnterpriseEndpoint = grpctransport.NewClient(conn,
			"fs.base.user.User",
			"FindByEnterprise",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_user.FindResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		findByEnterpriseEndpoint = limiter(findByEnterpriseEndpoint)
		findByEnterpriseEndpoint = opentracing.TraceClient(otTracer, "FindByEnterprise")(findByEnterpriseEndpoint)
		findByEnterpriseEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "FindByEnterprise",
			Timeout: 5 * time.Second,
		}))(findByEnterpriseEndpoint)
	}

	return Endpoints{
		AddEndpoint:              addEndpoint,
		UpdatePhoneEndpoint:      updatePhoneEndpoint,
		UpdateEmailEndpoint:      updateEmailEndpoint,
		UpdatePasswordEndpoint:   updatePasswordEndpoint,
		UpdateEnterpriseEndpoint: updateEnterpriseEndpoint,
		FindByUserIdEndpoint:     findByUserIdEndpoint,
		FindByEnterpriseEndpoint: findByEnterpriseEndpoint,
		FindByPhoneEndpoint:      findByPhoneEndpoint,
		FindByEmailEndpoint:      findByEmailEndpoint,
	}
}
