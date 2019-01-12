package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetUpdateProjectStrategyFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/strategy",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新策略",
		En:    "UpdateStrategy",
		Func:  "7551e6a14638",
		Level: fs_constants.LEVEL_ADMIN,
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
	}
	return f
}
