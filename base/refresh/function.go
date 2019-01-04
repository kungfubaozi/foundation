package refresh

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetRefreshFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/auth",
		Infix:  "/refresh",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "刷新鉴权",
		En:    "RefreshToken",
		Func:  "74d9a5fff564",
		Level: 1,
	}
	return f
}
