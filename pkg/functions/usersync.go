package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetAddUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/sync",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "添加用户同步",
		En:    "AddUserSyncHook",
		Func:  "6fce56f355b7",
		Type:  fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetRemoveUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/sync",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Zh:   "移除用户同步",
		En:   "RemoveUserSyncHook",
		Func: "4b99224bc064", Type: fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetUpdateUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/sync",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Zh:   "更新用户同步",
		En:   "UpdateUserSyncHook",
		Func: "f716c335dbf3", Type: fs_constants.TYPE_HIDE,
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}
