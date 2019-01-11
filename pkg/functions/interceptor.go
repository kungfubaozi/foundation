package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/tags"
)

func GetInterceptFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/interceptor",
		Infix:  "/auth",
		Function: &fs_base_function.Func{
			En:    "Server Intercept",
			Zh:    "鉴权拦截",
			Func:  fs_function_tags.GetInterceptFuncTag(),
			Fcv:   names.F_FCV_AUTH | names.F_FCV_SESSION,
			Level: 1,
			Type:  names.F_FUNC_TYPE_HIDE,
		},
	}
}
