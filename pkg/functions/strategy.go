package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetUpdateProjectStrategyFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/strategy",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新策略",
		En:    "UpdateStrategy",
		Func:  "7551e6a14638",
		Level: 5,
		Fcv:   names.F_FCV_AUTH | names.F_FCV_FACE,
		Type:  names.F_FUNC_TYPE_HIDE,
	}
	return f
}
