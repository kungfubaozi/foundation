package invite

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/invite/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/match"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/tool/encrypt"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_invite.GetRequest) (*fs_base_invite.GetResponse, error)

	MoveToUser(ctx context.Context, in *fs_base_invite.MoveToUserRequest) (*fs_base.Response, error)
}

type inviteService struct {
	session     *mgo.Session
	usercli     user.Service
	messagecli  messagecli.Channel
	reportercli reportercli.Channel
}

func (svc *inviteService) GetRepo() repository {
	return &inviteRepository{session: svc.session.Clone()}
}

func (svc *inviteService) Add(ctx context.Context, in *fs_base_invite.AddRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

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

	if fs_regx_match.Phone(in.Account) {
		m, err = repo.Get("phone", in.Account)
		m.Phone = in.Account
	} else if fs_regx_match.Email(in.Account) {
		m, err = repo.Get("email", in.Account)
		m.Email = in.Account
	} else {
		return errno.ErrResponse(errno.ErrInviteAccount)
	}

	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	if len(m.Username) > 0 {
		return errno.ErrResponse(errno.ErrInviteExists)
	}

	m.Level = in.Level
	m.Username = in.Username
	m.Enterprise = in.Enterprise
	m.CreateAt = time.Now().UnixNano()
	m.ExpireAt = strategy.Events.OnInviteUser.ExpireTime * 60 * 1e9

	code := utils.GetRandNumber()
	m.Code = fs_tools_encrypt.SHA1_256_512(code)

	err = repo.Add(m)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	//send message
	if len(m.Phone) > 0 {
		svc.messagecli.SendSMS(&fs_base.DirectMessage{
			To:      m.Phone,
			Content: fmt.Sprintf(osenv.GetInviteMessage(), m.Code),
		})
	} else {
		svc.messagecli.SendEmail(&fs_base.DirectMessage{
			To:      m.Phone,
			Content: fmt.Sprintf(osenv.GetInviteMessage(), m.Code),
		})
	}

	svc.reportercli.Write(fs_function_tags.GetInviteTag(), meta.UserId, meta.ClientId)

	return errno.ErrResponse(errno.Ok)
}

//移动完成后需要删除对应的邀请数据
func (svc *inviteService) MoveToUser(ctx context.Context, in *fs_base_invite.MoveToUserRequest) (*fs_base.Response, error) {
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
		return &fs_base_invite.GetResponse{
			State:    errno.Ok,
			InviteId: m.InviteId.Hex(),
		}, nil
	}

	return resp(errno.ErrInvalid)
}

func NewSerivce(session *mgo.Session, messagecli messagecli.Channel, reportercli reportercli.Channel, usercli user.Service) Service {
	var svc Service
	{
		svc = &inviteService{session: session, reportercli: reportercli, messagecli: messagecli, usercli: usercli}
	}
	return svc
}
