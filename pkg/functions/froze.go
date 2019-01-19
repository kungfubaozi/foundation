package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetRequestFrozeFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/safety/froze",
		Infix:  "/lock",
		Function: &fs_base_function.Func{
			Zh:    "冻结账户",
			En:    "FrozeAccount",
			Func:  "e66f94e8b5cc",
			Fcv:   fs_constants.FCV_NONE,
			Level: fs_constants.LEVEL_TOURIST,
		},
	}
}
