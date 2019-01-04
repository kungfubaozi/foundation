package project

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
)

func GetCreateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/project",
		Infix:  "/create",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建项目",
		En:    "CreateProject",
		Func:  "3c7cec044485",
		Level: 1,
	}
	return f
}

func GetRemoveProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/project",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "移除项目",
		En:    "RemoveProject",
		Func:  "4974f1c3a33b",
		Level: 1,
	}
	return f
}

func GetUpdateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/api/fds/env/project",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新项目",
		En:    "UpdateProject",
		Func:  "5451dbc0b529",
		Level: 1,
	}
	return f
}
