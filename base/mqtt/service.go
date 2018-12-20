package mqtt

import (
	"context"
	"github.com/eclipse/paho.mqtt.golang"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
)

type Service interface {
	SendMessage(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error)

	SendBroadcast(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error)

	SendOffline(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error)
}

type mqttService struct {
	client mqtt.Client
}

func (svc *mqttService) SendBroadcast(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error) {
	err := svc.client.Publish("broadcast", 1, false, in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *mqttService) SendOffline(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error) {
	err := svc.client.Publish("offline/"+in.To, 1, false, in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *mqttService) SendMessage(ctx context.Context, in *fs_base.TsMessage) (*fs_base.Response, error) {
	err := svc.client.Publish("user/"+in.To, 1, false, in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func NewService(client mqtt.Client) Service {
	var svc Service
	{
		svc = &mqttService{
			client: client,
		}
	}
	return svc
}
