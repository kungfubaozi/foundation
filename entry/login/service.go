package login

import (
	"context"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/entry/login/pb"
)

type Service interface {
	EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_base.Response, error)

	EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_base.Response, error)

	EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_base.Response, error)

	EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_base.Response, error)
}

type loginService struct {
}
