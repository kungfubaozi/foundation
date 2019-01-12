package messagesvc

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/base/message"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/message/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
	"zskparker.com/foundation/pkg/serv"
)

func StartService() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	opts := mqtt.NewClientOptions().AddBroker(os.Getenv("MQTT_BROKER")).
		SetClientID("foundation").
		SetUsername(os.Getenv("MQTT_USERNAME")).
		SetPassword(os.Getenv("MQTT_PASSWORD"))

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("mqtt server connect to broker error.")
		panic(token.Error())
	}

	var otTracer stdopentracing.Tracer
	{
		otTracer = stdopentracing.GlobalTracer()
	}

	service := message.NewService(c, logger)
	zipkinTracer, reporter := serv.NewZipkin(osenv.GetZipkinAddr(), fs_constants.SVC_MESSAGE, osenv.GetMicroPortString())
	defer reporter.Close()
	endpoints := message.NewEndpoints(service, zipkinTracer, logger)
	srv := message.MakeGRPCServer(endpoints, otTracer, zipkinTracer, logger)

	gs := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	fs_base_message.RegisterMessageServer(gs, srv)

	errc := make(chan error)

	registration.NewRegistrar(gs, fs_constants.SVC_MESSAGE, osenv.GetConsulAddr())

	go func() {
		grpcListener, err := net.Listen("tcp", osenv.GetMicroPortString())
		if err != nil {
			fmt.Println(err)
			errc <- err
		}
		errc <- gs.Serve(grpcListener)
	}()

	go func() {
		addr := osenv.GetMessageAMQPAddr()
		if len(addr) > 0 {
			conn, err := amqp.Dial(addr)
			if err != nil {
				fmt.Println("connect to message queue error.")
				panic(err)
			}
			ch, err := conn.Channel()
			_, err = ch.QueueDeclare("foundation.message", true, true, false, false, nil)
			if err != nil {
				fmt.Println("queue error.")
				panic(err)
			}
			messages, err := ch.Consume("foundation.message", "", true, false, false, false, nil)
			if err != nil {
				fmt.Println("message queue consume error.")
				panic(err)
			}
			cli := messagecli.NewClient(osenv.GetConsulAddr(), zipkinTracer)
			for m := range messages {
				go func() {
					msg := &fs_base.DirectMessage{}
					err := msgpack.Unmarshal(m.Body, msg)
					if err == nil {
						switch m.Type {
						case "message":
							cli.SendMessage(context.Background(), msg)
							break
						case "sms":
							cli.SendSMS(context.Background(), msg)
							break
						case "email":
							cli.SendEmail(context.Background(), msg)
							break
						case "offline":
							cli.SendOffline(context.Background(), msg)
							break
						case "broadcast":
							cli.SendBroadcast(context.Background(), msg)
							break
						}
					}
				}()
			}
		}
	}()

	// Run!
	logger.Log("exit", <-errc)
}
