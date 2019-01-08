package main

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"net/http"
	"os"
	"zskparker.com/foundation/entry/register"
	"zskparker.com/foundation/entry/register/cmd/registercli"
	"zskparker.com/foundation/pkg/osenv"
)

//login register apigateway
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
		//register
		endpoints := registercli.NewEndpoints(zipkinTracer)
		r.PathPrefix(register.GetAdminFunc().Prefix).Handler(http.StripPrefix(register.GetAdminFunc().Prefix,
			register.MakeHTTPHandler(endpoints, tracer, zipkinTracer, logger)))
	}

	errc := make(chan error)
	// HTTP transport.
	go func() {
		errc <- http.ListenAndServe(osenv.GetMicroPortString(), r)
	}()

	// Run!
	logger.Log("exit", <-errc)
}
