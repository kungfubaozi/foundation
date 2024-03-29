package main

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"net/http"
	"os"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/safety/blacklist"
	"zskparker.com/foundation/safety/blacklist/cmd/blacklistcli"
	"zskparker.com/foundation/safety/update"
	"zskparker.com/foundation/safety/update/cmd/updatecli"
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
		endpoints := updatecli.NewEndpoints(zipkinTracer)
		r.PathPrefix(fs_functions.GetUpdateEmailFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetUpdateEmailFunc().Prefix, update.MakeHTTPHandler(
			endpoints, tracer, zipkinTracer, logger)))
	}

	{
		//verification
		endpoints := verificationcli.NewEndpoints(osenv.GetConsulAddr(), zipkinTracer)
		r.PathPrefix(fs_functions.GetVerificationRegisterFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetVerificationRegisterFunc().Prefix, verification.MakeHTTPHandler(
			endpoints, tracer, zipkinTracer, logger)))
	}

	{
		//blacklist
		endpoints := blacklistcli.NewEndpoint(zipkinTracer)
		r.PathPrefix(fs_functions.GetAddBlacklistFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetAddBlacklistFunc().Prefix, blacklist.MakeHTTPHandler(
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
