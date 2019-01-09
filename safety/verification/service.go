package verification

import (
	"context"
	"github.com/go-kit/kit/log"
	"regexp"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
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
	if len(in.Func) == 0 {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrRequest,
		}, nil
	}
	svc.logger.Log("verification", "new", "func", in.Func)
	req := &fs_base_validate.CreateRequest{}
	meta := fs_metadata_transport.ContextToMeta(ctx)
	strategy := fs_metadata_transport.ContextToStrategy(ctx)
	switch in.Func {
	case fs_function_tags.GetAdminFuncTag(), fs_function_tags.GetFromAPFuncTag(): //注册
		req.Mode = strategy.Events.OnRegister.Mode //获取注册模式
		svc.logger.Log("function", "register")
		break
	case fs_function_tags.GetEntryByValidateCodeFuncTag(): //使用验证码登录
		svc.logger.Log("function", "entry by validate code")
		break

	default: //未匹配的方法
		svc.logger.Log("function", "no mactch")
		return &fs_safety_verification.NewResponse{State: errno.ErrRequest}, nil
	}
	if req.Mode == 1 { //手机号验证
		reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(in.To) {
			return &fs_safety_verification.NewResponse{State: errno.ErrPhoneNumber}, nil
		}
		svc.logger.Log("register", "phone")
	} else if req.Mode == 2 { //邮箱验证
		reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(in.To) {
			return &fs_safety_verification.NewResponse{State: errno.ErrEmailFormat}, nil
		}
		svc.logger.Log("register", "email")
	} else {
		svc.logger.Log("register", "unsupported")
		return &fs_safety_verification.NewResponse{State: errno.ErrSystem}, nil
	}
	req.OnVerification = strategy.Events.OnVerification
	req.To = in.To
	req.Func = in.Func
	req.Metadata = meta
	resp, err := svc.validatecli.Create(context.Background(), req)
	if err != nil {
		svc.logger.Log("validate create", "err", "info", err)
		return &fs_safety_verification.NewResponse{
			State: errno.ErrSystem,
		}, nil
	}
	svc.logger.Log("validate create", "ok", "info", resp)
	return &fs_safety_verification.NewResponse{
		State: resp.State,
		VerId: resp.VerId,
	}, nil
}

func NewService(validatecli validate.Service, reportercli reportercli.Channel, logger log.Logger) Service {
	var svc Service
	{
		svc = &verificationService{validatecli: validatecli, reportercli: reportercli, logger: logger}
	}
	return svc
}
