package vedscli

import (
	"fmt"
	"testing"
	"time"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/pkg/tool/veds"
)

func TestNewClient(t *testing.T) {

	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", fs_constants.SVC_VEDS, "58099")
	defer reporter.Close()

	c := NewClient(zipkinTracer)

	fmt.Println(time.Now().Unix())

	v := fs_service_veds.Encrypt(c, "this is test message", "this is test 2 message", "this is test 3 message",
		"this is test 3 message", "this is test 3 message", "this is test 3 message", "this is test 3 message", "this is test 3 message", "this is test 3 message",
		"this is test 3 message", "this is test 3 message", "this is test 3 message", "this is test 3 message", "this is test 3 message")

	fmt.Println(v.State)

	fmt.Println(v.Values)

	v1 := fs_service_veds.Decrypt(c, v.Values)
	fmt.Println(v1.State)
	fmt.Println(v1.Values)

	fmt.Println(time.Now().Unix())

}
