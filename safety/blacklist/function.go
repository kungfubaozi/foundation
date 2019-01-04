package blacklist

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

//添加黑名单
func GetAddBlacklistFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/blacklist",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "添加黑名单",
		En:    "AddBlacklist",
		Func:  "01d69bf4f2f9",
		Fcv:   names.F_FCV_AUTH,
		Level: 5,
	}
	return f
}

//移除黑名单
func GetRemoveBlacklistFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/blacklist",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "移除黑名单",
		En:    "RemoveBlacklist",
		Func:  "17bb8ce81629",
		Fcv:   names.F_FCV_AUTH,
		Level: 5,
	}
	return f
}
