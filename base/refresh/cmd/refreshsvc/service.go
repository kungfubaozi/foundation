package refreshsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/authenticate/cmd/authenticatecli"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/refresh"
	"zskparker.com/foundation/base/refresh/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state/cmd/statecli"
	"zskparker.com/foundation/base/user/cmd/usercli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/pkg/sync"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_REFRESH, osenv.GetMicroPortString())
	defer reporter.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	rch, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_REFRESH)
	if err != nil {
		panic(err)
	}
	defer rch.Close()

	service := refresh.NewService(authenticatecli.NewClient(zipkinTracer), rch, logger, statecli.NewClient(zipkinTracer), fs_redisync.Create(pool), usercli.NewClient(zipkinTracer))
	endpoints := refresh.NewEndpoints(service, zipkinTracer, logger, functionmw.NewFunctionMWClient(zipkinTracer))
	svc := refresh.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_refresh.RegisterRefreshServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_REFRESH, osenv.GetConsulAddr())

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
