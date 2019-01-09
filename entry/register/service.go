package register

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
)

type Service interface {
	FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error)

	FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error)

	Admin(ctx context.Context, in *fs_entry_register.AdminRequest) (*fs_base.Response, error)
}

type registerService struct {
	usercli     user.Service
	facecli     face.Service
	reportercli reportercli.Channel
	session     *mgo.Session
}

func (svc *registerService) GetRepo() repository {
	return &registerRepository{session: svc.session.Clone()}
}

//level:  1:游客 2:用户 3:开发人员  4:应用管理员  5:系统管理员
func (svc *registerService) Admin(ctx context.Context, in *fs_entry_register.AdminRequest) (*fs_base.Response, error) {
	if len(in.Meta.Face) == 0 {
		return errno.ErrResponse(errno.ErrInvalidFace)
	}
	if len(in.Phone) == 0 && len(in.Email) == 0 && len(in.Password) < 6 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	if ctx.Value(fs_metadata_transport.MetadataTransportKey) == nil {
		return errno.ErrResponse(errno.ErrTransfer)
	}
	//project := ctx.Value("project").(*fs_base_project.ProjectInfo)
	//strategy := ctx.Value("strategy").(*fs_base.ProjectStrategy)
	//to := ctx.Value("validate_to").(string)
	meta := ctx.Value("meta").(*fs_base.Metadata)
	resp, err := svc.usercli.Add(context.Background(), &fs_base_user.AddRequest{
		Level:         5,
		Password:      in.Password,
		Phone:         in.Phone,
		Email:         in.Email,
		FromProjectId: meta.ProjectId,
		FromClientId:  meta.ClientId,
	})

	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !resp.State.Ok {
		return errno.ErrResponse(resp.State)
	}

	r, err := svc.facecli.Upsert(context.Background(), &fs_base_face.UpsertRequest{
		Base64Face: in.Meta.Face,
		UserId:     resp.Content,
	})

	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !r.State.Ok {
		return errno.ErrResponse(r.State)
	}

	svc.reportercli.Write(fs_function_tags.GetAdminFuncTag(), resp.Content, meta.Ip)

	return errno.ErrResponse(errno.Ok)
}

func (svc *registerService) FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error) {
	if len(in.Phone) == 0 && len(in.Password) < 6 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	panic(errno.ERROR)
}

//从第三方注册不需要验证码
func (svc *registerService) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {
	panic(errno.ERROR)
}

func NewService(usercli user.Service, repotercli reportercli.Channel, facecli face.Service) Service {
	var svc Service
	{
		svc = &registerService{usercli: usercli, reportercli: repotercli, facecli: facecli}
	}
	return svc
}
