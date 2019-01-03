package authenticatecli

import (
	"context"
	"fmt"
	"testing"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/serv"
)

func TestNewClient(t *testing.T) {
	zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_SVC_SAFETY_AUTHENTICATE, "58082")
	defer reporter.Close()

	c := NewClient(zipkinTracer)
	resp, err := c.New(context.Background(), &fs_base_authenticate.NewRequest{
		MaxOnlineCount: 1,
		Authorize: &fs_base_authenticate.Authorize{
			UserId:    "32c2e9c2-1523-4f5b-85d1-d2f9b97d608f",
			ProjectId: "025c3b8c-a717-43fe-9307-7c721a1c07b6",
			ClientId:  "f41d5548-1292-468e-835f-e25be694189b",
			Device:    "test",
			Platform:  names.F_PLATFORM_ANDROID,
			UserAgent: "this is useragent",
			Ip:        "192.168.80.60",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.State)
	fmt.Println(resp.AccessToken)
	fmt.Println(resp.RefreshToken)
	fmt.Println(resp.Session)
}
