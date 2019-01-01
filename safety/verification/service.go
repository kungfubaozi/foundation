package verification

import (
	"context"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/safety/verification/pb"
)

type Service interface {
	New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error)

	GetFuncVerMode(ctx context.Context, in *fs_safety_verification.GetFuncVerModeRequest) (*fs_safety_verification.GetFuncVerModeResponse, error)
}

type verificationService struct {
	validatecli validate.Service
	functioncli function.Service
	reportercli reporter.Service
}

func (svc *verificationService) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	if in.Do <= 0 && in.FuncId <= 0 {
		return &fs_safety_verification.NewResponse{
			State: errno.ErrRequest,
		}, nil
	}
	resp, err := svc.validatecli.Create(ctx, &fs_base_validate.CreateRequest{})
}

func (svc *verificationService) GetFuncVerMode(ctx context.Context, in *fs_safety_verification.GetFuncVerModeRequest) (*fs_safety_verification.GetFuncVerModeResponse, error) {
	panic("implement me")
}

func NewService(validatecli validate.Service, functioncli function.Service, reportercli reporter.Service) Service {
	var svc Service
	{
		svc = &verificationService{validatecli: validatecli, functioncli: functioncli, reportercli: reportercli}
	}
	return svc
}
