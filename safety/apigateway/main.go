package main

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"net/http"
	"os"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/safety/verification"
	"zskparker.com/foundation/safety/verification/cmd/verificationcli"
)

//safety apigateway
func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Transport domain.
	tracer := stdopentracing.GlobalTracer() // no-op
	zipkinTracer, _ := stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))
	r := mux.NewRouter()

	{
		//update
		//svc := update.MakeGRPCClient()
		//r.PathPrefix("/update").Handler(http.StripPrefix("/update", update.MakeHTTPHandler(svc, tracer, zipkinTracer, logger)))
	}

	{
		//verification
		endpoints := verificationcli.NewEndpoints(osenv.GetConsulAddr(), zipkinTracer)
		r.PathPrefix("/api/rpfds/safety/verification").Handler(http.StripPrefix("/api/rpfds/safety/verification", verification.MakeHTPPHandler(
			endpoints, tracer, zipkinTracer, logger)))
	}

	errc := make(chan error)
	// HTTP transport.
	go func() {
		errc <- http.ListenAndServe(osenv.GetMicroPortString(), r)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
