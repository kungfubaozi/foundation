package login

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetEntryByFaceFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/entry",
		Infix:  "/face",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "刷脸登录",
		En:    "EntryByFace",
		Func:  "547ad4cc8ddc",
		Level: 5, //刷脸登录目前只对最高管理员开放
	}
	return f
}

func GetEntryByAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/entry",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "密码登录",
		En:    "EntryByAP",
		Func:  "8423a0ce0eb3",
		Level: 1,
	}
	return f
}

func GetEntryByOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/entry",
		Infix:  "/oatuh",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "第三方登录",
		En:    "EntryByOAuth",
		Func:  "3a8a4ba0f66a",
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
		Func:  "bf7a04fcc618",
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
		Func:  "84f9cb3a0619",
		Level: 1,
	}
	return f
}
