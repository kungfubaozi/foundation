package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetInterceptFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/intercept",
		Infix:  "/authorize",
		Function: &fs_base_function.Func{
			En:    "Authorize Intercept",
			Zh:    "鉴权拦截",
			Func:  fs_function_tags.GetInterceptFuncTag(),
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_TOURIST,
			Type:  fs_constants.TYPE_HIDE,
		},
	}
}
