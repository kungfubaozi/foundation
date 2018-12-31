package message

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
	"zskparker.com/foundation/base/message/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/format"
)

type GRPCServer struct {
	sendBroadcast grpctransport.Handler
	sendMessage   grpctransport.Handler
	sendOffline   grpctransport.Handler
	sendSMS       grpctransport.Handler
	sendEmail     grpctransport.Handler
}

func MakeGRPCServer(endpoints Endpoints, otTracer stdopentracing.Tracer, tracer *stdzipkin.Tracer, logger log.Logger) fs_base_message.MessageServer {
	zipkinServer := zipkin.GRPCServerTrace(tracer)

	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		zipkinServer,
	}

	return &GRPCServer{
		sendBroadcast: grpctransport.NewServer(
			endpoints.SendBroadcastEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "SendBroadcast", logger)))...),
		sendMessage: grpctransport.NewServer(
			endpoints.SendMessageEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "SendMessage", logger)))...),
		sendOffline: grpctransport.NewServer(
			endpoints.SendOfflineEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "SendOffline", logger)))...),
		sendEmail: grpctransport.NewServer(
			endpoints.SendEmailEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "SendEmail", logger)))...),
		sendSMS: grpctransport.NewServer(
			endpoints.SendSMSEndpoint,
			format.GrpcMessage,
			format.GrpcMessage,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer, "SendSMS", logger)))...),
	}
}

func MakeGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) Service {
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	zipkinClient := zipkin.GRPCClientTrace(zipkinTracer)

	options := []grpctransport.ClientOption{
		zipkinClient,
		grpctransport.ClientBefore(format.GRPCMetadata()),
	}

	var sendBroadcastEndpoint endpoint.Endpoint
	{
		sendBroadcastEndpoint = grpctransport.NewClient(conn,
			"fs.base.message.Message",
			"SendBroadcast",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		sendBroadcastEndpoint = limiter(sendBroadcastEndpoint)
		sendBroadcastEndpoint = opentracing.TraceClient(otTracer, "SendBroadcast")(sendBroadcastEndpoint)
		sendBroadcastEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "SendBroadcast",
			Timeout: 5 * time.Second,
		}))(sendBroadcastEndpoint)
	}

	var sendOfflineEndpoint endpoint.Endpoint
	{
		sendOfflineEndpoint = grpctransport.NewClient(conn,
			"fs.base.message.Message",
			"SendOffline",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		sendOfflineEndpoint = limiter(sendOfflineEndpoint)
		sendOfflineEndpoint = opentracing.TraceClient(otTracer, "SendBroadcast")(sendOfflineEndpoint)
		sendOfflineEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "SendOffline",
			Timeout: 5 * time.Second,
		}))(sendOfflineEndpoint)
	}

	var sendMessageEndpoint endpoint.Endpoint
	{
		sendMessageEndpoint = grpctransport.NewClient(conn,
			"fs.base.message.Message",
			"SendMessage",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		sendMessageEndpoint = limiter(sendMessageEndpoint)
		sendMessageEndpoint = opentracing.TraceClient(otTracer, "SendMessage")(sendMessageEndpoint)
		sendMessageEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "SendMessage",
			Timeout: 5 * time.Second,
		}))(sendMessageEndpoint)
	}

	var sendSMSEndpoint endpoint.Endpoint
	{
		sendSMSEndpoint = grpctransport.NewClient(conn,
			"fs.base.message.Message",
			"SendSMS",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		sendSMSEndpoint = limiter(sendSMSEndpoint)
		sendSMSEndpoint = opentracing.TraceClient(otTracer, "SendSMS")(sendSMSEndpoint)
		sendSMSEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "SendSMS",
			Timeout: 5 * time.Second,
		}))(sendSMSEndpoint)
	}

	var sendEmailEndpoint endpoint.Endpoint
	{
		sendEmailEndpoint = grpctransport.NewClient(conn,
			"fs.base.message.Message",
			"SendEmail",
			format.GrpcMessage,
			format.GrpcMessage,
			fs_base.Response{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...).Endpoint()
		sendEmailEndpoint = limiter(sendEmailEndpoint)
		sendEmailEndpoint = opentracing.TraceClient(otTracer, "SendEmail")(sendEmailEndpoint)
		sendEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "SendEmail",
			Timeout: 5 * time.Second,
		}))(sendEmailEndpoint)
	}

	return &Endpoints{
		SendOfflineEndpoint:   sendOfflineEndpoint,
		SendMessageEndpoint:   sendMessageEndpoint,
		SendBroadcastEndpoint: sendBroadcastEndpoint,
		SendEmailEndpoint:     sendEmailEndpoint,
		SendSMSEndpoint:       sendSMSEndpoint,
	}
}

func (g *GRPCServer) SendMessage(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	_, resp, err := g.sendMessage.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) SendBroadcast(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	_, resp, err := g.sendBroadcast.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) SendOffline(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	_, resp, err := g.sendOffline.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) SendSMS(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	_, resp, err := g.sendSMS.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *GRPCServer) SendEmail(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	_, resp, err := g.sendEmail.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}
