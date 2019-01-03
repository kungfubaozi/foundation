package verification

import (
	"context"
	"fmt"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy/def"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
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
	fmt.Println("new verification")
	if in.Do <= 0 && in.FuncId <= 0 {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrRequest,
		}, nil
	}
	req := &fs_base_validate.CreateRequest{}
	meta := &fs_base.Metadata{}
	meta.Ip = "localhost"

	if ctx.Value("meta") != nil {
		meta = ctx.Value("meta").(*fs_base.Metadata)
	}
	strategy := &fs_base.ProjectStrategy{}
	if ctx.Value("strategy") != nil {
		strategy = ctx.Value("strategy").(*fs_base.ProjectStrategy)
		req.OnVerification = strategy.Events.OnVerification
	} else {
		req.OnVerification = strategydef.GetOnVerificationDefault()
	}
	req.To = in.To
	req.Do = in.Do
	req.Metadata = meta
	if meta != nil { //鉴权信息不为空的情况下

	}
	resp, _ := svc.validatecli.Create(ctx, req)
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
