package authenticatemw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/authenticate"
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
			return next(ctx, request)
		}
	}
}
