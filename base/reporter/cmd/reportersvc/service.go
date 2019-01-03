package reportersvc

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"net"
	"os"
	"zskparker.com/foundation/analysis/statistics/cmd/statisticscli"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/reporter/pb"
	"zskparker.com/foundation/pkg/db"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/registration"
)

func StartService() {

	var l log.Logger
	{
		l = log.NewLogfmtLogger(os.Stderr)
		l = log.With(l, "ts", log.DefaultTimestampUTC)
		l = log.With(l, "caller", log.DefaultCaller)
	}

	session, err := db.CreateSession(osenv.GetMongoDBAddr())
	if err != nil {
		panic(err)
	}
	defer session.Close()

	s, err := statisticscli.NewConnect(osenv.GetStatisticsKafkaAddr())
	if err != nil {
		panic(err)
	}
	service := reporter.NewService(session, s)
	endpoints := reporter.NewEndpoints(service)
	svc := reporter.MakeGRPCServer(endpoints)

	gs := grpc.NewServer()
	fs_base_reporter.RegisterReporterServer(gs, svc)

	errc := make(chan error)

	registration.NewRegistrar(gs, names.F_SVC_REPORTER, osenv.GetConsulAddr())

	go func() {
		grpcListener, err := net.Listen("tcp", osenv.GetMicroPortString())
		if err != nil {
			fmt.Println(err)
			errc <- err
		}
		errc <- gs.Serve(grpcListener)
	}()

	go func() {
		addr := osenv.GetReporterAMQPAddr()
		if len(addr) > 0 {
			conn, err := amqp.Dial(addr)
			if err != nil {
				fmt.Println("connect to message queue error.")
				panic(err)
			}
			ch, err := conn.Channel()
			_, err = ch.QueueDeclare("foundation.reporter", true, true, false, false, nil)
			if err != nil {
				fmt.Println("queue error.")
				panic(err)
			}
			messages, err := ch.Consume("foundation.reporter", "", true, false, false, false, nil)
			if err != nil {
				fmt.Println("message queue consume error.")
				panic(err)
			}
			cli := reportercli.NewClient(osenv.GetConsulAddr())
			for m := range messages {
				go func() {
					msg := &fs_base_reporter.WriteRequest{}
					err := msgpack.Unmarshal(m.Body, msg)
					if err == nil {
						cli.Write(context.Background(), msg)
					}
				}()
			}
		}
	}()

	l.Log("exit", <-errc)
}
