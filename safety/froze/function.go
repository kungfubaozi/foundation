package froze

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetRequestFrozeFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/froze",
		Infix:  "/lock",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "冻结账户",
		En:    "FrozeAccount",
		Func:  "e66f94e8b5cc",
		Fcv:   names.F_FCV_NONE,
		Level: 1,
	}
	return f
}
