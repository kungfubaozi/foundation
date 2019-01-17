package strategycli

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
	"zskparker.com/foundation/base/strategy"
	"zskparker.com/foundation/base/strategy/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/osenv"
)

func NewEndpoints(tracer *zipkin.Tracer) strategy.Endpoints {
	var (
		retryMax     = 3
		retryTimeout = 3000 * time.Millisecond
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
		endpoints   = strategy.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, fs_constants.SVC_STRATEGY, tags, passingOnly)
	)

	{
		factory := Factory(strategy.MakeGetEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetEndpoint = retry
	}

	{
		factory := Factory(strategy.MakeUpsertEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.UpsertEndpoint = retry
	}

	return endpoints
}

func NewClient(tracer *zipkin.Tracer) fs_base_strategy.StrategyServer {
	return NewEndpoints(tracer)
}

func Factory(makeEndpoint func(service strategy.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := strategy.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
