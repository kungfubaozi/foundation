package fs_pkg_model

import "zskparker.com/foundation/base/function/pb"

type APIFunction struct {
	Prefix   string
	Infix    string
	Function *fs_base_function.Func
}
