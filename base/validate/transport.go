package validate

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
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	verification grpctransport.Handler
	create       grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_validate.ValidateServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		verification: grpctransport.NewServer(
			endpoints.VerificationEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Verification", logger)))...),
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Create", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
	}

	var verificationEndpoint endpoint.Endpoint
	{
		verificationEndpoint = grpctransport.NewClient(conn,
			"fs.base.validate.Validate",
			"Verification",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_validate.VerificationResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		verificationEndpoint = limiter(verificationEndpoint)
		verificationEndpoint = opentracing.TraceClient(otTracer, "Verification")(verificationEndpoint)
		verificationEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Verification",
			Timeout: 5 * time.Second,
		}))(verificationEndpoint)
	}

	var createEndpoint endpoint.Endpoint
	{
		createEndpoint = grpctransport.NewClient(conn,
			"fs.base.validate.Validate",
			"Create",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_validate.CreateResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		createEndpoint = limiter(createEndpoint)
		createEndpoint = opentracing.TraceClient(otTracer, "Create")(createEndpoint)
		createEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Create",
			Timeout: 5 * time.Second,
		}))(createEndpoint)
	}

	return &Endpoints{
		VerificationEndpoint: verificationEndpoint,
		CreateEndpoint:       createEndpoint,
	}
}

func (g *GRPCServer) Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base_validate.VerificationResponse, error) {
	_, resp, err := g.verification.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_validate.VerificationResponse), nil
}

func (g *GRPCServer) Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error) {
	_, resp, err := g.create.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_validate.CreateResponse), nil
}
