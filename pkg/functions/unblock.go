package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

//通过手机验证即可解除对账号的封锁
func GetUnlockFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/unblock",
		Infix:  "/unlock",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "账号解封",
		En:    "UnlockAccount",
		Func:  "0d42b31eee5e",
		Fcv:   fs_constants.FCV_VALIDATE_CODE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}
