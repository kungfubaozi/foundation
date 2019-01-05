package facecli

import (
	"github.com/openzipkin/zipkin-go"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/pkg/errno"
)

func NewClient(tracer *zipkin.Tracer) face.Service {
	panic(errno.ERROR)
}
