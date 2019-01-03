package messagecli

import (
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/hashicorp/consul/api"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	"github.com/streadway/amqp"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/grpc"
	"io"
	"os"
	"time"
	"zskparker.com/foundation/base/message"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/names"
)

type Channel interface {
	SendMessage(directMessage *fs_base.DirectMessage)

	SendSMS(directMessage *fs_base.DirectMessage)

	SendEmail(directMessage *fs_base.DirectMessage)

	SendOffline(directMessage *fs_base.DirectMessage)

	SendBroadcast(directMessage *fs_base.DirectMessage)

	Close()
}

type mqMessageImpl struct {
	channel *amqp.Channel
	conn    *amqp.Connection
}

func (impl *mqMessageImpl) Close() {
	impl.conn.Close()
	impl.channel.Close()
}

func (impl *mqMessageImpl) SendMessage(directMessage *fs_base.DirectMessage) {
	impl.send("message", directMessage)
}

func (impl *mqMessageImpl) SendSMS(directMessage *fs_base.DirectMessage) {
	impl.send("sms", directMessage)
}

func (impl *mqMessageImpl) SendEmail(directMessage *fs_base.DirectMessage) {
	impl.send("email", directMessage)
}

func (impl *mqMessageImpl) SendOffline(directMessage *fs_base.DirectMessage) {
	impl.send("offline", directMessage)
}

func (impl *mqMessageImpl) SendBroadcast(directMessage *fs_base.DirectMessage) {
	impl.send("broadcast", directMessage)
}

func (impl *mqMessageImpl) send(key string, directMessage *fs_base.DirectMessage) {
	b, _ := msgpack.Marshal(directMessage)
	impl.channel.Publish("", "foundation.message", false, false, amqp.Publishing{
		Body: b,
		Type: key,
	})
}

func NewMQClient(messageAMQPAddr string) (Channel, error) {
	conn, err := amqp.Dial(messageAMQPAddr)
	if err != nil {
		fmt.Println("message connect to message queue error.")
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("message channel error.")
		return nil, err
	}

	return &mqMessageImpl{
		channel: ch,
		conn:    conn,
	}, nil
}

func NewClient(consulAddr string, tracer *zipkin.Tracer) message.Service {
	var (
		retryMax     = 3
		retryTimeout = 500 * time.Millisecond
	)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var otTracer stdopentracing.Tracer
	{
		otTracer = stdopentracing.GlobalTracer()
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
		endpoints   = message.Endpoints{}
		instancer   = consulsd.NewInstancer(client, logger, names.F_SVC_MESSAGE, tags, passingOnly)
	)
	{
		factory := Factory(message.MakeSendOfflineEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.SendOfflineEndpoint = retry
	}
	{
		factory := Factory(message.MakeSendMessageEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.SendMessageEndpoint = retry
	}
	{
		factory := Factory(message.MakeSendBroadcastEndpoint, otTracer, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.SendBroadcastEndpoint = retry
	}

	return &endpoints
}

func Factory(makeEndpoint func(service message.Service) endpoint.Endpoint, otTracer stdopentracing.Tracer, tracer *zipkin.Tracer, logger log.Logger) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		service := message.MakeGRPCClient(conn, otTracer, tracer, logger)
		e := makeEndpoint(service)
		return e, conn, nil
	}
}
