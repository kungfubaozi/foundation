package reporter

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	write grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints) fs_base_reporter.ReporterServer {
	options := []grpctransport.ServerOption{}

	return &GRPCServer{
		write: grpctransport.NewServer(
			endpoints.WriteEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			options...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	options := []grpctransport.ClientOption{

		grpctransport.ClientBefore(format.GRPCMetadata()),
	}
	var writeEndpoint endpoint.Endpoint
	{
		writeEndpoint = grpctransport.NewClient(conn,
			"fs.base.reporter.Reporter",
			"Write",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			options...).Endpoint()
		writeEndpoint = limiter(writeEndpoint)
		writeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Write",
			Timeout: 5 * time.Second,
		}))(writeEndpoint)
	}

	return Endpoints{
		WriteEndpoint: writeEndpoint,
	}
}

func (g *GRPCServer) Write(ctx context.Context, in *fs_base_reporter.WriteRequest) (*fs_base.Response, error) {
	_, resp, err := g.write.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}
