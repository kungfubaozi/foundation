package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

//添加黑名单
func GetAddBlacklistFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/blacklist",
		Infix:  "/add",
		Function: &fs_base_function.Func{
			Zh:    "添加黑名单",
			En:    "AddBlacklist",
			Func:  fs_function_tags.GetAddBlacklist(),
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_ADMIN,
		},
	}
}

//移除黑名单
func GetRemoveBlacklistFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/blacklist",
		Infix:  "/remove",
		Function: &fs_base_function.Func{
			Zh:    "移除黑名单",
			En:    "RemoveBlacklist",
			Func:  fs_function_tags.GetRemoveBlacklist(),
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_ADMIN,
		},
	}
}
