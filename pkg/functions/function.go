package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetAddFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "添加功能",
		En:    "AddFunction",
		Func:  "9dfb23d9db2c",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetRemoveFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "移除功能",
		En:    "RemoveFunction",
		Func:  "713fd38bf13d",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetUpdateFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新功能",
		En:    "UpdateFunction",
		Func:  "3a2a43a79f2e",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetAllFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/all",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "获取所有功能",
		En:    "AllFunction",
		Func:  "1b84740cec7e",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetFindFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/find",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "查找功能",
		En:    "FindFunction",
		Func:  "c8cefd194d28",
		Fcv:   fs_constants.FCV_AUTH,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}
