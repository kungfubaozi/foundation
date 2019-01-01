package main

import (
	"context"
	"github.com/go-kit/kit/examples/addsvc/pkg/addtransport"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"net/http"
	"os"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/safety/update"
)

//safety apigateway
func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Service discovery domain. In this example we use Consul.
	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(*consulAddr) > 0 {
			consulConfig.Address = *consulAddr
		}
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	// Transport domain.
	tracer := stdopentracing.GlobalTracer() // no-op
	zipkinTracer, _ := stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))
	ctx := context.Background()
	r := mux.NewRouter()

	//update
	r.PathPrefix("/update").Handler(http.StripPrefix("/update", update.MakeHTTPHandler()))

	errc := make(chan error)
	// HTTP transport.
	go func() {
		errc <- http.ListenAndServe(osenv.GetMicroPortString(), r)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
