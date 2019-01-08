package strategy

import (
	"context"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	Get(ctx context.Context, in *fs_base_strategy.GetRequest) (*fs_base_strategy.GetResponse, error)

	Upsert(ctx context.Context, in *fs_base_strategy.UpsertRequest) (*fs_base.Response, error)
}

type strategyService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
}

func (svc *strategyService) Get(ctx context.Context, in *fs_base_strategy.GetRequest) (*fs_base_strategy.GetResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	p, err := repo.Get(in.ProjectId)
	if err != nil {
		return &fs_base_strategy.GetResponse{State: errno.ErrSystem}, nil
	}
	return &fs_base_strategy.GetResponse{
		State:    errno.Ok,
		Strategy: p,
	}, nil
}

func (svc *strategyService) Upsert(ctx context.Context, in *fs_base_strategy.UpsertRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	p, err := repo.Get(in.Strategy.ProjectId)
	if err != nil && err == mgo.ErrNotFound {
		err = nil
		p = &fs_base.ProjectStrategy{
			Configuration: &fs_base.Configuration{},
			Events: &fs_base.Events{
				OnCommonEquipmentChanges: &fs_base.OnCommonEquipmentChanges{},
				OnCancelFrozen:           &fs_base.OnCancelFrozen{},
				OnChangeFace:             &fs_base.OnChangeFace{},
				OnChangeOAuth:            &fs_base.OnChangeOAuth{},
				OnResetPassword:          &fs_base.OnResetPassword{},
				OnElsewhereLogin:         &fs_base.OnElsewhereLogin{},
				OnSubmitReview:           &fs_base.OnSubmitReview{},
				OnInviteUser:             &fs_base.OnInviteUser{},
				OnLogin: &fs_base.OnLogin{
					MaxCountOfOnline: &fs_base.MaxCountOfOnline{},
				},
				OnRegister:          &fs_base.OnRegister{},
				OnQRLogin:           &fs_base.OnQRLogin{},
				OnFaceLogin:         &fs_base.OnFaceLogin{},
				OnVerification:      &fs_base.OnVerification{},
				OnRequestFrozen:     &fs_base.OnRequestFrozen{},
				OnChangePhoneNumber: &fs_base.OnChangePhoneNumber{},
				OnChangeEmail:       &fs_base.OnChangeEmail{},
			},
		}
	}
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	p.ProjectId = in.Strategy.ProjectId
	p.Version = p.Version + 1

	p.CreateAt = time.Now().UnixNano()

	p.Creator = in.Strategy.Creator

	if in.Strategy.Configuration != nil {
		if len(in.Strategy.Configuration.RegisterReviewId) > 0 {
			p.Configuration.RegisterReviewId = in.Strategy.Configuration.RegisterReviewId
		}
		if len(in.Strategy.Configuration.OpenTime) > 0 {
			p.Configuration.OpenTime = in.Strategy.Configuration.OpenTime
		}
	}

	if in.Strategy.Events != nil {
		if in.Strategy.Events.OnVerification != nil {
			if in.Strategy.Events.OnVerification.EffectiveTime > 0 {
				p.Events.OnVerification.EffectiveTime = in.Strategy.Events.OnVerification.EffectiveTime
			}
			if in.Strategy.Events.OnVerification.CombinationMode > 0 {
				p.Events.OnVerification.CombinationMode = in.Strategy.Events.OnVerification.CombinationMode
			}
		}
		if in.Strategy.Events.OnFaceLogin != nil {
			if in.Strategy.Events.OnFaceLogin.Degree > 0 && in.Strategy.Events.OnFaceLogin.Degree > 72.0 {
				p.Events.OnFaceLogin.Degree = in.Strategy.Events.OnFaceLogin.Degree
			}
		}
		if in.Strategy.Events.OnQRLogin != nil {
			if in.Strategy.Events.OnQRLogin.RefreshDuration > 0 {
				p.Events.OnQRLogin.RefreshDuration = in.Strategy.Events.OnQRLogin.RefreshDuration
			}
		}
		if in.Strategy.Events.OnRegister != nil {
			if in.Strategy.Events.OnRegister.OpenReview > 0 {
				p.Events.OnRegister.OpenReview = in.Strategy.Events.OnRegister.OpenReview
			}
			if in.Strategy.Events.OnRegister.Submitlal > 0 {
				p.Events.OnRegister.Submitlal = in.Strategy.Events.OnRegister.Submitlal
			}
			if in.Strategy.Events.OnRegister.AnDeviceRegistrationInterval > 0 {
				p.Events.OnRegister.AnDeviceRegistrationInterval = in.Strategy.Events.OnRegister.AnDeviceRegistrationInterval
			}
			if in.Strategy.Events.OnRegister.AnIPRegistrationInterval > 0 {
				p.Events.OnRegister.AnIPRegistrationInterval = in.Strategy.Events.OnRegister.AnIPRegistrationInterval
			}
			if in.Strategy.Events.OnRegister.Mode > 0 {
				p.Events.OnRegister.Mode = in.Strategy.Events.OnRegister.Mode
			}
		}
		if in.Strategy.Events.OnLogin != nil {
			if in.Strategy.Events.OnLogin.Mode != nil && len(in.Strategy.Events.OnLogin.Mode) > 0 {
				p.Events.OnLogin.Mode = in.Strategy.Events.OnLogin.Mode
			}
			if in.Strategy.Events.OnLogin.Submitlal > 0 {
				p.Events.OnLogin.Submitlal = in.Strategy.Events.OnLogin.Submitlal
			}
			if in.Strategy.Events.OnLogin.Verification > 0 {
				p.Events.OnLogin.Verification = in.Strategy.Events.OnLogin.Verification
			}
			if in.Strategy.Events.OnLogin.AllowOtherProjectUserToLogin > 0 {
				p.Events.OnLogin.AllowOtherProjectUserToLogin = in.Strategy.Events.OnLogin.AllowOtherProjectUserToLogin
			}
			if in.Strategy.Events.OnLogin.MaxCountOfOnline != nil {
				if in.Strategy.Events.OnLogin.MaxCountOfOnline.IOS > 0 {
					p.Events.OnLogin.MaxCountOfOnline.IOS = in.Strategy.Events.OnLogin.MaxCountOfOnline.IOS
				}
				if in.Strategy.Events.OnLogin.MaxCountOfOnline.Windows > 0 {
					p.Events.OnLogin.MaxCountOfOnline.Windows = in.Strategy.Events.OnLogin.MaxCountOfOnline.Windows
				}
				if in.Strategy.Events.OnLogin.MaxCountOfOnline.Web > 0 {
					p.Events.OnLogin.MaxCountOfOnline.Web = in.Strategy.Events.OnLogin.MaxCountOfOnline.Web
				}
				if in.Strategy.Events.OnLogin.MaxCountOfOnline.MacOS > 0 {
					p.Events.OnLogin.MaxCountOfOnline.MacOS = in.Strategy.Events.OnLogin.MaxCountOfOnline.MacOS
				}
				if in.Strategy.Events.OnLogin.MaxCountOfOnline.Android > 0 {
					p.Events.OnLogin.MaxCountOfOnline.Android = in.Strategy.Events.OnLogin.MaxCountOfOnline.Android
				}
			}
			if in.Strategy.Events.OnLogin.MaxCountOfInvalidAccount != nil && len(in.Strategy.Events.OnLogin.MaxCountOfInvalidAccount) > 0 {
				p.Events.OnLogin.MaxCountOfInvalidAccount = in.Strategy.Events.OnLogin.MaxCountOfInvalidAccount
			}
			if in.Strategy.Events.OnLogin.MaxCountOfErrorPassword != nil && len(in.Strategy.Events.OnLogin.MaxCountOfErrorPassword) > 0 {
				p.Events.OnLogin.MaxCountOfErrorPassword = in.Strategy.Events.OnLogin.MaxCountOfErrorPassword
			}
		}
		if in.Strategy.Events.OnInviteUser != nil {
			if in.Strategy.Events.OnInviteUser.Review > 0 {
				p.Events.OnInviteUser.Review = in.Strategy.Events.OnInviteUser.Review
			}
			if in.Strategy.Events.OnInviteUser.ExpireTime > 0 {
				p.Events.OnInviteUser.ExpireTime = in.Strategy.Events.OnInviteUser.ExpireTime
			}
		}
		if in.Strategy.Events.OnSubmitReview != nil {
			if len(in.Strategy.Events.OnSubmitReview.ReportUserId) > 0 {
				p.Events.OnSubmitReview.ReportUserId = in.Strategy.Events.OnSubmitReview.ReportUserId
			}
			if len(in.Strategy.Events.OnSubmitReview.ResponseMessage) > 0 {
				p.Events.OnSubmitReview.ResponseMessage = in.Strategy.Events.OnSubmitReview.ResponseMessage
			}
		}
		if in.Strategy.Events.OnElsewhereLogin != nil {
			if in.Strategy.Events.OnElsewhereLogin.Verification > 0 {
				p.Events.OnElsewhereLogin.Verification = in.Strategy.Events.OnElsewhereLogin.Verification
			}
			if in.Strategy.Events.OnElsewhereLogin.SendMessageToUser > 0 {
				p.Events.OnElsewhereLogin.SendMessageToUser = in.Strategy.Events.OnElsewhereLogin.SendMessageToUser
			}
		}
		if in.Strategy.Events.OnResetPassword != nil {
			if in.Strategy.Events.OnResetPassword.Email > 0 {
				p.Events.OnResetPassword.Email = in.Strategy.Events.OnResetPassword.Email
			}
			if in.Strategy.Events.OnResetPassword.Phone > 0 {
				p.Events.OnResetPassword.Phone = in.Strategy.Events.OnResetPassword.Phone
			}
		}
		if in.Strategy.Events.OnChangeOAuth != nil {
			if in.Strategy.Events.OnChangeOAuth.Verification > 0 {
				p.Events.OnChangeOAuth.Verification = in.Strategy.Events.OnChangeOAuth.Verification
			}
		}
		if in.Strategy.Events.OnChangeFace != nil {
			if in.Strategy.Events.OnChangeFace.Verification > 0 {
				p.Events.OnChangeFace.Verification = in.Strategy.Events.OnChangeFace.Verification
			}
		}
		if in.Strategy.Events.OnChangeEmail != nil {
			if in.Strategy.Events.OnChangeEmail.Verification > 0 {
				p.Events.OnChangeEmail.Verification = in.Strategy.Events.OnChangeEmail.Verification
			}
		}
		if in.Strategy.Events.OnChangePhoneNumber != nil {
			if in.Strategy.Events.OnChangePhoneNumber.Verification > 0 {
				p.Events.OnChangePhoneNumber.Verification = in.Strategy.Events.OnChangePhoneNumber.Verification
			}
		}
		if in.Strategy.Events.OnCancelFrozen != nil {
			if in.Strategy.Events.OnCancelFrozen.Verification > 0 {
				p.Events.OnCancelFrozen.Verification = in.Strategy.Events.OnCancelFrozen.Verification
			}
		}
		if in.Strategy.Events.OnRequestFrozen != nil {
			if in.Strategy.Events.OnRequestFrozen.Verification > 0 {
				p.Events.OnRequestFrozen.Verification = in.Strategy.Events.OnRequestFrozen.Verification
			}
		}
		if in.Strategy.Events.OnCommonEquipmentChanges != nil {
			if in.Strategy.Events.OnCommonEquipmentChanges.SendMessageToUser > 0 {
				p.Events.OnCommonEquipmentChanges.SendMessageToUser = in.Strategy.Events.OnCommonEquipmentChanges.SendMessageToUser
			}
		}
	}

	err = repo.Upsert(p)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)

}

func (svc *strategyService) GetRepo() repository {
	return &strategyRepository{session: svc.session.Clone()}
}

func NewService(session *mgo.Session, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &strategyService{session: session, reportercli: reportercli}
	}
	return svc
}
