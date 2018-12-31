package serv

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func NewZipkin(zipkinV2URL, serviceName, hostPort string) (*zipkin.Tracer, reporter.Reporter) {
	r := zipkinhttp.NewReporter(zipkinV2URL)
	zEP, _ := zipkin.NewEndpoint(serviceName, hostPort)
	zipkinTracer, err := zipkin.NewTracer(
		r, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(false),
	)
	if err != nil {
		panic(err)
	}
	return zipkinTracer, r
}
