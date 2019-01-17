package fs_functions

import (
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/tags"
)

func GetCreateProject() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/create",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "新建项目",
		En:    "CreateProject",
		Func:  fs_function_tags.GetCreateProjectTag(),
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
		Func:  fs_function_tags.GetRemoveProjectTag(),
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
		Func:  fs_function_tags.GetUpdateProjectTag(),
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}

func GetEnablePlatform() *fs_pkg_model.APIFunction {
	f := &fs_pkg_model.APIFunction{
		Prefix: "/fds/api/env/project",
		Infix:  "/enablePlatform",
	}
	f.Function = &fs_base_function.Func{
		Zh:    "开启或关闭项目的终端",
		En:    "EnablePlatform",
		Func:  fs_function_tags.GetUpdateProjectTag(),
		Fcv:   fs_constants.FCV_AUTH | fs_constants.FCV_FACE,
		Type:  fs_constants.TYPE_HIDE,
		Level: fs_constants.LEVEL_PROJECT_MANAGER,
	}
	return f
}
