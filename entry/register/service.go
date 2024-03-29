package register

import (
	"context"
	"gopkg.in/mgo.v2"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Service interface {
	FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error)

	FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error)
}

type registerService struct {
	usercli      user.Service
	facecli      face.Service
	reportercli  reportercli.Channel
	session      *mgo.Session
	redisync     *fs_redisync.Redisync
	blacklistcli blacklist.Service
}

func (svc *registerService) GetRepo() repository {
	return &registerRepository{session: svc.session.Clone()}
}

func (svc *registerService) FromAP(ctx context.Context, in *fs_entry_register.FromAPRequest) (*fs_base.Response, error) {
	if len(in.Password) < 6 || len(in.Meta.Face) > 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	strategy := fs_metadata_transport.ContextToStrategy(ctx)

	if strategy.Events.OnRegister.AllowNewRegistrations == 1 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)

	mode := strategy.Events.OnRegister.Mode
	var v string
	if mode == 1 { //phone
		if len(in.Email) > 0 {
			return errno.ErrResponse(errno.ErrSupport)
		}
		if !fs_regx_match.Phone(in.Phone) {
			return errno.ErrResponse(errno.ErrPhoneNumber)
		}
		v = in.Phone
	} else if mode == 2 { //email
		if len(in.Phone) > 0 {
			return errno.ErrResponse(errno.ErrSupport)
		}
		if !fs_regx_match.Email(in.Email) {
			return errno.ErrResponse(errno.ErrEmail)
		}
		v = in.Email
	} else { //不支持的操作
		return errno.ErrResponse(errno.ErrSupport)
	}

	if s := fs_metadata_transport.CheckValidateAccount(ctx, v); !s.Ok {
		return errno.ErrResponse(s)
	}

	b, e := svc.blacklistcli.CheckAccount(context.Background(), &fs_safety_blacklist.CheckAccountRequest{
		Meta:    meta,
		Account: v,
	})

	if e != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	if !b.State.Ok {
		return errno.ErrResponse(b.State)
	}

	//锁住当前操作的注册值
	if s := svc.redisync.Lock(fs_function_tags.GetFromAPFuncTag(), v, 3); s != nil {
		return errno.ErrResponse(s)
	}

	resp, err := svc.usercli.Add(context.Background(), &fs_base_user.AddRequest{
		Level:         fs_constants.LEVEL_USER,
		Password:      in.Password,
		Phone:         in.Phone,
		Email:         in.Email,
		FromClientId:  meta.ClientId,
		FromProjectId: meta.ProjectId,
		Scope:         fs_constants.SCOPE_TYPE_OUTTER,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if !resp.State.Ok {
		return errno.ErrResponse(resp.State)
	}

	fs_metadata_transport.MetaToReporter(svc.reportercli, ctx, resp.Content, fs_constants.STATE_OK)

	return errno.ErrResponse(errno.Ok)
}

//从第三方注册不需要验证码
func (svc *registerService) FromOAuth(ctx context.Context, in *fs_entry_register.FromOAuthRequest) (*fs_base.Response, error) {
	panic(errno.ERROR)
}

func NewService(usercli user.Service, repotercli reportercli.Channel, facecli face.Service,
	redisync *fs_redisync.Redisync, blacklistcli blacklist.Service) Service {
	var svc Service
	{
		svc = &registerService{usercli: usercli, reportercli: repotercli, facecli: facecli,
			redisync: redisync, blacklistcli: blacklistcli}
	}
	return svc
}
