package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetInviteUserFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/invite",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "邀请用户",
		En:    "InviteUser",
		Func:  "72777f71a26d",
		Level: fs_constants.LEVEL_ADMIN,
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
	}
	return f
}
