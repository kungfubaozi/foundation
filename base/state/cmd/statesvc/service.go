package statesvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/db"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_STATE, osenv.GetMicroPortString())
	defer reporter.Close()

	pool := db.CreatePool(osenv.GetRedisAddr())
	defer pool.Close()

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	service := state.NewService(pool, session)
	endpoints := state.NewEndpoints(service, zipkinTracer, logger)
	svc := state.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_state.RegisterStateServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_STATE, osenv.GetConsulAddr())

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
