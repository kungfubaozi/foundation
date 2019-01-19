package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

//添加黑名单
func GetRefreshFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/auth",
		Infix:  "/refresh",
		Function: &fs_base_function.Func{
			Zh:    "刷新Token",
			En:    "RefreshToken",
			Func:  fs_function_tags.GetAuthenticateRefreshTag(),
			Fcv:   fs_constants.FCV_NONE,
			Level: fs_constants.LEVEL_TOURIST,
		},
	}
}
