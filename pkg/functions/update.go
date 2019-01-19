package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetUpdatePhoneFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/update",
		Infix:  "/phone",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新手机",
		En:    "UpdatePhone",
		Func:  fs_function_tags.GetUpdatePhoneFuncTag(),
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_VALIDATE_CODE,
		Level: fs_constants.LEVEL_USER,
	}
	return f
}

func GetUpdateEmailFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/update",
		Infix:  "/email",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新邮箱",
		En:    "UpdateEmail",
		Func:  fs_function_tags.GetUpdateEmailFuncTag(),
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_VALIDATE_CODE,
		Level: fs_constants.LEVEL_USER,
	}
	return f
}

//更新指定用户的企业号
func GetResetPasswordFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/update",
		Infix:  "/resetPW",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "重置密码",
		En:    "ResetPassword",
		Func:  fs_function_tags.GetResetPasswordFuncTag(),
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

//更新密码
func GetUpdatePasswordFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/update",
		Infix:  "/password",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新密码",
		En:    "UpdatePassword",
		Func:  fs_function_tags.GetUpdatePasswordFuncTag(),
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_VALIDATE_CODE,
		Level: fs_constants.LEVEL_USER,
	}
	return f
}
