package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetInviteUserFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/invite",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "邀请用户",
		En:    "InviteUser",
		Func:  "72777f71a26d",
		Level: 5,
		Fcv:   names.F_FCV_AUTH | names.F_FCV_FACE,
	}
	return f
}
