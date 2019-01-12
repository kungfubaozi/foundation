package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
)

func GetCreateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/create",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建项目",
		En:    "CreateProject",
		Func:  "3c7cec044485",
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetRemoveProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/remove",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "移除项目",
		En:    "RemoveProject",
		Func:  "4974f1c3a33b",
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetUpdateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/update",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "更新项目",
		En:    "UpdateProject",
		Func:  "5451dbc0b529",
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}
