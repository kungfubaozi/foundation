package verification

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
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
	if len(in.Func) == 0 {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrRequest,
		}, nil
	}
	req := &fs_base_validate.CreateRequest{}
	meta := ctx.Value("meta").(*fs_base.Metadata)
	strategy := ctx.Value("strategy").(*fs_base.ProjectStrategy)
	req.OnVerification = strategy.Events.OnVerification
	req.To = in.To
	req.Func = in.Func
	req.Metadata = meta
	resp, _ := svc.validatecli.Create(context.Background(), req)
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
