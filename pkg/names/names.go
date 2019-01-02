package names

const (
	F_SVC_MESSAGE             = "foundation.svc.message"
	F_SVC_STATE               = "foundation.svc.state"
	F_SVC_VALIDATE            = "foundation.svc.validate"
	F_SVC_REPORTER            = "foundation.svc.reporter"
	F_SVC_USER                = "foundation.svc.user"
	F_SVC_SAFETY_UPDATE       = "foundation.svc.safety.update"
	F_SVC_SAFETY_VERIFICATION = "foundation.svc.safety.verification"
)

const (
	F_DO_LOGIN             int64 = 10101
	F_DO_LOGIN_VER         int64 = 10101
	F_DO_REGISTER          int64 = 10201
	F_DO_UPDATE_PASSWORD   int64 = 10301
	F_DO_UPDATE_PHONE      int64 = 10302
	F_DO_UPDATE_EMAIL      int64 = 10303
	F_DO_UPDATE_ENTERPRISE int64 = 10304
)

const (
	F_STATE_OK int64 = 3 << 1
)

const (
	F_PLATFORM_ANDROID = 4 << 1
	F_PLATFORM_IOS     = 4 << 2
	F_PLATFORM_WINDOWD = 4 << 3
	F_PLATFORM_MAC_OS  = 4 << 4
	F_PLATFORM_WEB     = 4 << 5
)
