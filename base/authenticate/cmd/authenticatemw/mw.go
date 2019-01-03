package authenticatemw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
)

//验证步骤
//auth：
//	检查用户
//	检查状态
//  检查鉴权信息
//mw：
//  当auth成功后检查策略
func Middleware(authenticatecli authenticate.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			strategy := ctx.Value("strategy").(*fs_base.ProjectStrategy)
			meta := ctx.Value("meta").(*fs_base.Metadata)

			var c int64
			if meta.Platform == names.F_PLATFORM_ANDROID {
				c = strategy.Events.OnLogin.MaxCountOfOnline.Android
			} else if meta.Platform == names.F_PLATFORM_WEB {
				c = strategy.Events.OnLogin.MaxCountOfOnline.Web
			} else if meta.Platform == names.F_PLATFORM_MAC_OS {
				c = strategy.Events.OnLogin.MaxCountOfOnline.MacOS
			} else if meta.Platform == names.F_PLATFORM_WINDOWD {
				c = strategy.Events.OnLogin.MaxCountOfOnline.Windows
			} else if meta.Platform == names.F_PLATFORM_IOS {
				c = strategy.Events.OnLogin.MaxCountOfOnline.IOS
			}

			resp, err := authenticatecli.Check(ctx, &fs_base_authenticate.CheckRequest{
				Metadata:                     meta,
				MaxOnlineCount:               c,
				AllowOtherProjectUserToLogin: strategy.Events.OnLogin.AllowOtherProjectUserToLogin == 2,
			})

			if err != nil {
				return errno.ErrResponse(errno.ErrSystem)
			}

			if !resp.State.Ok {
				return errno.ErrResponse(resp.State)
			}

			meta.ClientId = resp.ClientId
			meta.ProjectId = resp.ProjectId
			meta.UserId = resp.UserId
			meta.Level = resp.Level

			return next(ctx, request)
		}
	}
}
