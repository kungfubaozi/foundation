package invite

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/tool/encrypt"
	"zskparker.com/foundation/pkg/tool/number"
	"zskparker.com/foundation/pkg/transport"
)

type Service interface {
	Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_invite.GetRequest) (*fs_base_invite.GetResponse, error)

	Update(ctx context.Context, in *fs_base_invite.UpdateRequest) (*fs_base.Response, error)

	GetInvites(ctx context.Context, in *fs_base_invite.GetInvitesRequest) (*fs_base_invite.GetInvitesResponse, error)
}

type inviteService struct {
	session     *mgo.Session
	messagecli  messagecli.Channel
	reportercli reportercli.Channel
	usercli     user.Service
}

func NewService(session *mgo.Session, messagecli messagecli.Channel, reportercli reportercli.Channel, usercli user.Service) Service {
	var svc Service
	{
		svc = &inviteService{session: session, messagecli: messagecli, reportercli: reportercli, usercli: usercli}
	}
	return svc
}

func (svc *inviteService) GetRepo() repository {
	return &inviteRepository{session: svc.session.Clone()}
}

func (svc *inviteService) Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	fmt.Println("enter")

	if len(in.Account) < 10 {
		return errno.ErrResponse(errno.ErrInviteAccount)
	}

	if len(in.Username) < 2 {
		return errno.ErrResponse(errno.ErrInviteUsername)
	}

	//限定，只能邀请用户，开发者，应用管理员
	if in.Level != fs_constants.LEVEL_USER && in.Level != fs_constants.LEVEL_DEVELOPER && in.Level != fs_constants.LEVEL_PROJECT_MANAGER {
		return errno.ErrResponse(errno.ErrInviteLevel)
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)
	strategy := fs_metadata_transport.ContextToStrategy(ctx)

	m := &model{}
	var err error
	var u *fs_base_user.FindResponse

	if fs_regx_match.Phone(in.Account) {
		m, err = repo.FindInviteByAccount(in.Account, true)
		u, err = svc.usercli.FindByPhone(context.Background(), &fs_base_user.FindRequest{
			Value: in.Account,
		})
		m.Phone = in.Account
	} else if fs_regx_match.Email(in.Account) {
		m, err = repo.FindInviteByAccount(in.Account, false)
		u, err = svc.usercli.FindByEmail(context.Background(), &fs_base_user.FindRequest{
			Value: in.Account,
		})
		m.Email = in.Account
	} else {
		return errno.ErrResponse(errno.ErrInviteAccount)
	}

	if err != nil && err == mgo.ErrNotFound {
		err = nil
		m.Ok = false
	}

	if err != nil || u == nil {
		fmt.Println("err", err)
		fmt.Println("u", u)
		return errno.ErrResponse(errno.ErrSystem)
	}

	if u != nil && u.State.Code != errno.ErrInvalidUser.Code {
		return errno.ErrResponse(errno.ErrUserAlreadyExists)
	}

	//是否已经邀请过且通过了
	if m.Ok {
		return errno.ErrResponse(errno.ErrAlreadyInvited)
	}

	//查找用户是否存在

	//如果存在且没有过期则返回错误
	if len(m.InviteId) > 0 && m.ExpireAt-time.Now().UnixNano() > 0 {
		return errno.ErrResponse(errno.ErrInviteExists)
	}

	if in.Scope != fs_constants.SCOPE_TYPE_OUTTER && in.Scope != fs_constants.SCOPE_TYPE_INNER {
		return errno.ErrResponse(errno.ErrRequest)
	}

	m.Scope = in.Scope
	m.Level = in.Level
	m.Username = in.Username
	m.Enterprise = in.Enterprise
	m.CreateAt = time.Now().UnixNano()
	m.ExpireAt = time.Now().UnixNano() + strategy.Events.OnInviteUser.ExpireTime*60*1e9
	m.Ok = false
	m.OperateUserId = meta.UserId //操作人
	code := fs_tools_number.GetRndNumber(6)
	m.Code = fs_tools_encrypt.SHA1_256_512(code)

	//存在重新邀请的情况
	if len(m.InviteId.Hex()) == 0 {
		m.InviteId = bson.NewObjectId()
	}

	err = repo.Add(m)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	//send message
	if len(m.Phone) > 0 {
		svc.messagecli.SendSMS(&fs_base.DirectMessage{
			To:      m.Phone,
			Content: fmt.Sprintf(osenv.GetInviteMessage(), code),
		})
	} else {
		svc.messagecli.SendEmail(&fs_base.DirectMessage{
			To:      m.Email,
			Content: fmt.Sprintf(osenv.GetInviteMessage(), code),
		})
	}

	fs_metadata_transport.MetaToReporter(svc.reportercli, ctx, in.Account, fs_constants.STATUS_OK)

	return errno.ErrResponse(errno.Ok)
}

//移动完成后需要更新对应的邀请数据
func (svc *inviteService) Update(ctx context.Context, in *fs_base_invite.UpdateRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.Update(in.Account, in.InviteId)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

//获取邀请列表
func (svc *inviteService) GetInvites(ctx context.Context, in *fs_base_invite.GetInvitesRequest) (*fs_base_invite.GetInvitesResponse, error) {
	panic("implement me")
}

func (svc *inviteService) Get(ctx context.Context, in *fs_base_invite.GetRequest) (*fs_base_invite.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	resp := func(s *fs_base.State) (*fs_base_invite.GetResponse, error) {
		return &fs_base_invite.GetResponse{State: s}, nil
	}

	m, err := repo.FindInvite(fs_tools_encrypt.SHA1_256_512(in.Code))

	if err != nil {
		return resp(errno.ErrSystem)
	}

	if m != nil && len(m.InviteId) > 10 {

		//邀请过期
		if m.ExpireAt-time.Now().UnixNano() < 0 {
			return resp(errno.ErrExpired)
		}

		if m.Ok {
			return resp(errno.ErrInviteCodeExpired)
		}

		return &fs_base_invite.GetResponse{
			State:    errno.Ok,
			InviteId: m.InviteId.Hex(),
			Detail: &fs_base_invite.InviteInfo{
				Phone:         m.Phone,
				Email:         m.Email,
				OkAt:          m.OkTime,
				OperateUserId: m.OperateUserId,
				Enterprise:    m.Enterprise,
				Username:      m.Username,
				RealName:      m.RealName,
				Level:         m.Level,
			},
		}, nil
	}

	return resp(errno.ErrInvalid)
}
