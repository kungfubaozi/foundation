package interceptorcli

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"
	"io"
	"os"
	"time"
	"zskparker.com/foundation/base/interceptor"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/osenv"
)

func NewEndpoints(tracer *zipkin.Tracer) interceptor.Endpoints {
	var (
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)

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

	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = osenv.GetConsulAddr()
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}
		client = consulsd.NewClient(consulClient)
	}

	var (
		tags        []string
		passingOnly = true
		endpoints   = interceptor.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, fs_constants.SVC_INTERCEPTOR, tags, passingOnly)
	)
	{
		factory := Factory(interceptor.MakeAuthEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.AuthEndpoint = retry
	}

	return endpoints
}

func NewClient(tracer *zipkin.Tracer) interceptor.Service {
	return NewEndpoints(tracer)
}

func Factory(makeEndpoint func(service interceptor.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := interceptor.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
