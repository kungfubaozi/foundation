package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetAddUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/sync",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "添加用户同步",
		En:    "AddUserSyncHook",
		Func:  "6fce56f355b7",
		Level: 1,
	}
	return f
}

func GetRemoveUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/sync",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "移除用户同步",
		En:    "RemoveUserSyncHook",
		Func:  "4b99224bc064",
		Level: 1,
	}
	return f
}

func GetUpdateUserSyncHookFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/sync",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新用户同步",
		En:    "UpdateUserSyncHook",
		Func:  "f716c335dbf3",
		Level: 1,
	}
	return f
}
