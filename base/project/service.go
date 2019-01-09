package project

import (
	"context"
	"fmt"
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
}

type projectService struct {
	session     *mgo.Session
	strategycli strategy.Service
	reportercli reportercli.Channel
}

func (svc *projectService) GetRepo() repository {
	return &projectRepository{session: svc.session.Clone()}
}

func (svc *projectService) EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	meta := ctx.Value("meta").(*fs_base.Metadata)

	err := repo.Enable(meta.ProjectId, in.Platform, in.Enable)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *projectService) Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	resp := &fs_base_project.GetResponse{
		State: errno.ErrRequest,
	}

	p, err := repo.Get(in.ClientId)

	if err != nil {
		fmt.Println("e0", err)
		return resp, nil
	}

	gp := &fs_base_project.ProjectInfo{
		Logo: p.Logo,
		Desc: p.Desc,
		En:   p.EN,
		Zh:   p.ZH,
	}

	if len(p.Platforms) != 5 {
		fmt.Println("e1")
		return resp, nil
	}

	for _, v := range p.Platforms {
		if !v.Enabled { //未开启平台
			resp.State = errno.ErrPlatform
			return resp, nil
		}
		gp.Platforms = append(gp.Platforms, &fs_base_project.Platform{
			ClientId: v.ClientId,
			Enabled:  v.Enabled,
			Platform: v.Platform,
		})
	}

	r, err := svc.strategycli.Get(context.Background(), &fs_base_strategy.GetRequest{
		ProjectId: p.Id.Hex(),
	})
	if err != nil {
		fmt.Println("err", err)
		return resp, nil
	}
	if !r.State.Ok {
		resp.State = r.State
		return resp, nil
	}

	resp.Info = gp

	return &fs_base_project.GetResponse{
		State:    errno.Ok,
		Info:     gp,
		Strategy: r.Strategy,
	}, nil
}

func (svc *projectService) New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	meta := ctx.Value(fs_metadata_transport.MetadataTransportKey).(*fs_base.Metadata)
	node := utils.NodeGenerate()
	p := defProject(in, meta.UserId, bson.NewObjectId())
	p.Session = node.Generate().Base64()
	p.Platforms = []*platform{
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
	}

	errc := make(chan error, 2)

	go func() {
		errc <- repo.Save(p)
	}()

	go func() {
		svc.strategycli.Upsert(ctx, &fs_base_strategy.UpsertRequest{
			Strategy: strategydef.DefStrategy(p.Id.Hex(), meta.UserId),
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

func InsertDef(session *mgo.Session) {
	p := defProject(&fs_base_project.NewRequest{
		Zh:   "Foundation",
		En:   "Foundation",
		Desc: "Root project",
	}, "admin", bson.ObjectIdHex("5c345ba1133cf43acf167bd9"))
	p.Platforms = []*platform{
		{
			ClientId: "MTA4MjY0NzQ2ODA3MzU1MzkyMA==",
			Platform: names.F_PLATFORM_ANDROID,
			Enabled:  true,
		},
		{
			ClientId: "MTA4MjY0NzQ2ODA3MzU1MzkyMQ==",
			Platform: names.F_PLATFORM_IOS,
			Enabled:  true,
		},
		{
			ClientId: "MTA4MjY0NzQ2ODA3MzU1MzkyMg==",
			Platform: names.F_PLATFORM_WINDOWD,
			Enabled:  true,
		},
		{
			ClientId: "MTA4MjY0NzQ2ODA3MzU1MzkyMw==",
			Platform: names.F_PLATFORM_MAC_OS,
			Enabled:  true,
		},
		{
			ClientId: "MTA4MjY0NzQ2ODA3MzU1MzkyNA==",
			Platform: names.F_PLATFORM_WEB,
			Enabled:  true,
		},
	}
	session.DB("foundation").C("project").Upsert(bson.M{"_id": p.Id}, p)
}

func defProject(in *fs_base_project.NewRequest, creator string, id bson.ObjectId) *project {
	return &project{
		Id:       id,
		Logo:     in.Logo,
		ZH:       in.Zh,
		EN:       in.En,
		Creator:  creator,
		CreateAt: time.Now().UnixNano(),
		Desc:     in.Desc,
	}
}
