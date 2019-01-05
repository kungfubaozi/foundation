package projectcli

import (
	"github.com/openzipkin/zipkin-go"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/pkg/errno"
)

func NewClient(tracer *zipkin.Tracer) project.Service {
	panic(errno.ERROR)
}
