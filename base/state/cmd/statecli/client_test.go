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

	c := NewClient(zipkinTracer)

	ctx := context.Background()

	resp, err := c.Upsert(ctx, &fs_base_state.UpsertRequest{
		Key:    "1111",
		Status: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.State)
}
