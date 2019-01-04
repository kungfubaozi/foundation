package unblock

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

//通过手机验证即可解除对账号的封锁
func GetUnlockFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/unblock",
		Infix:  "/unlock",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "账号解封",
		En:    "UnlockAccount",
		Func:  "0d42b31eee5e",
		Fcv:   names.F_FCV_PHONE,
		Level: 1,
	}
	return f
}
