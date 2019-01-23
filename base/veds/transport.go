package veds

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
	"zskparker.com/foundation/base/veds/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	encrypt grpctransport.Handler
	decrypt grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_veds.VEDSServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		encrypt: grpctransport.NewServer(
			endpoints.EncryptEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Encrypt", logger)))...),
		decrypt: grpctransport.NewServer(
			endpoints.DecryptEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Decrypt", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) fs_base_veds.VEDSServer {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1000))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
	}

	var encryptEndpoint endpoint.Endpoint
	{
		encryptEndpoint = grpctransport.NewClient(conn,
			"fs.base.veds.VEDS",
			"Encrypt",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_veds.CryptResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		encryptEndpoint = limiter(encryptEndpoint)
		encryptEndpoint = opentracing.TraceClient(otTracer, "Encrypt")(encryptEndpoint)
		encryptEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Encrypt",
			Timeout: 5 * time.Second,
		}))(encryptEndpoint)
	}

	var decryptEndpoint endpoint.Endpoint
	{
		decryptEndpoint = grpctransport.NewClient(conn,
			"fs.base.veds.VEDS",
			"Decrypt",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_veds.CryptResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		decryptEndpoint = limiter(decryptEndpoint)
		decryptEndpoint = opentracing.TraceClient(otTracer, "Decrypt")(decryptEndpoint)
		decryptEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Decrypt",
			Timeout: 5 * time.Second,
		}))(decryptEndpoint)
	}

	return Endpoints{
		EncryptEndpoint: encryptEndpoint,
		DecryptEndpoint: decryptEndpoint,
	}
}

func (g *GRPCServer) Encrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	_, resp, err := g.encrypt.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_veds.CryptResponse), nil
}

func (g *GRPCServer) Decrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	_, resp, err := g.decrypt.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_veds.CryptResponse), nil
}
