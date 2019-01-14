package login

import (
	"context"
	"sync"
	"time"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/invite"
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/entry/login/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/tool/number"
	"zskparker.com/foundation/pkg/transport"
)

type Service interface {
	EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error)

	EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_entry_login.EntryResponse, error)

	EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_entry_login.EntryResponse, error)

	EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_entry_login.EntryResponse, error)

	//admin
	EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error)

	//使用邀请码登录
	EntryByInvite(ctx context.Context, in *fs_entry_login.EntryByInviteRequest) (*fs_entry_login.EntryResponse, error)
}

type loginService struct {
	usercli         user.Service
	reportercli     reportercli.Channel
	authenticatecli authenticate.Service
	validatecli     validate.Service
	facecli         face.Service
	redisync        *fs_redisync.Redisync
	invitecli       invite.Service
}

func (svc *loginService) getMaxOnlineCount(platform int64, strategy *fs_base.MaxCountOfOnline) int64 {
	if platform == fs_constants.PLATFORM_ANDROID {
		return strategy.Android
	} else if platform == fs_constants.PLATFORM_WEB {
		return strategy.Web
	} else if platform == fs_constants.PLATFORM_IOS {
		return strategy.IOS
	} else if platform == fs_constants.PLATFORM_WINDOWD {
		return strategy.Windows
	} else if platform == fs_constants.PLATFORM_MAC_OS {
		return strategy.MacOS
	}
	return -1
}

func (svc *loginService) checkProjectLevel(level int64, ctx context.Context) bool {
	p := fs_metadata_transport.ContextToProject(ctx)
	return level >= p.Level
}

func (svc *loginService) getAuthorize(meta *fs_base.Metadata, userId, mode string, level, platform int64) *fs_base_authenticate.Authorize {
	return &fs_base_authenticate.Authorize{
		UserId:    userId,
		LoginMode: mode,
		Timestamp: time.Now().UnixNano(),
		ProjectId: meta.ProjectId,
		Platform:  platform,
		ClientId:  meta.ClientId,
		Ip:        meta.Ip,
		Level:     level,
		Device:    meta.Device,
		UserAgent: meta.UserAgent,
	}
}

//如果有邀请
//输入手机号和邀请码
func (svc *loginService) EntryByInvite(ctx context.Context, in *fs_entry_login.EntryByInviteRequest) (*fs_entry_login.EntryResponse, error) {
	resp := func(s *fs_base.State) (*fs_entry_login.EntryResponse, error) {
		return &fs_entry_login.EntryResponse{State: s}, nil
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)

	//邀请码为8位数字码
	if len(in.Code) != 8 {
		return resp(errno.ErrRequest)
	}

	ir, err := svc.invitecli.Get(context.Background(), &fs_base_invite.GetRequest{
		Code: in.Code,
	})
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if !ir.State.Ok {
		return resp(ir.State)
	}

	//找到用户就把用户移动到用户表里并删除在邀请里的用户

	wg := sync.WaitGroup{}
	wg.Add(2)

	ps := errno.Ok

	errc := func(s *fs_base.State) {
		if ps.Ok {
			ps = s
		}
		wg.Done()
	}

	//写入用户表
	go func() {
		resp, err := svc.usercli.Add(context.Background(), &fs_base_user.AddRequest{
			Password:      fs_tools_number.GetRndNumber(8), //设置随机密码
			Enterprise:    ir.Detail.Enterprise,
			Username:      ir.Detail.Username,
			Phone:         ir.Detail.Phone,
			Email:         ir.Detail.Email,
			FromProjectId: meta.ProjectId,
			FromClientId:  meta.ClientId,
			Level:         ir.Detail.Level,
			RealName:      ir.Detail.RealName,
			Reset_:        true, //需要重置密码
		})
		if err != nil {
			errc(errno.ErrSystem)
			return
		}
		errc(resp.State)
	}()

	//更新
	go func() {
		resp, err := svc.invitecli.Update(context.Background(), &fs_base_invite.UpdateRequest{
			InviteCode: in.Code,
			InviteId:   ir.InviteId,
			Account:    ir.Detail.Phone + ir.Detail.Email,
		})
		if err != nil {
			errc(errno.ErrSystem)
			return
		}
		errc(resp.State)
	}()

	//重置密码
	return &fs_entry_login.EntryResponse{State: errno.ErrResetPassword}, nil
}

