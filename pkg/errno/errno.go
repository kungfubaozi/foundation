package errno

import (
	"errors"
	"zskparker.com/foundation/base/pb"
)

func ErrResponse(state *fs_base.State) (*fs_base.Response, error) {
	return &fs_base.Response{State: state}, nil
}

var (
	ERROR = errors.New("err")
)

var (
	Ok                 = &fs_base.State{Ok: true, Message: "ok", Code: 0}
	ErrRequest         = &fs_base.State{Ok: false, Message: "err request", Code: 10001}
	ErrSystem          = &fs_base.State{Ok: false, Message: "err system", Code: 10002}
	ErrBusy            = &fs_base.State{Ok: false, Message: "err busy", Code: 10003}
	ErrSupport         = &fs_base.State{Ok: false, Message: "err support", Code: 10004}
	ErrInvalid         = &fs_base.State{Ok: false, Message: "err invalid", Code: 10005}
	ErrExpired         = &fs_base.State{Ok: false, Message: "err expired", Code: 10006}
	ErrFaceCompare     = &fs_base.State{Ok: false, Message: "err face compare", Code: 10007}
	ErrFaceProcess     = &fs_base.State{Ok: false, Message: "err face process", Code: 10008}
	ErrMetaValidate    = &fs_base.State{Ok: false, Message: "err validate meta", Code: 10009}
	ErrFroze           = &fs_base.State{Ok: false, Message: "err user froze", Code: 10010}
	ErrData            = &fs_base.State{Ok: false, Message: "err data", Code: 10011}
	ErrPlatform        = &fs_base.State{Ok: false, Message: "err platform", Code: 10012}
	ErrPhoneValidate   = &fs_base.State{Ok: false, Message: "err phone validate", Code: 10013}
	ErrEmailValidate   = &fs_base.State{Ok: false, Message: "err email validate", Code: 10014}
	ErrFaceValidate    = &fs_base.State{Ok: false, Message: "err face validate", Code: 10015}
	ErrToken           = &fs_base.State{Ok: false, Message: "err token", Code: 10016}
	ErrTokenExpired    = &fs_base.State{Ok: false, Message: "err token expired", Code: 10017}
	ErrUserState       = &fs_base.State{Ok: false, Message: "err user state", Code: 10018}
	ErrProjectAccess   = &fs_base.State{Ok: false, Message: "err project access", Code: 10020}
	ErrFunction        = &fs_base.State{Ok: false, Message: "err function", Code: 10021}
	ErrFunctionInvalid = &fs_base.State{Ok: false, Message: "err function invalid", Code: 10022}
)
