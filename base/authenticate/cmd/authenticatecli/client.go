package authenticatecli

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
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
)

func NewEndpoints(tracer *zipkin.Tracer) authenticate.Endpoints {
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
		endpoints   = authenticate.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, names.F_SVC_SAFETY_AUTHENTICATE, tags, passingOnly)
	)
	{
		factory := Factory(authenticate.MakeRefreshEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.RefreshEndpoint = retry
	}
	{
		factory := Factory(authenticate.MakeNewEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.NewEndpoint = retry
	}
	{
		factory := Factory(authenticate.MakeCheckEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CheckEndpoint = retry
	}
	{
		factory := Factory(authenticate.MakeOfflineUserEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.OfflineUserEndpoint = retry
	}

	return endpoints
}

func NewClient(tracer *zipkin.Tracer) authenticate.Service {
	return NewEndpoints(tracer)
}

func Factory(makeEndpoint func(service authenticate.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := authenticate.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
