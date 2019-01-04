package login

import (
	"context"
	"time"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/entry/login/pb"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error)

	EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_entry_login.EntryResponse, error)

	EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_entry_login.EntryResponse, error)

	EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_entry_login.EntryResponse, error)

	EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error)
}

type loginService struct {
	usercli         user.Service
	reportercli     reportercli.Channel
	authenticatecli authenticate.Service
	validatecli     validate.Service
	facecli         face.Service
}

func (svc *loginService) EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error) {
	meta := ctx.Value("meta").(*fs_base.Metadata)
	//strategy := ctx.Value("strategy").(*fs_base.ProjectStrategy)
	//查找人脸库
	fr, _ := svc.facecli.Search(context.Background(), &fs_base_face.SearchRequest{
		Base64Face: in.Meta.Face,
	})
	if !fr.State.Ok {
		return &fs_entry_login.EntryResponse{State: fr.State}, nil
	}
	ar, _ := svc.authenticatecli.New(context.Background(), &fs_base_authenticate.NewRequest{
		Authorize: &fs_base_authenticate.Authorize{
			UserId:    fr.UserId,
			Timestamp: time.Now().UnixNano(),
			ProjectId: meta.ProjectId,
			ClientId:  meta.ClientId,
			Ip:        meta.Ip,
			Level:     fr.Level,
			Device:    meta.Device,
			UserAgent: meta.UserAgent,
		},
	})
	if !ar.State.Ok {
		return &fs_entry_login.EntryResponse{State: ar.State}, nil
	}
	//svc.reportercli.Write(login, in.Authorize.UserId, in.Authorize.Ip)
	return &fs_entry_login.EntryResponse{
		State:        errno.Ok,
		Session:      ar.Session,
		RefreshToken: ar.RefreshToken,
		AccessToken:  ar.AccessToken,
	}, nil
}

func (svc *loginService) EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error) {
	panic("implement me")
}

func (svc *loginService) EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_entry_login.EntryResponse, error) {
	panic("implement me")
}

func (svc *loginService) EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_entry_login.EntryResponse, error) {
	panic("implement me")
}

func (svc *loginService) EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_entry_login.EntryResponse, error) {
	panic("implement me")
}

func NewService(usercli user.Service, reportercli reportercli.Channel, authenticatecli authenticate.Service,
	validatecli validate.Service, facecli face.Service) Service {
	var service Service
	{
		service = &loginService{usercli: usercli, reportercli: reportercli,
			authenticatecli: authenticatecli, validatecli: validatecli, facecli: facecli}
	}
	return service
}
