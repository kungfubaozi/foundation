package vedssvc

import (
	"fmt"
	"github.com/go-kit/kit/log"
	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/veds"
	"zskparker.com/foundation/base/veds/pb"
	"zskparker.com/foundation/pkg/constants"
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

	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_VEDS, osenv.GetMicroPortString())
	defer reporter.Close()

	key := osenv.GetSecretKey()
	if len(key) != 32 {
		panic("The veds secret key length must = 32")
	}

	service := veds.NewService(key)
	endpoints := veds.NewEndpoints(service, zipkinTracer, logger)
	svc := veds.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer()
	fs_base_veds.RegisterVEDSServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_VEDS, osenv.GetConsulAddr())

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
