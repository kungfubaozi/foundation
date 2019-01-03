package functionmw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
)

func Middleware(functioncli function.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			resp, err := functioncli.Get(ctx, &fs_base_function.GetRequest{
				Api: meta.Api,
			})
			if err != nil {
				return errno.ErrResponse(errno.ErrSystem)
			}
			if !resp.State.Ok {
				if resp.State == errno.ErrInvalid {
					return next(ctx, request)
				}
				return errno.ErrResponse(resp.State)
			}
			ctx = context.WithValue(ctx, "func", resp.Func)
			return next(ctx, request)
		}
	}
}
