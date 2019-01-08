package register

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/tags"
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
		Fcv:   names.F_FCV_VALIDATE_CODE,
		Func:  fs_function_tags.GetFromOAuthFuncTag(),
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
		Fcv:   names.F_FCV_VALIDATE_CODE,
		Func:  fs_function_tags.GetFromAPFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetAdminFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/register",
		Infix:  "/init",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "初始化",
		En:    "InitSys",
		Fcv:   names.F_FCV_VALIDATE_CODE,
		Func:  fs_function_tags.GetAdminFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}
