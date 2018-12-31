package reportersvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/analysis/statistics/cmd/statisticscli"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/reporter/pb"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
)

func StartService() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	s, err := statisticscli.NewConnect(osenv.GetStatisticsKafkaAddr())
	if err != nil {
		panic(err)
	}
	service := reporter.NewService(session, s)
	endpoints := reporter.NewEndpoints(service)
	svc := reporter.MakeGRPCServer(endpoints, logger)

	gs := grpc.NewServer()
	fs_base_reporter.RegisterReporterServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, names.F_SVC_REPORTER, osenv.GetConsulAddr())

	go func() {
		grpcListener, err := net.Listen("tcp", osenv.GetMicroPortString())
		if err != nil {
			fmt.Println(err)
			errc <- err
		}
		errc <- gs.Serve(grpcListener)
	}()

	logger.Log("exit", <-errc)
}
