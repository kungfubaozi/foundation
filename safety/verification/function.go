package verification

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetNewFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/verification",
		Infix:  "/new",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建验证码",
		En:    "NewVerification",
		Func:  "508f0bb9c09d",
		Level: 1,
	}
	return f
}
