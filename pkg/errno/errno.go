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
	Ok                      = &fs_base.State{Ok: true, Message: "ok", Code: 0}
	ErrRequest              = &fs_base.State{Ok: false, Message: "err request", Code: 10001}
	ErrSystem               = &fs_base.State{Ok: false, Message: "err system", Code: 10002}
	ErrBusy                 = &fs_base.State{Ok: false, Message: "err busy", Code: 10003}
	ErrSupport              = &fs_base.State{Ok: false, Message: "err support", Code: 10004}
	ErrInvalid              = &fs_base.State{Ok: false, Message: "err invalid", Code: 10005}
	ErrExpired              = &fs_base.State{Ok: false, Message: "err expired", Code: 10006}
	ErrFaceCompare          = &fs_base.State{Ok: false, Message: "err face compare", Code: 10007}
	ErrFaceProcess          = &fs_base.State{Ok: false, Message: "err face process", Code: 10008}
	ErrMetaValidate         = &fs_base.State{Ok: false, Message: "err validate meta", Code: 10009}
	ErrFroze                = &fs_base.State{Ok: false, Message: "err user froze", Code: 10010}
	ErrData                 = &fs_base.State{Ok: false, Message: "err data", Code: 10011}
	ErrPlatform             = &fs_base.State{Ok: false, Message: "err platform", Code: 10012}
	ErrPhoneValidate        = &fs_base.State{Ok: false, Message: "err phone validate", Code: 10013}
	ErrEmailValidate        = &fs_base.State{Ok: false, Message: "err email validate", Code: 10014}
	ErrFaceValidate         = &fs_base.State{Ok: false, Message: "err face validate", Code: 10015}
	ErrToken                = &fs_base.State{Ok: false, Message: "err token", Code: 10016}
	ErrTokenExpired         = &fs_base.State{Ok: false, Message: "err token expired", Code: 10017}
	ErrUserState            = &fs_base.State{Ok: false, Message: "err user state", Code: 10018}
	ErrProjectAccess        = &fs_base.State{Ok: false, Message: "err project access", Code: 10020}
	ErrFunction             = &fs_base.State{Ok: false, Message: "err function", Code: 10021}
	ErrFunctionInvalid      = &fs_base.State{Ok: false, Message: "err function invalid", Code: 10022}
	ErrInvalidFace          = &fs_base.State{Ok: false, Message: "err invalid face", Code: 10023}  //未发现face
	ErrPhoneNumber          = &fs_base.State{Ok: false, Message: "err phone number", Code: 10024}  //错误手机好吗
	ErrEmailFormat          = &fs_base.State{Ok: false, Message: "err email format", Code: 10025}  //错误邮箱地址
	ErrValidateCode         = &fs_base.State{Ok: false, Message: "err validate code", Code: 10026} //验证码错误
	ErrTransfer             = &fs_base.State{Ok: false, Message: "err system data transfer", Code: 10027}
	ErrAlreadyBind          = &fs_base.State{Ok: false, Message: "err already bind", Code: 10028}
	ErrEmail                = &fs_base.State{Ok: false, Message: "err email", Code: 10029}
	ErrUserAlreadyExists    = &fs_base.State{Ok: false, Message: "err user already exists", Code: 10030}
	ErrProjectPermission    = &fs_base.State{Ok: false, Message: "err project permission", Code: 10031} //当前用户没有权限访问当前项目
	ErrProject              = &fs_base.State{Ok: false, Message: "err project", Code: 10032}
	ErrInvalidOrAccount     = &fs_base.State{Ok: false, Message: "err invalid account or password", Code: 10033}
	ErrRequestPermission    = &fs_base.State{Ok: false, Message: "err request permission", Code: 10034}
	ErrProjectAlreadyExists = &fs_base.State{Ok: false, Message: "err project exists", Code: 10035}
	ErrClient               = &fs_base.State{Ok: false, Message: "err client", Code: 10036}
	ErrInviteAccount        = &fs_base.State{Ok: false, Message: "err invite user account", Code: 10037}
	ErrInviteUsername       = &fs_base.State{Ok: false, Message: "err invite user username", Code: 10038}
	ErrInviteExists         = &fs_base.State{Ok: false, Message: "err invite user already exits", Code: 10039}
	ErrInviteLevel          = &fs_base.State{Ok: false, Message: "err invite user level", Code: 10040}
	ErrInvalidInviteKey     = &fs_base.State{Ok: false, Message: "err invalid invite code", Code: 10041}
	ErrAlreadyInvited       = &fs_base.State{Ok: false, Message: "err already invited", Code: 10042}
	ErrResetPassword        = &fs_base.State{Ok: false, Message: "err reset password", Code: 10043}
	ErrAlreadyExists        = &fs_base.State{Ok: false, Message: "err already exists", Code: 10044}
)
