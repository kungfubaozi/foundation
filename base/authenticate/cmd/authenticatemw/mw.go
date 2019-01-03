package authenticatemw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
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
			stragety := ctx.Value("strategy").(*fs_base.ProjectStrategy)
			meta := ctx.Value("meta").(*fs_base.Metadata)

			var c int64
			if meta.Platform == names.F_PLATFORM_ANDROID {
				c = stragety.Events.OnLogin.MaxCountOfOnline.Android
			} else if meta.Platform == names.F_PLATFORM_WEB {
				c = stragety.Events.OnLogin.MaxCountOfOnline.Web
			} else if meta.Platform == names.F_PLATFORM_MAC_OS {
				c = stragety.Events.OnLogin.MaxCountOfOnline.MacOS
			} else if meta.Platform == names.F_PLATFORM_WINDOWD {
				c = stragety.Events.OnLogin.MaxCountOfOnline.Windows
			} else if meta.Platform == names.F_PLATFORM_IOS {
				c = stragety.Events.OnLogin.MaxCountOfOnline.IOS
			}

			resp, err := authenticatecli.Check(ctx, &fs_base_authenticate.CheckRequest{
				Metadata:       meta,
				MaxOnlineCount: c,
			})

			return next(ctx, request)
		}
	}
}
