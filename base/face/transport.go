package face

import (
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
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	upsert  grpctransport.Handler
	search  grpctransport.Handler
	remove  grpctransport.Handler
	compare grpctransport.Handler
}

func (g *GRPCServer) Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error) {
	_, resp, err := g.compare.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error) {
	_, resp, err := g.search.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_face.SearchResponse), nil
}

func (g *GRPCServer) Upsert(ctx context.Context, in *fs_base_face.UpsertRequest) (*fs_base.Response, error) {
	_, resp, err := g.upsert.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) RemoveFace(ctx context.Context, in *fs_base_face.RemoveFaceRequest) (*fs_base.Response, error) {
	_, resp, err := g.remove.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeHTPPHandler() {

}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_face.FaceServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		upsert: grpctransport.NewServer(
			endpoints.UpsertEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Upsert", logger)))...),
		compare: grpctransport.NewServer(
			endpoints.CompareEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Compare", logger)))...),
		remove: grpctransport.NewServer(
			endpoints.RemoveFaceEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Remove", logger)))...),
		search: grpctransport.NewServer(
			endpoints.SearchEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "Search", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(format.GRPCMetadata()),
	}

	var compareEndpoint endpoint.Endpoint
	{
		compareEndpoint = grpctransport.NewClient(conn,
			"fs.base.face.Face",
			"Compare",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		compareEndpoint = limiter(compareEndpoint)
		compareEndpoint = opentracing.TraceClient(otTracer, "Compare")(compareEndpoint)
		compareEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Compare",
			Timeout: 30 * time.Second,
		}))(compareEndpoint)
	}

	var searchEndpoint endpoint.Endpoint
	{
		searchEndpoint = grpctransport.NewClient(conn,
			"fs.base.face.Face",
			"Search",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base_face.SearchResponse{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		searchEndpoint = limiter(searchEndpoint)
		searchEndpoint = opentracing.TraceClient(otTracer, "Search")(searchEndpoint)
		searchEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Search",
			Timeout: 30 * time.Second,
		}))(searchEndpoint)
	}

	var upsertEndpoint endpoint.Endpoint
	{
		upsertEndpoint = grpctransport.NewClient(conn,
			"fs.base.face.Face",
			"Upsert",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		upsertEndpoint = limiter(upsertEndpoint)
		upsertEndpoint = opentracing.TraceClient(otTracer, "Upsert")(upsertEndpoint)
		upsertEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Upsert",
			Timeout: 30 * time.Second,
		}))(upsertEndpoint)
	}

	var removeEndpoint endpoint.Endpoint
	{
		removeEndpoint = grpctransport.NewClient(conn,
			"fs.base.face.Face",
			"Remove",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		removeEndpoint = limiter(removeEndpoint)
		removeEndpoint = opentracing.TraceClient(otTracer, "Remove")(removeEndpoint)
		removeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Remove",
			Timeout: 30 * time.Second,
		}))(removeEndpoint)
	}

	return Endpoints{
		CompareEndpoint:    compareEndpoint,
		SearchEndpoint:     searchEndpoint,
		UpsertEndpoint:     upsertEndpoint,
		RemoveFaceEndpoint: removeEndpoint,
	}
}
