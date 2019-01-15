package main

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"net/http"
	"os"
	"zskparker.com/foundation/base/interceptor"
	"zskparker.com/foundation/base/interceptor/cmd/interceptorcli"
	"zskparker.com/foundation/base/invite"
	"zskparker.com/foundation/base/invite/cmd/invitecli"
	"zskparker.com/foundation/base/refresh"
	"zskparker.com/foundation/base/refresh/cmd/refreshcli"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/osenv"
)

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
		endpoints := interceptorcli.NewEndpoints(zipkinTracer)
		r.PathPrefix(fs_functions.GetInterceptFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetInterceptFunc().Prefix, interceptor.MakeHTTPHandler(
			endpoints, tracer, zipkinTracer, logger)))
	}

	{
		endpoints := refreshcli.NewEndpoints(zipkinTracer)
		r.PathPrefix(fs_functions.GetRefreshFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetRefreshFunc().Prefix, refresh.MakeHTTPHandler(
			endpoints, tracer, zipkinTracer, logger)))
	}

	{
		endpoints := invitecli.NewEndpoint(zipkinTracer)
		r.PathPrefix(fs_functions.GetInviteUserFunc().Prefix).Handler(http.StripPrefix(fs_functions.GetInviteUserFunc().Prefix, invite.MakeHTTPHandler(
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
