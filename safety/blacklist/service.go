package blacklist

import (
	"context"
	"gopkg.in/mgo.v2"
	"sync"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Service interface {
	CheckMeta(ctx context.Context, in *fs_safety_blacklist.CheckMetaRequest) (*fs_base.Response, error)

	CheckAccount(ctx context.Context, in *fs_safety_blacklist.CheckAccountRequest) (*fs_base.Response, error)

	Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error)
}

type blacklistService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
}

func (svc *blacklistService) GetRepo() repository {
	return &blacklistRepository{session: svc.session.Clone()}
}

func (svc *blacklistService) CheckAccount(ctx context.Context, in *fs_safety_blacklist.CheckAccountRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	err := repo.Get(in.Account, ACCOUNT)
	if err != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}
	fs_metadata_transport.MetaToReporterByMetadata(svc.reportercli, in.Meta, in.Account, fs_function_tags.GetBlacklistCheckIP(), fs_constants.STATUS_OK)
	return errno.ErrResponse(errno.Ok)
}

func (svc *blacklistService) Add(ctx context.Context, in *fs_safety_blacklist.AddRequest) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *blacklistService) CheckMeta(ctx context.Context, in *fs_safety_blacklist.CheckMetaRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	var wg sync.WaitGroup

	ps := errno.Ok
	errc := func(s *fs_base.State) {
		if ps.Ok {
			ps = s
		}
		wg.Done()
	}

	if len(in.Ip) > 0 {
		wg.Add(1)
		go func() {
			err := repo.Get(in.Ip, IP)
			if err == nil {
				errc(errno.Ok)
			} else {
				fs_metadata_transport.MetaToReporterByMetadata(svc.reportercli, in.Meta, in.Ip, fs_function_tags.GetBlacklistCheckIP(), fs_constants.STATUS_OK)
				errc(errno.ErrRequest)
			}
		}()
	}

	if len(in.Device) > 0 {
		wg.Add(1)
		go func() {
			err := repo.Get(in.Ip, DEVICE)
			if err == nil {
				errc(errno.Ok)
			} else {
				fs_metadata_transport.MetaToReporterByMetadata(svc.reportercli, in.Meta, in.Device, fs_function_tags.GetBlacklistCheckDevice(), fs_constants.STATUS_OK)
				errc(errno.ErrRequest)
			}
		}()
	}

	if len(in.UserId) > 0 {
		wg.Add(1)
		go func() {
			err := repo.Get(in.UserId, USER)
			if err == nil {
				errc(errno.Ok)
			} else {
				fs_metadata_transport.MetaToReporterByMetadata(svc.reportercli, in.Meta, in.UserId, fs_function_tags.GetBlacklistCheckUser(), fs_constants.STATUS_OK)
				errc(errno.ErrRequest)
			}
		}()
	}

	wg.Wait()

	return errno.ErrResponse(ps)
}

func NewService(session *mgo.Session, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &blacklistService{session: session, reportercli: reportercli}
	}
	return svc
}
