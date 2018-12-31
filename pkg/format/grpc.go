package format

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
)

func GrpcMessage(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func GRPCMetadata() grpc.ClientRequestFunc {
	return func(i context.Context, md *metadata.MD) context.Context {
		return i
	}
}

func GRPCServerMetadata() grpc.ServerRequestFunc {
	return func(i context.Context, mds metadata.MD) context.Context {
		return i
	}
}
