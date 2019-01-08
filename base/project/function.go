package project

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/names"
)

func GetCreateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/create",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "新建项目",
		En:    "CreateProject",
		Func:  "3c7cec044485",
		Fcv:   names.F_FCV_AUTH | names.F_FCV_FACE,
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 4,
	}
	return f
}

func GetRemoveProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "移除项目",
		En:    "RemoveProject",
		Func:  "4974f1c3a33b",
		Fcv:   names.F_FCV_AUTH | names.F_FCV_FACE,
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 4,
	}
	return f
}

func GetUpdateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Api:   f.Prefix + f.Infix,
		Zh:    "更新项目",
		En:    "UpdateProject",
		Func:  "5451dbc0b529",
		Fcv:   names.F_FCV_AUTH | names.F_FCV_FACE,
		Type:  names.F_FUNC_TYPE_HIDE,
		Level: 4,
	}
	return f
}
