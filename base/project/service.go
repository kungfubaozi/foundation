package project

import (
	"context"
	"encoding/base64"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/strategy/def"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error)

	EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error)

	Init(ctx context.Context, in *fs_base_project.InitRequest) (*fs_base_project.InitResponse, error)
}

type projectService struct {
	session     *mgo.Session
	strategycli strategy.Service
	reportercli reportercli.Channel
}

func (svc *projectService) Init(ctx context.Context, in *fs_base_project.InitRequest) (*fs_base_project.InitResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	if repo.Size() > 0 {
		return &fs_base_project.InitResponse{State: errno.ErrRequest}, nil
	}

	p := defProject(in.Logo, in.Zh, in.En, in.Desc, in.UserId, bson.NewObjectId(), 3)
	r, err := svc.Create(ctx, p)
	if err != nil {
		return &fs_base_project.InitResponse{State: errno.ErrSystem}, nil
	}
	if !r.State.Ok {
		return &fs_base_project.InitResponse{State: r.State}, nil
	}
	return &fs_base_project.InitResponse{
		State:     errno.Ok,
		Session:   p.Session,
		ProjectId: p.Id.Hex(),
	}, nil
}

func (svc *projectService) GetRepo() repository {
	return &projectRepository{session: svc.session.Clone()}
}

func (svc *projectService) EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	meta := fs_metadata_transport.ContextToMeta(ctx)

	err := repo.Enable(meta.ProjectId, in.Platform, in.Enable)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *projectService) Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	resp := func(state *fs_base.State) (*fs_base_project.GetResponse, error) {
		return &fs_base_project.GetResponse{State: state}, nil
	}

	p, err := repo.Get(in.ClientId)

	if err != nil {
		return resp(errno.ErrRequest)
	}

	gp := &fs_base_project.ProjectInfo{
		Logo:    p.Logo,
		Desc:    p.Desc,
		En:      p.EN,
		Zh:      p.ZH,
		Level:   p.Level,
		Session: p.Session,
	}

	if len(p.Platforms) != 5 {
		return resp(errno.ErrSystem)
	}

	for _, v := range p.Platforms {
		if v.ClientId == in.ClientId {
			if !v.Enabled { //未开启平台
				return resp(errno.ErrPlatform)
			}
			gp.Platform = &fs_base_project.Platform{
				ClientId: v.ClientId,
				Platform: v.Platform,
			}
			break
		}
	}

	if gp.Platform == nil {
		return &fs_base_project.GetResponse{}, nil
	}

	r, err := svc.strategycli.Get(context.Background(), &fs_base_strategy.GetRequest{
		ProjectId: p.Id.Hex(),
	})
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if !r.State.Ok {
		return resp(r.State)
	}

	return &fs_base_project.GetResponse{
		State:    errno.Ok,
		Info:     gp,
		Strategy: r.Strategy,
	}, nil
}

func (svc *projectService) New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error) {
	meta := fs_metadata_transport.ContextToMeta(ctx)
	p := defProject(in.Logo, in.Zh, in.En, in.Desc, meta.UserId, bson.NewObjectId(), 3)
	return svc.Create(ctx, p)
}

func (svc *projectService) Create(ctx context.Context, p *project) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.Exists(p.EN)
	if err != nil && err == errno.ERROR {
		err = nil
		return errno.ErrResponse(errno.ErrProjectAlreadyExists)
	} else if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	errc := make(chan error, 2)

	go func() {
		errc <- repo.Save(p)
	}()

	go func() {
		svc.strategycli.Upsert(ctx, &fs_base_strategy.UpsertRequest{
			Strategy: strategydef.DefStrategy(p.Id.Hex(), p.Creator),
		})
	}()

	if err := <-errc; err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func NewService(session *mgo.Session, strategycli strategy.Service, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &projectService{session: session, strategycli: strategycli, reportercli: reportercli}
	}
	return svc
}

func defProject(logo, zh, en, desc, creator string, id bson.ObjectId, level int64) *project {
	node := utils.NodeGenerate()
	return &project{
		Id:       id,
		Logo:     logo,
		ZH:       zh,
		EN:       en,
		Creator:  creator,
		CreateAt: time.Now().UnixNano(),
		Desc:     desc,
		Level:    level,
		Session:  base64.StdEncoding.EncodeToString([]byte(uuid.NewV4().String())),
		Platforms: []*platform{
			{
				ClientId: node.Generate().Base64(),
				Platform: names.F_PLATFORM_ANDROID,
				Enabled:  true,
			},
			{
				ClientId: node.Generate().Base64(),
				Platform: names.F_PLATFORM_IOS,
				Enabled:  true,
			},
			{
				ClientId: node.Generate().Base64(),
				Platform: names.F_PLATFORM_WINDOWD,
				Enabled:  true,
			},
			{
				ClientId: node.Generate().Base64(),
				Platform: names.F_PLATFORM_MAC_OS,
				Enabled:  true,
			},
			{
				ClientId: node.Generate().Base64(),
				Platform: names.F_PLATFORM_WEB,
				Enabled:  true,
			},
		},
	}
}
