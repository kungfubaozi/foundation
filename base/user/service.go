package user

import (
	"context"
	"github.com/go-kit/kit/log"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	logger   log.Logger
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
		UserId:        user.UserId.Hex(),
		FromProjectId: user.FromProjectId,
		FromClientId:  user.FromClientId,
		Level:         user.Level,
		Phone:         user.Phone,
		Email:         user.Email,
		Enterprise:    user.Enterprise,
	}, nil
}

func (svc *userService) FindByUserId(ctx context.Context, in *fs_base_user.FindRequest) (*fs_base_user.FindResponse, error) {
	return svc.findByKey(ctx, "_id", in.Value)
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

	resp, _ := svc.findByKey(ctx, "phone", in.Value)
	if len(resp.UserId) > 0 {
		return errno.ErrResponse(errno.ErrAlreadyBind)
	}

	err := repo.UpdatePhone(meta.UserId, in.Value)
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

	resp, _ := svc.findByKey(ctx, "email", in.Value)
	if len(resp.UserId) > 0 {
		return errno.ErrResponse(errno.ErrAlreadyBind)
	}

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

	resp, _ := svc.findByKey(ctx, "enterprise", in.Value)
	if len(resp.UserId) > 0 {
		return errno.ErrResponse(errno.ErrAlreadyBind)
	}

	err := repo.UpdateEnterprise(meta.UserId, in.Value)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *userService) UpdatePassword(ctx context.Context, in *fs_base_user.UpdateRequest) (*fs_base.Response, error) {
	if len(in.Value) < 6 {
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

	err = repo.UpdatePassword(meta.UserId, string(p))
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
	u := &User{
		CreateAt:      time.Now().UnixNano(),
		Password:      string(p),
		UserId:        bson.NewObjectId(),
		Email:         in.Email,
		Level:         in.Level,
		Phone:         in.Phone,
		FromClientId:  in.FromClientId,
		FromProjectId: in.FromProjectId,
	}
	repo := svc.GetRepo()
	defer repo.Close()
	//注册管理员(系统管理员)
	if in.Level == 5 {
		i := repo.FindAdmin()
		if i != 0 {
			return errno.ErrResponse(errno.ErrRequest)
		}
	} else {
		//查找相同的
		err = repo.FindSame(in.Phone, in.Email, in.Enterprise)
		if err != nil {
			svc.logger.Log("match", "err", "info", err)
			return errno.ErrResponse(errno.ErrUserAlreadyExists)
		}
	}
	err = repo.Add(u)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	resp, err := svc.statecli.Upsert(ctx, &fs_base_state.UpsertRequest{
		Key:    u.UserId.Hex(),
		Status: names.F_STATE_OK,
	})
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return &fs_base.Response{
		State:   resp.State,
		Content: u.UserId.Hex(),
	}, nil
}

func NewService(session *mgo.Session, statecli state.Service, logger log.Logger) Service {
	var svc Service
	{
		svc = &userService{session: session, statecli: statecli, logger: logger}
	}
	return svc
}
