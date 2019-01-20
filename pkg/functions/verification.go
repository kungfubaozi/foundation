package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetVerificationRegisterFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/register",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建注册验证码",
		En:    "NewRegisterVerification",
		Func:  fs_function_tags.GetVerificationRegisterFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_NONE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetVerificationLoginFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/login",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建登录验证码",
		En:    "NewLoginVerification",
		Func:  fs_function_tags.GetVerificationLoginFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_NONE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetVerificationUpdatePasswordFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/updatePw",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建更新密码验证码",
		En:    "NewUpdatePasswordVerification",
		Func:  fs_function_tags.GetVerificationUpdatePasswordFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_NONE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetVerificationResetPasswordFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/verification",
		Infix:  "/resetPw",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建重置密码验证码",
		En:    "NewResetPasswordVerification",
		Func:  fs_function_tags.GetVerificationResetPasswordFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_NONE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}
