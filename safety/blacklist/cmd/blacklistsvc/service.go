package blacklistsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/middlewares/mwclients"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/blacklist/pb"
)

func StartService() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var otTracer stdopentracing.Tracer
	{
		otTracer = stdopentracing.GlobalTracer()
	}

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_BLACKLIST, osenv.GetMicroPortString())
	defer reporter.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	rc, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_BLACKLIST)
	if err != nil {
		panic(err)
	}
	defer rc.Close()

	service := blacklist.NewService(session, rc)
	endpoints := blacklist.NewEndpoints(service, zipkinTracer, logger, mwclients.NewMiddleware(logger, zipkinTracer))
	svc := blacklist.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_safety_blacklist.RegisterBlacklistServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_BLACKLIST, osenv.GetConsulAddr())

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
