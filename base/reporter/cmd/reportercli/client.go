package reportercli

import (
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/hashicorp/consul/api"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"io"
	"os"
	"time"
	"zskparker.com/foundation/base/reporter"
	"zskparker.com/foundation/base/reporter/pb"
	"zskparker.com/foundation/pkg/constants"
)

type Channel interface {
	Write(function, who, where string, status int64)

	Close()
}

type channel struct {
	channel *amqp.Channel
	conn    *amqp.Connection
	svc     string
}

func (c *channel) Write(function, who, where string, status int64) {
	b, _ := msgpack.Marshal(&fs_base_reporter.WriteRequest{
		Svc:       c.svc,
		Func:      function,
		Who:       who,
		Where:     where,
		Status:    status,
		Timestamp: time.Now().UnixNano(),
	})
	c.channel.Publish("", "foundation.reporter", false, false, amqp.Publishing{
		Body: b,
		Type: "reporter",
	})
}

func (c *channel) Close() {
	c.conn.Close()
	c.channel.Close()
}

func NewMQConnect(reporterAMQTAddr string, svc string) (Channel, error) {
	conn, err := amqp.Dial(reporterAMQTAddr)
	if err != nil {
		fmt.Println("message connect to message queue error.")
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("message channel error.")
		return nil, err
	}
	return &channel{conn: conn, channel: ch, svc: svc}, nil
}

func NewClient(consulAddr string) reporter.Service {

	var (
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)

	var l log.Logger
	{
		l = log.NewLogfmtLogger(os.Stderr)
		l = log.With(l, "ts", log.DefaultTimestampUTC)
		l = log.With(l, "caller", log.DefaultCaller)
	}

	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(consulAddr) > 0 {
			consulConfig.Address = consulAddr
		}
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}
		client = consulsd.NewClient(consulClient)
	}

	var (
		tags        []string
		passingOnly = true
		endpoints   = reporter.Endpoints{}
		instancer   = consulsd.NewInstancer(client, l, fs_constants.SVC_REPORTER, tags, passingOnly)
	)
	{
		factory := Factory(reporter.MakeWriteEndpoint)
		endpointer := sd.NewEndpointer(instancer, factory, l)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.WriteEndpoint = retry
	}

	return endpoints
}

func Factory(makeEndpoint func(service reporter.Service) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := reporter.MakeGRPCClient(conn)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
