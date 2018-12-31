package blacklistmw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/blacklist/pb"
)

//拦截
func Middleware(blacklistcli blacklist.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			resp, err := blacklistcli.Check(ctx, &fs_safety_blacklist.CheckRequest{
				UserId: meta.UserId,
				Ip:     meta.Ip,
				Device: meta.Device,
			})
			if err != nil {
				return errno.ErrResponse(errno.ErrRequest)
			}
			if !resp.State.Ok {
				return errno.ErrResponse(resp.State)
			}
			return next(ctx, request)
		}
	}
}
