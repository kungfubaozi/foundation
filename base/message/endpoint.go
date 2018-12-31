package message

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/pb"
)

type Endpoints struct {
	SendBroadcastEndpoint endpoint.Endpoint
	SendMessageEndpoint   endpoint.Endpoint
	SendOfflineEndpoint   endpoint.Endpoint
	SendSMSEndpoint       endpoint.Endpoint
	SendEmailEndpoint     endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var sendBroadcastEndpoint endpoint.Endpoint
	{
		sendBroadcastEndpoint = MakeSendBroadcastEndpoint(svc)
		sendBroadcastEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendBroadcastEndpoint)
		sendBroadcastEndpoint = zipkin.TraceEndpoint(trace, "SendBroadcast")(sendBroadcastEndpoint)
	}

	var sendMessageEndpoint endpoint.Endpoint
	{
		sendMessageEndpoint = MakeSendMessageEndpoint(svc)
		sendMessageEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendMessageEndpoint)
		sendMessageEndpoint = zipkin.TraceEndpoint(trace, "SendMessage")(sendMessageEndpoint)
	}

	var sendOfflineEndpoint endpoint.Endpoint
	{
		sendOfflineEndpoint = MakeSendOfflineEndpoint(svc)
		sendOfflineEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendOfflineEndpoint)
		sendOfflineEndpoint = zipkin.TraceEndpoint(trace, "SendOffline")(sendOfflineEndpoint)
	}

	var sendEmailEndpoint endpoint.Endpoint
	{
		sendEmailEndpoint = MakeSendEmailEndpoint(svc)
		sendEmailEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendEmailEndpoint)
		sendEmailEndpoint = zipkin.TraceEndpoint(trace, "SendEmail")(sendEmailEndpoint)
	}

	var sendSMSEndpoint endpoint.Endpoint
	{
		sendSMSEndpoint = MakeSendSMSEndpoint(svc)
		sendSMSEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sendSMSEndpoint)
		sendSMSEndpoint = zipkin.TraceEndpoint(trace, "SendSMS")(sendSMSEndpoint)
	}

	return Endpoints{
		SendBroadcastEndpoint: sendBroadcastEndpoint,
		SendMessageEndpoint:   sendMessageEndpoint,
		SendOfflineEndpoint:   sendOfflineEndpoint,
		SendSMSEndpoint:       sendSMSEndpoint,
		SendEmailEndpoint:     sendEmailEndpoint,
	}
}

func (g *Endpoints) SendMessage(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	resp, err := g.SendMessageEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *Endpoints) SendBroadcast(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	resp, err := g.SendBroadcastEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *Endpoints) SendOffline(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	resp, err := g.SendOfflineEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *Endpoints) SendEmail(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	resp, err := g.SendEmailEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g *Endpoints) SendSMS(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	resp, err := g.SendSMSEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeSendBroadcastEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.SendBroadcast(ctx, request.(*fs_base.DirectMessage))
	}
}

func MakeSendMessageEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.SendMessage(ctx, request.(*fs_base.DirectMessage))
	}
}

func MakeSendOfflineEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.SendOffline(ctx, request.(*fs_base.DirectMessage))
	}
}

func MakeSendSMSEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.SendSMS(ctx, request.(*fs_base.DirectMessage))
	}
}

func MakeSendEmailEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.SendEmail(ctx, request.(*fs_base.DirectMessage))
	}
}
