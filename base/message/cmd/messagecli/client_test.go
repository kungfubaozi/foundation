package messagecli

import (
	"testing"
	"zskparker.com/foundation/base/pb"
)

func TestNewClient(t *testing.T) {
	//zipkinTracer, reporter := serv.NewZipkin("http://192.168.2.60:9411/api/v2/spans", names.F_MQTT, "58088")
	//defer reporter.Close()
	//
	//c := NewClient("http://192.168.80.67:8500", zipkinTracer)
	//
	//resp, err := c.SendOffline(context.TODO(), &fs_base.DirectMessage{
	//	To:      "e3d03f9728df48f5913ae36813be22fa",
	//	Content: "this is broadcast message",
	//})
	//
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//
	//fmt.Println(resp)

	message, err := NewMQClient("amqp://root:123456@192.168.2.60:5672/")
	if err != nil {
		panic(err)
	}

	message.SendMessage(&fs_base.DirectMessage{To: "e3d03f9728df48f5913ae36813be22fa", Content: "hi"})
	message.SendOffline(&fs_base.DirectMessage{To: "e3d03f9728df48f5913ae36813be22fa", Content: "hi"})

}
