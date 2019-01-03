package projectmw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/pkg/errno"
)

func Middleware(projectcli project.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			resp, err := projectcli.Get(ctx, &fs_base_project.GetRequest{
				ClientId: meta.ClientId,
			})
			if !resp.State.Ok {
				return errno.ErrResponse(resp.State)
			}
			ctx = context.WithValue(ctx, "strategy", resp.Strategy)
			ctx = context.WithValue(ctx, "project", resp.Info)
			return next(ctx, request)
		}
	}
}
