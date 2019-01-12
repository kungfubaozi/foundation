package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetRegisterFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/register",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建注册验证码",
		En:    "NewRegisterVerification",
		Func:  fs_function_tags.GetVerificationRegisterFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetLoginFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/login",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建登录验证码",
		En:    "NewLoginVerification",
		Func:  fs_function_tags.GetVerificationLoginFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}
