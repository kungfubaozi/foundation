package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetFromOAuthFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/register",
		Infix:  "/oauth",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "第三方账号注册",
		En:    "RegisterFromOAuth",
		Fcv:   fs_constants.FCV_VALIDATE_CODE,
		Func:  fs_function_tags.GetFromOAuthFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}

func GetFromAPFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/entry/register",
		Infix:  "/ap",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "账号注册",
		En:    "RegisterFromAP",
		Fcv:   fs_constants.FCV_VALIDATE_CODE,
		Func:  fs_function_tags.GetFromAPFuncTag(),
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_TOURIST,
	}
	return f
}
