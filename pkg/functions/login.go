package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetEntryByInviteFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/invite",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "邀请码登录",
		En:    "EntryByInvite",
		Func:  fs_function_tags.GetEntryByInviteTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetEntryByFaceFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/face",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "刷脸登录",
		En:    "EntryByFace",
		Func:  fs_function_tags.GetEntryByFaceFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_ADMIN, //刷脸登录目前只对最高管理员开放
	}
	return f
}

func GetEntryByAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "密码登录",
		En:    "EntryByAP",
		Func:  fs_function_tags.GetEntryByAPFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetEntryByOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/oauth",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "第三方登录",
		En:    "EntryByOAuth",
		Func:  fs_function_tags.GetEntryByOAuthFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetEntryByValidateCodeFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/vc",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "验证码登录",
		En:    "EntryByValidateCode",
		Func:  fs_function_tags.GetEntryByValidateCodeFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetEntryByQRCodeFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/login",
		Infix:  "/qrcode",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "二维码登录",
		En:    "EntryByQRCode",
		Func:  fs_function_tags.GetEntryByQRCodeFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}
