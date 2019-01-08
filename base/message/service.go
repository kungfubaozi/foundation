package message

import (
	"context"
	"encoding/json"
	"github.com/eclipse/paho.mqtt.golang"
	"net/http"
	"net/url"
	"time"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	SendMessage(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error)

	SendBroadcast(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error)

	SendOffline(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error)

	SendSMS(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error)

	SendEmail(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error)
}

type mqttService struct {
	client mqtt.Client
}

func (svc *mqttService) SendSMS(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	e := sms(in.To, in.Content)
	if e != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *mqttService) SendEmail(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	panic("implement me")
}

func (svc *mqttService) SendBroadcast(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	i, err := json.Marshal(in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	e := svc.client.Publish("broadcast", 1, false, string(i))
	if e.Error() != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *mqttService) SendOffline(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	i, err := json.Marshal(in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	e := svc.client.Publish("offline/"+in.To, 0, false, string(i))
	if e.Error() != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *mqttService) SendMessage(ctx context.Context, in *fs_base.DirectMessage) (*fs_base.Response, error) {
	i, err := json.Marshal(in)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	e := svc.client.Publish("message/"+in.To, 0, false, string(i))
	if e.Error() != nil {
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

func sms(mobile, c string) error {
	appId := "EUCP-EMY-SMS0-JBZOQ"
	secretKey := "8553947376603211"
	timestamp := utils.FormatDate(time.Now(), utils.YYYYMMDDHHMMSS)
	sign := utils.Md5(appId + secretKey + timestamp)
	values := url.Values{}
	values.Add("appId", appId)
	values.Add("timestamp", timestamp)
	values.Add("sign", sign)
	values.Add("mobiles", mobile)
	values.Add("content", c)
	_, err := http.PostForm("http://shmtn.b2m.cn:80/simpleinter/sendSMS", values)
	return err
}
