package verificationcli

import (
	"context"
	"fmt"
	"testing"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/serv"
	"zskparker.com/foundation/safety/verification/pb"
)

func TestNewClient(t *testing.T) {

	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_SAFETY_VERIFICATION, "58121")
	defer reporter.Close()

	c := NewClient("http://192.168.80.67:8500", zipkinTracer)
	resp, err := c.New(context.Background(), &fs_safety_verification.NewRequest{
		Do: names.F_DO_REGISTER,
		To: "13222021207",
	})
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(resp)

}