func (svc *loginService) EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error) {
	resp := func(state *fs_base.State) (*fs_entry_login.EntryResponse, error) {
		return &fs_entry_login.EntryResponse{State: state}, nil
	}
	meta := fs_metadata_transport.ContextToMeta(ctx)
	strategy := fs_metadata_transport.ContextToStrategy(ctx)

	//不允许登录
	if strategy.Events.OnLogin.AllowLogin == 1 {
		return resp(errno.ErrRequest)
	}

	project := fs_metadata_transport.ContextToProject(ctx)
	//查找人脸库
	fr, _ := svc.facecli.Search(context.Background(), &fs_base_face.SearchRequest{
		Base64Face: in.Meta.Face,
	})
	if !fr.State.Ok {
		return resp(fr.State)
	}

	if !svc.checkProjectLevel(fr.Level, ctx) {
		return resp(errno.ErrProjectPermission)
	}

	if s := svc.redisync.Lock(fs_function_tags.GetEntryByFaceFuncTag(), meta.UserId, 3); s != nil {
		return resp(s)
	}

	ar, _ := svc.authenticatecli.New(context.Background(), &fs_base_authenticate.NewRequest{
		Authorize:      svc.getAuthorize(meta, fr.UserId, "face", fr.Level, project.Platform.Platform),
		MaxOnlineCount: svc.getMaxOnlineCount(meta.Platform, strategy.Events.OnLogin.MaxCountOfOnline),
	})
	if !ar.State.Ok {
		return resp(ar.State)
	}

	svc.reportercli.Write(fs_function_tags.GetEntryByFaceFuncTag(), meta.UserId, meta.ProjectId)

	return &fs_entry_login.EntryResponse{
		State:        errno.Ok,
		Session:      ar.Session,
		RefreshToken: ar.RefreshToken,
		AccessToken:  ar.AccessToken,
	}, nil
}

func (svc *loginService) EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error) {
	if len(in.Account) < 6 || len(in.Password) < 6 {
		return &fs_entry_login.EntryResponse{State: errno.ErrRequest}, nil
	}

	resp := func(state *fs_base.State) (*fs_entry_login.EntryResponse, error) {
		return &fs_entry_login.EntryResponse{State: state}, nil
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)
	strategy := fs_metadata_transport.ContextToStrategy(ctx)

	//不允许登录
	if strategy.Events.OnLogin.AllowLogin == 1 {
		return resp(errno.ErrRequest)
	}

	project := fs_metadata_transport.ContextToProject(ctx)
	var u *fs_base_user.FindResponse
	var err error
	req := &fs_base_user.FindRequest{
		Value:    in.Account,
		Password: in.Password,
	}
	if fs_regx_match.Phone(in.Account) {
		u, err = svc.usercli.FindByPhone(context.Background(), req)
	} else if fs_regx_match.Email(in.Account) {
		u, err = svc.usercli.FindByEmail(context.Background(), req)
	} else { //enterprise
		u, err = svc.usercli.FindByEnterprise(context.Background(), req)
	}
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if !u.State.Ok {
		return resp(u.State)
	}

	if !svc.checkProjectLevel(u.Level, ctx) {
		return resp(errno.ErrProjectPermission)
	}

	//锁3s
	if s := svc.redisync.Lock(fs_function_tags.GetEntryByAPFuncTag(), u.UserId, 3); s != nil {
		return resp(s)
	}

	a, err := svc.authenticatecli.New(context.Background(), &fs_base_authenticate.NewRequest{
		MaxOnlineCount: svc.getMaxOnlineCount(meta.Platform, strategy.Events.OnLogin.MaxCountOfOnline),
		Authorize:      svc.getAuthorize(meta, u.UserId, "ap", u.Level, project.Platform.Platform),
	})
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if !a.State.Ok {
		return resp(a.State)
	}

	svc.reportercli.Write(fs_function_tags.GetEntryByAPFuncTag(), meta.UserId, meta.ProjectId)

	return &fs_entry_login.EntryResponse{
		State:        errno.Ok,
		RefreshToken: a.RefreshToken,
		Session:      a.Session,
		AccessToken:  a.AccessToken,
	}, nil
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
	validatecli validate.Service, facecli face.Service, redisync *fs_redisync.Redisync) Service {
	var service Service
	{
		service = &loginService{usercli: usercli, reportercli: reportercli,
			authenticatecli: authenticatecli, validatecli: validatecli, facecli: facecli,
			redisync: redisync}
	}
	return service
}
