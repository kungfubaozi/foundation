package interceptorsvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/interceptor"
	"zskparker.com/foundation/base/interceptor/pb"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_INTERCEPTOR, osenv.GetMicroPortString())
	defer reporter.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	ch, err := messagecli.NewMQClient(osenv.GetMessageAMQPAddr())
	if err != nil {
		panic(err)
	}
	rp, err := reportercli.NewMQConnect(osenv.GetReporterAMQPAddr(), fs_constants.SVC_INTERCEPTOR)
	defer rp.Close()

	service := interceptor.NewService(rp, ch)
	endpoints := interceptor.NewEndpoints(service, zipkinTracer, logger, mwclients.NewMiddleware(logger, zipkinTracer))
	svc := interceptor.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_interceptor.RegisterInterceptorServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_INTERCEPTOR, osenv.GetConsulAddr())

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
