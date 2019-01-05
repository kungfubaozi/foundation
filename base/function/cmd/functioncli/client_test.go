package functioncli

import (
	"context"
	"fmt"
	"testing"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/serv"
)

func TestNewClient(t *testing.T) {
	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_FUNCTION, "58088")
	defer reporter.Close()

	c := NewClient(zipkinTracer)

	resp, _ := c.Get(context.Background(), &fs_base_function.GetRequest{
		Api: "/api/fds/env/entry/face",
	})

	fmt.Println(resp.State)
	fmt.Println(resp.Func.Zh)
}
