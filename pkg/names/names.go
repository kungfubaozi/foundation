package names

const (
	F_SVC_MESSAGE             = "foundation.svc.message"
	F_SVC_STATE               = "foundation.svc.state"
	F_SVC_VALIDATE            = "foundation.svc.validate"
	F_SVC_REPORTER            = "foundation.svc.reporter"
	F_SVC_USER                = "foundation.svc.user"
	F_SVC_SAFETY_UPDATE       = "foundation.svc.safety.update"
	F_SVC_SAFETY_VERIFICATION = "foundation.svc.safety.verification"
	F_SVC_AUTHENTICATE        = "foundation.svc.authenticate"
	F_SVC_FUNCTION            = "foundation.svc.function"
	F_SVC_FACE                = "foundation.svc.face"
	F_SVC_STRATEGY            = "foundation.svc.strategy"
	F_SVC_PROJECT             = "foundation.svc.project"
	F_SVC_ENTRY_REGISTER      = "foundation.svc.entry.register"
	F_SVC_ENTRY_LOGIN         = "foundation.svc.entry.login"
)

const (
	F_STATE_OK         int64 = 3 << 1
	F_USER_STATE_FROZE int64 = 3 << 2 //用户状态冻结
	F_USER_STATE_LOCK  int64 = 3 << 3 //用户被锁定
	F_USER_VALIDATE    int64 = 3 << 3 //需要通过手机验证才可继续
)

const (
	F_FUNC_TYPE_HIDE int64 = 8 << 2
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

const (
	F_FCV_FACE          = 6 << 6
	F_FCV_VALIDATE_CODE = 6 << 7
	F_FCV_AUTH          = 6 << 9
	F_FCV_NONE          = 6 << 10
)
