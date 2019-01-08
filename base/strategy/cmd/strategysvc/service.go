package strategysvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/names"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), names.F_SVC_STRATEGY, osenv.GetMicroPortString())
	defer reporter.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	rs, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), names.F_SVC_STRATEGY)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	service := strategy.NewService(session, rs)
	endpoints := strategy.NewEndpoints(service, zipkinTracer, logger)
	svc := strategy.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_strategy.RegisterStrategyServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, names.F_SVC_STRATEGY, osenv.GetConsulAddr())

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
