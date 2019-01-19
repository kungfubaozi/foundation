package function

import (
	"context"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/model"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error)

	Init(ctx context.Context, in *fs_base_function.InitRequest) (*fs_base.Response, error)
}

type functionService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
}

func upsert(c *mgo.Collection, session string, f *fs_pkg_model.APIFunction) {
	c.Upsert(bson.M{"tag": f.Function.Tag}, &Model{
		Func:     f.Function.Func,
		ZH:       f.Function.Zh,
		Level:    f.Function.Level,
		Fcv:      f.Function.Fcv,
		EN:       f.Function.En,
		CreateAt: time.Now().UnixNano(),
		Tag:      utils.Md5(f.Prefix + f.Infix + session),
	})
}

func (svc *functionService) Init(ctx context.Context, in *fs_base_function.InitRequest) (*fs_base.Response, error) {
	c := svc.session.Clone().DB(fs_constants.DB_BASE).C("functions")
	i, _ := c.Count()
	if i > 0 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	//login functions
	upsert(c, in.Session, fs_functions.GetEntryByFaceFunc())
	upsert(c, in.Session, fs_functions.GetEntryByValidateCodeFunc())
	upsert(c, in.Session, fs_functions.GetEntryByAPFunc())
	upsert(c, in.Session, fs_functions.GetEntryByOAuthFunc())
	upsert(c, in.Session, fs_functions.GetEntryByQRCodeFunc())
	upsert(c, in.Session, fs_functions.GetEntryByInviteFunc())

	//safety verification functions
	upsert(c, in.Session, fs_functions.GetVerificationUpdatePasswordFunc())
	upsert(c, in.Session, fs_functions.GetVerificationLoginFunc())
	upsert(c, in.Session, fs_functions.GetVerificationRegisterFunc())
	upsert(c, in.Session, fs_functions.GetVerificationResetPasswordFunc())

	//register functions
	upsert(c, in.Session, fs_functions.GetFromAPFunc())
	upsert(c, in.Session, fs_functions.GetFromOAuthFunc())

	//safety update functions
	upsert(c, in.Session, fs_functions.GetUpdateEmailFunc())
	upsert(c, in.Session, fs_functions.GetResetPasswordFunc())
	upsert(c, in.Session, fs_functions.GetUpdatePasswordFunc())
	upsert(c, in.Session, fs_functions.GetUpdatePhoneFunc())

	//unblock
	upsert(c, in.Session, fs_functions.GetUnlockFunc())

	//blacklist
	upsert(c, in.Session, fs_functions.GetAddBlacklistFunc())
	upsert(c, in.Session, fs_functions.GetRemoveBlacklistFunc())

	//function
	upsert(c, in.Session, fs_functions.GetAddFunc())
	upsert(c, in.Session, fs_functions.GetAllFunc())
	upsert(c, in.Session, fs_functions.GetFindFunc())
	upsert(c, in.Session, fs_functions.GetRemoveFunc())
	upsert(c, in.Session, fs_functions.GetUpdateFunc())

	//project functions
	upsert(c, in.Session, fs_functions.GetCreateProject())
	upsert(c, in.Session, fs_functions.GetRemoveProject())
	upsert(c, in.Session, fs_functions.GetUpdateProject())
	upsert(c, in.Session, fs_functions.GetEnablePlatform())

	//usersync functions
	upsert(c, in.Session, fs_functions.GetAddUserSyncHookFunc())
	upsert(c, in.Session, fs_functions.GetRemoveUserSyncHookFunc())
	upsert(c, in.Session, fs_functions.GetUpdateUserSyncHookFunc())

	//strategy functions
	upsert(c, in.Session, fs_functions.GetUpdateProjectStrategyFunc())

	//review functions

	//invite functions
	upsert(c, in.Session, fs_functions.GetInviteUserFunc())

	//froze
	upsert(c, in.Session, fs_functions.GetRequestFrozeFunc())

	//interceptor
	upsert(c, in.Session, fs_functions.GetInterceptFunc())

	//auth
	upsert(c, in.Session, fs_functions.GetRefreshFunc())
	return &fs_base.Response{State: errno.Ok}, nil
}

func (svc *functionService) Add(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base_function.UpsertResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	meta := fs_metadata_transport.ContextToProject(ctx)

	f, err := repo.Get(in.Api)
	if err == mgo.ErrNotFound {
		err = nil
	}
	if err != nil {
		return &fs_base_function.UpsertResponse{
			State: errno.ErrSystem,
		}, nil
	}
	if f != nil {
		return &fs_base_function.UpsertResponse{
			State: errno.ErrRequest,
		}, nil
	}
	uid, _ := uuid.NewV1()

	f = &Model{
		Func:  uuid.NewV5(uid, in.Api).String()[24:],
		ZH:    in.Zh,
		Tag:   utils.Md5(in.Api + meta.Session),
		Level: in.Level,
		EN:    in.En,
		Fcv:   in.Fcv,
	}
	err = repo.Add(f)
	if err != nil {
		return &fs_base_function.UpsertResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_function.UpsertResponse{
		State: errno.Ok,
		Func:  f.Func,
	}, nil
}

func (svc *functionService) Remove(ctx context.Context, in *fs_base_function.RemoveRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.Remove(in.Func, "func")
	if err == mgo.ErrNotFound {
		return errno.ErrResponse(errno.ErrInvalid)
	}
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *functionService) Update(ctx context.Context, in *fs_base_function.UpsertRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *functionService) Get(ctx context.Context, in *fs_base_function.GetRequest) (*fs_base_function.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	var function *Model
	var err error
	if len(in.Func) != 0 {
		function, err = repo.FindByFunc(in.Func)
	} else {
		function, err = repo.Get(in.Tag)
	}
	if err == mgo.ErrNotFound && len(in.Func) == 0 { //不是通过func查找的返回未找到
		return &fs_base_function.GetResponse{State: errno.Ok, Func: &fs_base_function.Func{
			Fcv:   fs_constants.FCV_AUTH,
			Level: fs_constants.LEVEL_USER,
		}}, nil
	}
	if err != nil {
		return &fs_base_function.GetResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_function.GetResponse{
		State: errno.Ok,
		Func: &fs_base_function.Func{
			Zh:   function.ZH,
			Tag:  function.Tag,
			En:   function.EN,
			Fcv:  function.Fcv,
			Func: function.Func,
		},
	}, nil
}

func NewService(session *mgo.Session, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &functionService{session: session, reportercli: reportercli}
	}
	return svc
}

func (svc *functionService) GetRepo() repository {
	return &functionRepository{session: svc.session.Clone()}
}
