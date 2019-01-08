package verification

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
	"zskparker.com/foundation/pkg/format"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/verification/pb"
)

type GRPCServer struct {
	new grpctransport.Handler
}

func MakeHTTPHandler(endpoints Endpoints, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) http.Handler {
	zipkinServer := zipkin.HTTPServerTrace(zipkinTracer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerBefore(fs_metadata_transport.HTTPToContext()),
		zipkinServer,
	}

	m := http.NewServeMux()

	//register
	m.Handle(GetRegisterFunc().Infix, httptransport.NewServer(
		endpoints.NewEndpoint,
		decodeHTTPNewRegister,
		format.EncodeHTTPGenericResponse,
		append(options,
			httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "NewRegister", logger)))...,
	))

	//register admin
	m.Handle(GetAdminRegisterFunc().Infix, httptransport.NewServer(
		endpoints.NewEndpoint,
		decodeHTTPNewAdminRegister,
		format.EncodeHTTPGenericResponse,
		append(options,
			httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "NewAdminRegister", logger)))...,
	))

	//login
	m.Handle(GetLoginFunc().Infix, httptransport.NewServer(
		endpoints.NewEndpoint,
		decodeHTTPNewLogin,
		format.EncodeHTTPGenericResponse,
		append(options,
			httptransport.ServerBefore(opentracing.HTTPToContext(otTracer, "NewLogin", logger)))...,
	))

	return m
}

//注册管理员验证码
func decodeHTTPNewAdminRegister(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_safety_verification.NewRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if req != nil {
		req.Func = fs_function_tags.GetAdminFuncTag()
	}
	return req, err
}

//使用密码账号注册验证码
func decodeHTTPNewRegister(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_safety_verification.NewRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if req != nil {
		req.Func = fs_function_tags.GetFromAPFuncTag()
	}
	return req, err
}

//登录验证码
func decodeHTTPNewLogin(_ context.Context, r *http.Request) (interface{}, error) {
	var req *fs_safety_verification.NewRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if req != nil {
		req.Func = fs_function_tags.GetEntryByValidateCodeFuncTag()
	}
	return req, err
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_safety_verification.VerificationServer {
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
	}
}

func (g *GRPCServer) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	_, resp, err := g.new.ServeGRPC(ctx, in)
	if err != nil {
		return &fs_safety_verification.NewResponse{State: fs_metadata_transport.GetResponseState(err, resp)}, nil
	}
	return resp.(*fs_safety_verification.NewResponse), nil
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
			"fs.safety.verification.Verification",
			"New",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_safety_verification.NewResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		newEndpoint = limiter(newEndpoint)
		newEndpoint = opentracing.TraceClient(otTracer, "New")(newEndpoint)
		newEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "New",
			Timeout: 5 * time.Second,
		}))(newEndpoint)
	}

	return Endpoints{
		NewEndpoint: newEndpoint,
	}
}
