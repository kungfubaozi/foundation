package statecli

import (
	"context"
	"fmt"
	"testing"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/serv"
)

func TestNewClient(t *testing.T) {

	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_STATE, "58092")
	defer reporter.Close()

	c := NewClient("http://192.168.80.67:8500", zipkinTracer)

	ctx := context.Background()

	resp, err := c.Get(ctx, &fs_base_state.GetRequest{
		Key: "1111",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.State.Message)
}
