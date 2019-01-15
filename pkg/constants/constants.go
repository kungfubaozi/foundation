package fs_constants

const (
	USER_TYPE_FROM_INVITE   = 5 << 10 //来自邀请
	USER_TYPE_FROM_REGISTER = 5 << 13 //来自注册
)

const (
	LEVEL_TOURIST         = 8 << 3 //游客
	LEVEL_USER            = 8 << 4 //用户
	LEVEL_DEVELOPER       = 8 << 5 //开发者
	LEVEL_PROJECT_MANAGER = 8 << 6 //项目管理员
	LEVEL_ADMIN           = 8 << 7 //最高管理员
)

const (
	TYPE_HIDE = 2 << 3
)

const (
	PLATFORM_ANDROID = 4 << 1
	PLATFORM_IOS     = 4 << 2
	PLATFORM_WINDOWD = 4 << 3
	PLATFORM_MAC_OS  = 4 << 4
	PLATFORM_WEB     = 4 << 5
)

const (
	FCV_FACE          = 6 << 4
	FCV_VALIDATE_CODE = 6 << 7
	FCV_AUTH          = 6 << 9
	FCV_NONE          = 6 << 10
	FCV_PASSWORD      = 6 << 11
)

const (
	STATE_OK                  int64 = 3 << 1
	STATE_USER_FROZE          int64 = 3 << 2 //用户状态冻结
	STATE_USER_LOCK           int64 = 3 << 3 //用户被锁定
	STATE_USER_VALIDATE       int64 = 3 << 6 //需要通过手机验证才可继续
	STATE_USER_RESET_PASSWORD int64 = 3 << 8 //需要通过手机验证才可继续
)

const (
	SVC_MESSAGE             = "fd.svc.message"
	SVC_STATE               = "fd.svc.state"
	SVC_VALIDATE            = "fd.svc.validate"
	SVC_REPORTER            = "fd.svc.reporter"
	SVC_USER                = "fd.svc.user"
	SVC_SAFETY_UPDATE       = "fd.svc.safety.update"
	SVC_SAFETY_VERIFICATION = "fd.svc.safety.verification"
	SVC_AUTHENTICATE        = "fd.svc.authenticate"
	SVC_FUNCTION            = "fd.svc.function"
	SVC_FACE                = "fd.svc.face"
	SVC_STRATEGY            = "fd.svc.strategy"
	SVC_PROJECT             = "fd.svc.project"
	SVC_ENTRY_REGISTER      = "fd.svc.entry.register"
	SVC_ENTRY_LOGIN         = "fd.svc.entry.login"
	SVC_INTERCEPTOR         = "fd.svc.interceptor"
	SVC_REFRESH             = "fd.svc.refresh"
	SVC_INVITE              = "fd.svc.invite"
)
