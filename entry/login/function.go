package login

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/tags"
)

func GetEntryByFaceFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/entry",
		Infix:  "/face",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "刷脸登录",
		En:    "EntryByFace",
		Func:  fs_function_tags.GetEntryByFaceFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 5, //刷脸登录目前只对最高管理员开放
	}
	return f
}

func GetEntryByAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/entry",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "密码登录",
		En:    "EntryByAP",
		Func:  fs_function_tags.GetEntryByAPFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetEntryByOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/fds/env/entry",
		Infix:  "/oatuh",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "第三方登录",
		En:    "EntryByOAuth",
		Func:  fs_function_tags.GetEntryByOAuthFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetEntryByValidateCodeFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/entry",
		Infix:  "/vc",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "验证码登录",
		En:    "EntryByValidateCode",
		Func:  fs_function_tags.GetEntryByValidateCodeFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}

func GetEntryByQRCodeFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/entry",
		Infix:  "/qrcode",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "二维码登录",
		En:    "EntryByQRCode",
		Func:  fs_function_tags.GetEntryByQRCodeFuncTag(),
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 1,
	}
	return f
}
