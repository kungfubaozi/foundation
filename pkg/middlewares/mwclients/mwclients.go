package mwclients

import (
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"zskparker.com/foundation/base/authenticate/cmd/authenticatecli"
	"zskparker.com/foundation/base/face/cmd/facecli"
	"zskparker.com/foundation/base/function/cmd/functioncli"
	"zskparker.com/foundation/base/project/cmd/projectcli"
	"zskparker.com/foundation/base/validate/cmd/validatecli"
	"zskparker.com/foundation/pkg/middlewares"
	"zskparker.com/foundation/safety/blacklist/cmd/blacklistcli"
)

func NewMiddleware(logger log.Logger, tracer *zipkin.Tracer) fs_endpoint_middlewares.Endpoint {
	return fs_endpoint_middlewares.Create(logger, functioncli.NewClient(tracer),
		authenticatecli.NewClient(tracer),
		facecli.NewClient(tracer),
		validatecli.NewClient(tracer),
		projectcli.NewClient(tracer), blacklistcli.NewClient(tracer))
}
