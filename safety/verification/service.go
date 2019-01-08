package verification

import (
	"context"
	"fmt"
	"regexp"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/safety/verification/pb"
)

type Service interface {
	New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error)
}

type verificationService struct {
	validatecli validate.Service
	reportercli reportercli.Channel
}

func (svc *verificationService) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	if len(in.Func) == 0 {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrRequest,
		}, nil
	}
	fmt.Println("new")
	req := &fs_base_validate.CreateRequest{}
	meta := ctx.Value("meta").(*fs_base.Metadata)
	strategy := ctx.Value("strategy").(*fs_base.ProjectStrategy)
	switch in.Func {
	case fs_function_tags.GetFromAPFuncTag():
	case fs_function_tags.GetAdminFuncTag(): //注册
		req.Mode = strategy.Events.OnRegister.Mode //获取注册模式
		break
	case fs_function_tags.GetEntryByValidateCodeFuncTag(): //使用验证码登录
		break

	default: //未匹配的方法
		return &fs_safety_verification.NewResponse{State: errno.ErrRequest}, nil
	}
	if req.Mode == 1 { //手机号验证
		reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(in.To) {
			return &fs_safety_verification.NewResponse{State: errno.ErrPhoneNumber}, nil
		}
		fmt.Println("phone")
	} else if req.Mode == 2 { //邮箱验证
		reg := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		rgx := regexp.MustCompile(reg)
		if !rgx.MatchString(in.To) {
			return &fs_safety_verification.NewResponse{State: errno.ErrEmailFormat}, nil
		}
		fmt.Println("email")
	} else {
		return &fs_safety_verification.NewResponse{State: errno.ErrSystem}, nil
	}
	req.OnVerification = strategy.Events.OnVerification
	req.To = in.To
	req.Func = in.Func
	req.Metadata = meta
	resp, err := svc.validatecli.Create(context.Background(), req)
	if err != nil {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrSystem,
		}, nil
	}
	return &fs_safety_verification.NewResponse{
		State: resp.State,
		VerId: resp.VerId,
	}, nil
}

func NewService(validatecli validate.Service, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &verificationService{validatecli: validatecli, reportercli: reportercli}
	}
	return svc
}
