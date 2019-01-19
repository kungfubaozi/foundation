package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetAddFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/add",
		Function: &fs_base_function.Func{
			Zh:    "添加功能",
			En:    "AddFunction",
			Func:  "9dfb23d9db2c",
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_PROJECT_MANAGER,
		},
	}
}

func GetRemoveFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/remove",
		Function: &fs_base_function.Func{
			Zh:    "移除功能",
			En:    "RemoveFunction",
			Func:  "713fd38bf13d",
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_PROJECT_MANAGER,
		},
	}
}

func GetUpdateFunc() *fs_pkg_model.APIFunction {

	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/update",
		Function: &fs_base_function.Func{
			Zh:    "更新功能",
			En:    "UpdateFunction",
			Func:  "3a2a43a79f2e",
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_PROJECT_MANAGER,
		},
	}
}

func GetAllFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/all",
		Function: &fs_base_function.Func{
			Zh:    "获取所有功能",
			En:    "AllFunction",
			Func:  "1b84740cec7e",
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_PROJECT_MANAGER,
		},
	}
}

func GetFindFunc() *fs_pkg_model.APIFunction {
	return &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/function",
		Infix:  "/find",
		Function: &fs_base_function.Func{
			Zh:    "查找功能",
			En:    "FindFunction",
			Func:  "c8cefd194d28",
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_PROJECT_MANAGER,
		},
	}
}
