package user

import (
	"context"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
)

//gRPC
type Service interface {
	Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error)

	FindByUserId(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error)

	FindByPhone(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error)

	FindByEnterprise(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error)

	FindByEmail(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error)

	UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)

	UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error)
}

type userService struct {
	session  *mgo.Session
	statecli state.Service
}

func (svc *userService) findByKey(ctx context.Context, key, value string) (*fs_base_user.FindResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	user, err := repo.Get(value, key)
	if err != nil {
		return &fs_base_user.FindResponse{
			State: errno.ErrInvalid,
		}, nil
	}
	return &fs_base_user.FindResponse{
		State:         errno.Ok,
		UserId:        user.UserId,
		FromProjectId: user.FromProjectId,
		FromAppId:     user.FromAppId,
		Level:         user.Level,
		Phone:         user.Phone,
		Email:         user.Email,
		Enterprise:    user.Enterprise,
	}, nil
}

func (svc *userService) FindByUserId(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	return svc.findByKey(ctx, "user_id", in.Value)
}

func (svc *userService) FindByPhone(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	return svc.findByKey(ctx, "phone", in.Value)
}

func (svc *userService) FindByEnterprise(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	return svc.findByKey(ctx, "enterprise", in.Value)
}

func (svc *userService) FindByEmail(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	return svc.findByKey(ctx, "email", in.Value)
}

func (svc *userService) GetRepo() repository {
	return &userRepository{session: svc.session.Clone()}
}

func (svc *userService) UpdatePhone(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	meta := ctx.Value("meta").(*fs_base.Metadata)

	//update
	repo := svc.GetRepo()
	defer repo.Close()

	p, err := bcrypt.GenerateFromPassword([]byte(in.Value), bcrypt.DefaultCost)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	err = repo.UpdatePhone(meta.UserId, string(p))
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdateEmail(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.UpdateEmail(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdateEnterprise(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.UpdateEnterprise(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) == 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}
	meta := ctx.Value("meta").(*fs_base.Metadata)

	//update
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.UpdatePassword(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) Add(ctx context.Context, in *fs_base_user.AddRequest) (*fs_base.Response, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	u := &user{
		CreateAt:      time.Now().UnixNano(),
		Password:      string(p),
		UserId:        uuid.NewV4().String(),
		Email:         in.Email,
		Level:         in.Level,
		Phone:         in.Phone,
		FromAppId:     in.FromAppId,
		FromProjectId: in.FromProjectId,
	}
	repo := svc.GetRepo()
	defer repo.Close()
	err = repo.Add(u)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	resp, err := svc.statecli.Upsert(ctx, &fs_base_state.UpsertRequest{
		Key:    u.UserId,
		Status: names.F_STATE_OK,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(resp.State)
}

func NewService(session *mgo.Session, statecli state.Service) Service {
	var svc Service
	{
		svc = &userService{session: session, statecli: statecli}
	}
	return svc
}
