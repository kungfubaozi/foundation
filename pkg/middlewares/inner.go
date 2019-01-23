package fs_endpoint_middlewares

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func ServiceAuth() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		}
	}
}
