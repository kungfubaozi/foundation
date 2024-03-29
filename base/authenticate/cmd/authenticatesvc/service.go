package authenticatesvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"runtime"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state/cmd/statecli"
	"zskparker.com/foundation/base/user/cmd/usercli"
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
	runtime.GOMAXPROCS(4)
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_AUTHENTICATE, osenv.GetMicroPortString())
	defer reporter.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	ch, err := messagecli.NewMQClient(osenv.GetMessageAMQPAddr())
	if err != nil {
		panic(err)
	}
	rp, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_AUTHENTICATE)
	defer rp.Close()

	service := authenticate.NewService(statecli.NewClient(zipkinTracer), usercli.NewClient(zipkinTracer),
		rp, pool,
		ch, fs_redisync.Create(pool), blacklistcli.NewClient(zipkinTracer))
	endpoints := authenticate.NewEndpoints(service, zipkinTracer, logger, mwclients.NewMiddleware(logger, zipkinTracer))
	svc := authenticate.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_authenticate.RegisterAuthenticateServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_AUTHENTICATE, osenv.GetConsulAddr())

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
