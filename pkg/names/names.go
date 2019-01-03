package names

const (
	F_SVC_MESSAGE             = "foundation.svc.message"
	F_SVC_STATE               = "foundation.svc.state"
	F_SVC_VALIDATE            = "foundation.svc.validate"
	F_SVC_REPORTER            = "foundation.svc.reporter"
	F_SVC_USER                = "foundation.svc.user"
	F_SVC_SAFETY_UPDATE       = "foundation.svc.safety.update"
	F_SVC_SAFETY_VERIFICATION = "foundation.svc.safety.verification"
	F_SVC_SAFETY_AUTHENTICATE = "foundation.svc.safety.authenticate"
)

const (
	F_FUNC_LOGIN             string = "8f537f6a"
	F_FUNC_LOGIN_VER         string = "d2d0371b"
	F_FUNC_REGISTER          string = "f83152ac"
	F_FUNC_UPDATE_PASSWORD   string = "b60ed019"
	F_FUNC_UPDATE_PHONE      string = "cd994e0f"
	F_FUNC_UPDATE_EMAIL      string = "ecf8ecc2"
	F_FUNC_UPDATE_ENTERPRISE string = "4fea4bba"
)

const (
	F_STATE_OK         int64 = 3 << 1
	F_USER_STATE_FROZE int64 = 3 << 2 //用户状态冻结
	F_USER_STATE_LOCK  int64 = 3 << 3 //用户被锁定
	F_USER_VALIDATE    int64 = 3 << 3 //需要通过手机验证才可继续
)

const (
	F_PLATFORM_ANDROID = 4 << 1
	F_PLATFORM_IOS     = 4 << 2
	F_PLATFORM_WINDOWD = 4 << 3
	F_PLATFORM_MAC_OS  = 4 << 4
	F_PLATFORM_WEB     = 4 << 5
)

const (
	F_VALIDATE_FACE  = 5 << 1
	F_VALIDATE_EMAIL = 5 << 2
	F_VALIDATE_PHONE = 5 << 3
)
