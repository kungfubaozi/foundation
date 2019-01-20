package verification

import (
	"context"
	"github.com/go-kit/kit/log"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/verification/pb"
)

type Service interface {
	New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error)
}

type verificationService struct {
	validatecli validate.Service
	reportercli reportercli.Channel
	logger      log.Logger
}

func (svc *verificationService) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	resp := func(state *fs_base.State) (*fs_safety_verification.NewResponse, error) {
		return &fs_safety_verification.NewResponse{State: state}, nil
	}

	if len(in.Func) == 0 {
		return resp(errno.ErrRequest)
	}
	svc.logger.Log("verification", "new", "func", in.Func)
	req := &fs_base_validate.CreateRequest{}
	meta := fs_metadata_transport.ContextToMeta(ctx)
	strategy := fs_metadata_transport.ContextToStrategy(ctx)

	//加入锁机制，内部避免再次经过数据库查询判断间隔时间
	//defer svc.redisync.Unlock("173e2bb3f601", meta.Ip)

	switch in.Func {
	case fs_function_tags.GetAdminFuncTag(), fs_function_tags.GetFromAPFuncTag(): //注册
		req.Mode = strategy.Events.OnRegister.Mode //获取注册模式
		svc.logger.Log("function", "register")
		break
	case fs_function_tags.GetEntryByValidateCodeFuncTag(): //使用验证码登录
		svc.logger.Log("function", "entry by validate code")
		break
	case fs_function_tags.GetResetPasswordFuncTag():
		req.Mode = 1
		break

	default: //未匹配的方法
		svc.logger.Log("function", "no mactch")
		return resp(errno.ErrRequest)
	}
	if req.Mode == 1 { //手机号验证
		if !fs_regx_match.Phone(in.To) {
			return resp(errno.ErrPhoneNumber)
		}
		svc.logger.Log("register", "phone")
	} else if req.Mode == 2 { //邮箱验证
		if !fs_regx_match.Email(in.To) {
			return resp(errno.ErrEmailFormat)
		}
		svc.logger.Log("register", "email")
	} else {
		svc.logger.Log("register", "unsupported")
		return resp(errno.ErrSystem)
	}
	req.To = in.To
	req.Func = in.Func
	req.Metadata = meta
	req.OnVerification = strategy.Events.OnVerification

	vr, err := svc.validatecli.Create(context.Background(), req)
	if err != nil {
		svc.logger.Log("validate create", "err", "info", err)
		return resp(errno.ErrSystem)
	}

	svc.logger.Log("validate create", "ok", "info", resp)
	return &fs_safety_verification.NewResponse{
		State: vr.State,
		VerId: vr.VerId,
	}, nil
}

func NewService(validatecli validate.Service, reportercli reportercli.Channel, logger log.Logger) Service {
	var svc Service
	{
		svc = &verificationService{validatecli: validatecli, reportercli: reportercli, logger: logger}
	}
	return svc
}
