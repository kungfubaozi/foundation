package projectsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy/cmd/strategycli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/middlewares/mwclients"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_PROJECT, osenv.GetMicroPortString())
	defer reporter.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	rs, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_PROJECT)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	//插入默认项目
	//project.InsertDef(session)

	service := project.NewService(session, strategycli.NewClient(zipkinTracer), rs)

	endpoints := project.NewEndpoints(service, zipkinTracer, logger, mwclients.NewMiddleware(logger, zipkinTracer))
	svc := project.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_project.RegisterProjectServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_PROJECT, osenv.GetConsulAddr())

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
