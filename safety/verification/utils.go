package verification

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
)

func FromRequestMeta(validatecli validate.Service, in *fs_base.Meta, do int64) *fs_base.State {
	if len(in.Validate) == 0 || len(in.Id) == 0 {
		return errno.ErrMetaValidate
	}
	resp, err := validatecli.Verification(context.Background(), &fs_base_validate.VerificationRequest{
		VerId: in.Id,
		Code:  in.Validate,
		Do:    do,
	})
	if err != nil {
		return errno.ErrSystem
	}
	return resp.State
}
