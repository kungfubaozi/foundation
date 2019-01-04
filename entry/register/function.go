package register

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetFromOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/register",
		Infix:  "/oauth",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "第三方账号注册",
		En:    "RegisterFromOAuth",
		Func:  "a1999c10319f",
		Level: 1,
	}
	return f
}

func GetFromAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/register",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "账号注册",
		En:    "RegisterFromAP",
		Func:  "ef274cc105ad",
		Level: 1,
	}
	return f
}
