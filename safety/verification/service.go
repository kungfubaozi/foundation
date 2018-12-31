package verification

import (
	"context"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/validate"
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

func (verificationService) New(ctx context.Context, in *fs_safety_verification.NewRequest) (*fs_safety_verification.NewResponse, error) {
	panic("implement me")
}

func (verificationService) GetFuncVerMode(ctx context.Context, in *fs_safety_verification.GetFuncVerModeRequest) (*fs_safety_verification.GetFuncVerModeResponse, error) {
	panic("implement me")
}

func NewService(validatecli validate.Service, functioncli function.Service, reportercli reporter.Service) Service {
	var svc Service
	{
		svc = &verificationService{validatecli: validatecli, functioncli: functioncli, reportercli: reportercli}
	}
	return svc
}
