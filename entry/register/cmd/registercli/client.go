package registercli

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
	"zskparker.com/foundation/entry/register"
	"zskparker.com/foundation/entry/register/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/osenv"
)

func NewClient(tracer *zipkin.Tracer) fs_entry_register.RegisterServer {
	return NewEndpoints(tracer)
}

func NewEndpoints(tracer *zipkin.Tracer) register.Endpoints {
	var (
		retryMax     = 3
		retryTimeout = 30 * time.Second
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
		endpoints   = register.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, fs_constants.SVC_ENTRY_REGISTER, tags, passingOnly)
	)

	{
		factory := Factory(register.MakeFromOAuthEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.FromOAuthEndpoint = retry
	}

	{
		factory := Factory(register.MakeFromAPEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.FromAPEndpoint = retry
	}

	return endpoints
}

func Factory(makeEndpoint func(service register.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := register.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
