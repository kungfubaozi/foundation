package reporter

import (
	"context"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/analysis/statistics/cmd/statisticscli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Write(ctx context.Context, in *fs_base_reporter.WriteRequest) (*fs_base.Response, error)
}

type reporterService struct {
	session    *mgo.Session
	statistics statisticscli.Statistics
}

func (svc *reporterService) GetRepo() repository {
	return &reporterRepository{session: svc.session.Clone()}
}

func (svc *reporterService) Write(ctx context.Context, in *fs_base_reporter.WriteRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	log := &logger{
		date:      utils.FormatDate(time.Now(), utils.YYYYMMDD),
		timestamp: in.Timestamp,
		do:        in.Do,
		where:     in.Where,
		who:       in.Who,
	}

	err := repo.Write(log)

	//发送到统计里
	svc.statistics.Event(in.Who, in.Where, in.Timestamp, in.Do)

	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	return errno.ErrResponse(errno.Ok)
}

func NewService(session *mgo.Session, statistics statisticscli.Statistics) Service {
	var svc Service
	{
		svc = &reporterService{session: session, statistics: statistics}
	}
	return svc
}
