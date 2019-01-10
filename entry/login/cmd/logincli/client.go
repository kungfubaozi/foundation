package logincli

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
	"zskparker.com/foundation/entry/login"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
)

func NewClient(tracer *zipkin.Tracer) login.Service {
	return NewEndpoints(tracer)
}

func NewEndpoints(tracer *zipkin.Tracer) login.Endpoints {
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
		endpoints   = login.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, names.F_SVC_ENTRY_LOGIN, tags, passingOnly)
	)

	{
		factory := Factory(login.MakeEntryByQRCodeEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.EntryByQRCodeEndpoint = retry
	}

	{
		factory := Factory(login.MakeEntryByValidateCodeEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.EntryByValidateCodeEndpoint = retry
	}

	{
		factory := Factory(login.MakeEntryByOAuthEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.EntryByOAuthEndpoint = retry
	}

	{
		factory := Factory(login.MakeEntryByFaceEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.EntryByFaceEndpoint = retry
	}

	{
		factory := Factory(login.MakeEntryByAPEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.EntryByAPEndpoint = retry
	}

	return endpoints
}

func Factory(makeEndpoint func(service login.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := login.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
