package registersvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/face/cmd/facecli"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user/cmd/usercli"
	"zskparker.com/foundation/entry/register"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/middlewares/mwclients"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/safety/blacklist/cmd/blacklistcli"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_ENTRY_REGISTER, osenv.GetMicroPortString())
	defer reporter.Close()

	rs, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_ENTRY_REGISTER)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	service := register.NewService(usercli.NewClient(zipkinTracer), rs, facecli.NewClient(zipkinTracer), fs_redisync.Create(pool),
		blacklistcli.NewClient(zipkinTracer))
	endpoints := register.NewEndpoints(service, zipkinTracer, logger, mwclients.NewMiddleware(logger, zipkinTracer))
	svc := register.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_entry_register.RegisterRegisterServer(gs, svc)

	registration.NewRegistrar(gs, fs_constants.SVC_ENTRY_REGISTER, osenv.GetConsulAddr())

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
