package loginsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/authenticate/cmd/authenticatecli"
	"zskparker.com/foundation/base/face/cmd/facecli"
	"zskparker.com/foundation/base/function/cmd/functionmw"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user/cmd/usercli"
	"zskparker.com/foundation/base/validate/cmd/validatecli"
	"zskparker.com/foundation/entry/login"
	"zskparker.com/foundation/entry/login/pb"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/names"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), names.F_SVC_ENTRY_LOGIN, osenv.GetMicroPortString())
	defer reporter.Close()

	rs, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), names.F_SVC_ENTRY_LOGIN)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	service := login.NewService(usercli.NewClient(zipkinTracer), rs, authenticatecli.NewClient(zipkinTracer), validatecli.NewClient(zipkinTracer),
		facecli.NewClient(zipkinTracer), fs_redisync.Create(pool))
	endpoints := login.NewEndpoints(service, zipkinTracer, logger, functionmw.NewFunctionMWClient(zipkinTracer))
	svc := login.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_entry_login.RegisterLoginServer(gs, svc)

	registration.NewRegistrar(gs, names.F_SVC_ENTRY_LOGIN, osenv.GetConsulAddr())

	errc := make(chan error)

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
