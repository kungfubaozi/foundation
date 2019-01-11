package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetAddFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/function",
		Infix:  "/add",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "添加功能",
		En:    "AddFunction",
		Func:  "9dfb23d9db2c",
		Fcv:   names.F_FCV_AUTH,
		Level: 4,
	}
	return f
}

func GetRemoveFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/function",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "移除功能",
		En:    "RemoveFunction",
		Func:  "713fd38bf13d",
		Fcv:   names.F_FCV_AUTH,
		Level: 4,
	}
	return f
}

func GetUpdateFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/function",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新功能",
		En:    "UpdateFunction",
		Func:  "3a2a43a79f2e",
		Fcv:   names.F_FCV_AUTH,
		Level: 4,
	}
	return f
}

func GetAllFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/function",
		Infix:  "/all",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "获取所有功能",
		En:    "AllFunction",
		Func:  "1b84740cec7e",
		Fcv:   names.F_FCV_AUTH,
		Level: 4,
	}
	return f
}

func GetFindFunc() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/function",
		Infix:  "/find",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "查找功能",
		En:    "FindFunction",
		Func:  "c8cefd194d28",
		Fcv:   names.F_FCV_AUTH,
		Level: 4,
	}
	return f
}
