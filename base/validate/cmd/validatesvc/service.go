package validatesvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/state/cmd/statecli"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/base/veds/cmd/vedscli"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_VALIDATE, osenv.GetMicroPortString())
	defer reporter.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	message, err := messagecli.NewMQClient(osenv.GetMessageAMQPAddr())
	if err != nil {
		panic(err)
	}

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	service := validate.NewService(session, message, statecli.NewClient(zipkinTracer), fs_redisync.Create(pool), vedscli.NewClient(zipkinTracer))
	endpoints := validate.NewEndpoints(service, zipkinTracer, logger)
	svc := validate.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_validate.RegisterValidateServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_VALIDATE, osenv.GetConsulAddr())

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
