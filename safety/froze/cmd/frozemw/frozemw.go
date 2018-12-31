package frozemw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/safety/froze"
	"zskparker.com/foundation/safety/froze/pb"
)

//检查冻结
func Middleware(svc froze.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			resp, err := svc.Check(ctx, &fs_safety_froze.CheckRequest{
				UserId: meta.UserId,
			})
			if err != nil {
				return errno.ErrResponse(errno.ErrSystem)
			}
			if !resp.State.Ok {
				return errno.ErrResponse(resp.State)
			}
			return next(ctx, request)
		}
	}
}
