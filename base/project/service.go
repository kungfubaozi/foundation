package project

import (
	"context"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/strategy/def"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
)

type Service interface {
	New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error)

	Get(ctx context.Context, in *fs_base_project.GetRequest) (*fs_base_project.GetResponse, error)

	EnablePlatform(ctx context.Context, in *fs_base_project.EnablePlatformRequest) (*fs_base.Response, error)
}

type projectService struct {
	session     *mgo.Session
	strategycli strategy.Service
	reportercli reporter.Service
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
	gp := &fs_base_project.ProjectInfo{
		Logo: p.Logo,
		Desc: p.Desc,
		En:   p.EN,
		Zh:   p.ZH,
	}

	if err != nil {
		return resp, nil
	}

	if len(p.Platforms) != 5 {
		return resp, nil
	}

	for _, v := range p.Platforms {
		if !v.Enabled {
			resp.State = errno.ErrPlatform
			return resp, nil
		}
		gp.Platforms = append(gp.Platforms, &fs_base_project.Platform{
			ClientId: v.ClientId,
			Enabled:  v.Enabled,
			Platform: v.Platform,
		})
	}

	resp.Info = gp

	r, err := svc.strategycli.Get(ctx, &fs_base_strategy.GetRequest{
		ProjectId: p.Id,
	})
	if err != nil {
		return resp, nil
	}
	if !r.State.Ok {
		resp.State = r.State
		return resp, nil
	}

	return &fs_base_project.GetResponse{
		State:    errno.Ok,
		Info:     gp,
		Strategy: r.Strategy,
	}, nil
}

func (svc *projectService) New(ctx context.Context, in *fs_base_project.NewRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	p := &project{
		Id:       uuid.NewV4().String(),
		Logo:     in.Logo,
		ZH:       in.Zh,
		EN:       in.En,
		CreateAt: time.Now().UnixNano(),
		Desc:     in.Desc,
		Platforms: []*platform{
			{
				ClientId: uuid.NewV4().String(),
				Platform: names.F_PLATFORM_ANDROID,
				Enabled:  true,
			},
			{
				ClientId: uuid.NewV4().String(),
				Platform: names.F_PLATFORM_IOS,
				Enabled:  true,
			},
			{
				ClientId: uuid.NewV4().String(),
				Platform: names.F_PLATFORM_WINDOWD,
				Enabled:  true,
			},
			{
				ClientId: uuid.NewV4().String(),
				Platform: names.F_PLATFORM_MAC_OS,
				Enabled:  true,
			},
			{
				ClientId: uuid.NewV4().String(),
				Platform: names.F_PLATFORM_WEB,
				Enabled:  true,
			},
		},
	}

	errc := make(chan error, 2)

	go func() {
		errc <- repo.Save(p)
	}()

	go func() {
		svc.strategycli.New(ctx, &fs_base.ProjectStrategy{
			ProjectId: p.Id,
			CreateAt:  p.CreateAt,
			Version:   1,
			Configuration: &fs_base.Configuration{
				OpenTime: "0-24",
			},
			Events: &fs_base.Events{
				OnRegister: &fs_base.OnRegister{},
				OnLogin: &fs_base.OnLogin{
					AllowOtherProjectUserToLogin: 2,
					Mode: []int64{
						1, 2, 3, 4, 5, 6,
					},
					MaxCountOfOnline: &fs_base.MaxCountOfOnline{
						Android: 1,
						IOS:     1,
						Windows: 1,
						MacOS:   1,
						Web:     0, //无限制
					},
					Verification: 1,
					MaxCountOfErrorPassword: []*fs_base.MaxCountOfErrorPassword{
						{
							Count:  3,
							Action: 4,
						},
						{
							Count:  5,
							Action: 5,
						},
						{
							Count:       8,
							Action:      3,
							ExpiredTime: 10 * 60, //10分钟
						},
					},
					MaxCountOfInvalidAccount: []*fs_base.MaxCountOfInvalidAccount{
						{
							Count:  3,
							Action: 4,
						},
						{
							Count:  5,
							Action: 5,
						},
						{
							Count:       8,
							Action:      3,
							ExpiredTime: 10 * 60, //10分钟
						},
					},
					Submitlal: 1, //不用提交经纬度
				},
				OnVerification: strategydef.GetOnVerificationDefault(),
				OnQRLogin: &fs_base.OnQRLogin{
					RefreshDuration: 60, //单位秒
				},
				OnFaceLogin: &fs_base.OnFaceLogin{
					Degree: 80,
				},
				OnCommonEquipmentChanges: &fs_base.OnCommonEquipmentChanges{
					SendMessageToUser: 2,
				},
				OnRequestFrozen: &fs_base.OnRequestFrozen{
					Verification: 2,
				},
				OnCancelFrozen: &fs_base.OnCancelFrozen{
					Verification: 2,
				},
				OnChangePhoneNumber: &fs_base.OnChangePhoneNumber{
					Verification: 2,
				},
				OnChangeEmail: &fs_base.OnChangeEmail{
					Verification: 2,
				},
				OnChangeFace: &fs_base.OnChangeFace{
					Verification: 2,
				},
				OnChangeOAuth: &fs_base.OnChangeOAuth{
					Verification: 1,
				},
				OnResetPassword: &fs_base.OnResetPassword{ //两种方式都可重置密码
					Phone: 2,
					Email: 2,
				},
				OnElsewhereLogin: &fs_base.OnElsewhereLogin{
					SendMessageToUser: 2,
					Verification:      2,
				},
				OnSubmitReview: &fs_base.OnSubmitReview{},
				OnInviteUser: &fs_base.OnInviteUser{
					ExpireTime: 10 * 60 * 24 * 3, //3天
					Review:     2,
				},
			},
		})
	}()

	if errc != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func NewService(session *mgo.Session, strategycli strategy.Service, reportercli reporter.Service) Service {
	var svc Service
	{
		svc = &projectService{session: session, strategycli: strategycli, reportercli: reportercli}
	}
	return svc
}
