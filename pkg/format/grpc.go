package format

import "context"

func GrpcMessage(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}
