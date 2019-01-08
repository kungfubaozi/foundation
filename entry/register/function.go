package register

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetFromOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/register",
		Infix:  "/oauth",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "第三方账号注册",
		En:    "RegisterFromOAuth",
		Func:  "a1999c10319f",
		Fcv:   names.F_FCV_PHONE,
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetFromAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/register",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "账号注册",
		En:    "RegisterFromAP",
		Fcv:   names.F_FCV_PHONE,
		Func:  "ef274cc105ad",
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetAdminFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/register",
		Infix:  "/init_system",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "初始化",
		En:    "InitSys",
		Func:  "934d601db20d",
		Fcv:   names.F_FCV_PHONE,
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}
