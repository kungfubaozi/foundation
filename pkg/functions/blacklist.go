package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

//添加黑名单
func GetAddBlacklistFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/blacklist",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "添加黑名单",
		En:    "AddBlacklist",
		Func:  "01d69bf4f2f9",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_ADMIN,
	}
	return f
}

//移除黑名单
func GetRemoveBlacklistFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/blacklist",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "移除黑名单",
		En:    "RemoveBlacklist",
		Func:  "17bb8ce81629",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_ADMIN,
	}
	return f
}
