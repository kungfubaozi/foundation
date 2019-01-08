package verification

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/tags"
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
		Func:  fs_function_tags.GetVerificationRegisterFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 0,
	}
	return f
}

func GetAdminRegisterFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification/new",
		Infix:  "/init",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建系统初始化验证码",
		En:    "NewInitVerification",
		Func:  fs_function_tags.GetVerificationAdminInitFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 0,
	}
	return f
}

func GetLoginFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification/new",
		Infix:  "/login",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建登录验证码",
		En:    "NewLoginVerification",
		Func:  fs_function_tags.GetVerificationLoginFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 0,
	}
	return f
}
