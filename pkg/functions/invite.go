package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetInviteUserFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/invite",
		Infix:  "/add",
		Function: &fs_base_function.Func{
			Zh:    "邀请用户",
			En:    "InviteUser",
			Func:  fs_function_tags.GetInviteTag(),
			Level: fs_constants.LEVEL_ADMIN,
			Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		},
	}
}
