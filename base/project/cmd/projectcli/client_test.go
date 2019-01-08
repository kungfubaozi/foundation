package projectcli

import (
	"context"
	"fmt"
	"testing"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/serv"
)

func TestNewClient(t *testing.T) {
	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_PROJECT, "58089")
	defer reporter.Close()

	c := NewClient(zipkinTracer)

	resp, err := c.New(context.Background(), &fs_base_project.NewRequest{
		Desc: "this is test project",
		En:   "test",
		Zh:   "test",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.State)
}
