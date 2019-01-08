package update

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/tags"
)

func GetUpdatePhoneFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/update",
		Infix:  "/phone",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新手机",
		En:    "UpdatePhone",
		Func:  fs_function_tags.GetUpdatePhoneFuncTag(),
		Fcv:   names.F_FCV_AUTH | names.F_FCV_VALIDATE_CODE,
		Level: 1,
	}
	return f
}

func GetUpdateEmailFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/update",
		Infix:  "/email",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新邮箱",
		En:    "UpdateEmail",
		Func:  fs_function_tags.GetUpdateEmailFuncTag(),
		Fcv:   names.F_FCV_AUTH | names.F_FCV_VALIDATE_CODE,
		Level: 1,
	}
	return f
}

//更新指定用户的企业号
func GetUpdateEnterpriseFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/update",
		Infix:  "/enterprise",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新企业号",
		En:    "UpdateEnterprise",
		Func:  fs_function_tags.GetUpdateEnterpriseFuncTag(),
		Fcv:   names.F_FCV_AUTH,
		Level: 5,
	}
	return f
}

//更新密码不需要登录，使用绑定的手机即可
func GetUpdatePasswordFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/safety/update",
		Infix:  "/password",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新密码",
		En:    "UpdatePassword",
		Func:  fs_function_tags.GetUpdatePasswordFuncTag(),
		Fcv:   names.F_FCV_VALIDATE_CODE,
		Level: 1,
	}
	return f
}
