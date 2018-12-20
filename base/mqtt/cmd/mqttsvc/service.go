package mqttsvc

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"os"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/utils"
)

func StartService() {
	var (
		name       = "foundation.svc.mqtt"
		consulAddr = osenv.GetConsulAddr()
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	opts := mqtt.NewClientOptions().AddBroker(os.Getenv("MQTT_BROKER")).
		SetClientID(utils.RandomMD5("mqtt_client_id")).
		SetUsername(os.Getenv("MQTT_USERNAME")).
		SetPassword(os.Getenv("MQTT_PASSWORD"))

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("mqtt server connect to broker error.")
		panic(token.Error())
	}

}
