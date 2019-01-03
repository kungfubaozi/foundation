package authenticatesvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state/cmd/statecli"
	"zskparker.com/foundation/base/user/cmd/usercli"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), names.F_SVC_SAFETY_AUTHENTICATE, osenv.GetMicroPortString())
	defer reporter.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	ch, err := messagecli.NewMQClient(osenv.GetMessageAMQPAddr())
	if err != nil {
		panic(err)
	}
	rp, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), names.F_SVC_SAFETY_AUTHENTICATE)
	defer rp.Close()

	service := authenticate.NewService(statecli.NewClient(osenv.GetConsulAddr(), zipkinTracer), usercli.NewClient(zipkinTracer),
		rp, pool,
		ch)
	endpoints := authenticate.NewEndpoints(service, zipkinTracer, logger)
	svc := authenticate.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_authenticate.RegisterAuthenticateServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, names.F_SVC_SAFETY_AUTHENTICATE, osenv.GetConsulAddr())

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
