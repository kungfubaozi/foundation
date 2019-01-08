package verification

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetRegisterFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification/new",
		Infix:  "/register",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建注册验证码",
		En:    "NewRegisterVerification",
		Func:  "3e4f9c2e8530",
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}
