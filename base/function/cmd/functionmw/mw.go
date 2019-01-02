package functionmw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
)

func Middleware(functioncli function.Service, metagetter func(request interface{}) *fs_base.Meta) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			m := metagetter(request)
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
			//需要身份验证
			if resp.Verification && len(m.Validate) < 5 && len(m.Id) != 32 {
				return errno.ErrResponse(errno.ErrMetaValidate)
			}
			return next(ctx, request)
		}
	}
}
