package login

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/entry/login/pb"
	"zskparker.com/foundation/pkg/middlewares"
)

type Endpoints struct {
	EntryByAPEndpoint           endpoint.Endpoint
	EntryByFaceEndpoint         endpoint.Endpoint
	EntryByOAuthEndpoint        endpoint.Endpoint
	EntryByValidateCodeEndpoint endpoint.Endpoint
	EntryByQRCodeEndpoint       endpoint.Endpoint
	EntryByInviteEndpoint       endpoint.Endpoint
}

func (g Endpoints) EntryByInvite(ctx context.Context, in *fs_entry_login.EntryByInviteRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByInviteEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g Endpoints) EntryByAP(ctx context.Context, in *fs_entry_login.EntryByAPRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByAPEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g Endpoints) EntryByOAuth(ctx context.Context, in *fs_entry_login.EntryByOAuthRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByOAuthEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g Endpoints) EntryByValidateCode(ctx context.Context, in *fs_entry_login.EntryByValidateCodeRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByValidateCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g Endpoints) EntryByQRCode(ctx context.Context, in *fs_entry_login.EntryByQRCodeRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByQRCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func (g Endpoints) EntryByFace(ctx context.Context, in *fs_entry_login.EntryByFaceRequest) (*fs_entry_login.EntryResponse, error) {
	resp, err := g.EntryByFaceEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_entry_login.EntryResponse), nil
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger, client fs_endpoint_middlewares.Endpoint) Endpoints {

	var entryByAPEndpoint endpoint.Endpoint
	{
		entryByAPEndpoint = MakeEntryByAPEndpoint(svc)
		entryByAPEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByAPEndpoint)
		entryByAPEndpoint = zipkin.TraceEndpoint(trace, "EntryByAp")(entryByAPEndpoint)

		entryByAPEndpoint = client.WithMeta()(entryByAPEndpoint)
	}

	var entryByFaceEndpoint endpoint.Endpoint
	{
		entryByFaceEndpoint = MakeEntryByFaceEndpoint(svc)
		entryByFaceEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByFaceEndpoint)
		entryByFaceEndpoint = zipkin.TraceEndpoint(trace, "EntryByFace")(entryByFaceEndpoint)

		entryByFaceEndpoint = client.WithMeta()(entryByFaceEndpoint)
	}

	var entryByOAuthEndpoint endpoint.Endpoint
	{
		entryByOAuthEndpoint = MakeEntryByOAuthEndpoint(svc)
		entryByOAuthEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByOAuthEndpoint)
		entryByOAuthEndpoint = zipkin.TraceEndpoint(trace, "EntryByOAuth")(entryByOAuthEndpoint)

		entryByOAuthEndpoint = client.WithMeta()(entryByOAuthEndpoint)
	}

	var entryByValidateCodeEndpoint endpoint.Endpoint
	{
		entryByValidateCodeEndpoint = MakeEntryByValidateCodeEndpoint(svc)
		entryByValidateCodeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByValidateCodeEndpoint)
		entryByValidateCodeEndpoint = zipkin.TraceEndpoint(trace, "EntryByValidateCode")(entryByValidateCodeEndpoint)

		entryByValidateCodeEndpoint = client.WithMeta()(entryByValidateCodeEndpoint)
	}

	var entryByQRCodeEndpoint endpoint.Endpoint
	{
		entryByQRCodeEndpoint = MakeEntryByQRCodeEndpoint(svc)
		entryByQRCodeEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByQRCodeEndpoint)
		entryByQRCodeEndpoint = zipkin.TraceEndpoint(trace, "EntryByQRCode")(entryByQRCodeEndpoint)

		entryByQRCodeEndpoint = client.WithMeta()(entryByQRCodeEndpoint)
	}

	var entryByInviteEndpoint endpoint.Endpoint
	{
		entryByInviteEndpoint = MakeEntryByInviteEndpoint(svc)
		entryByInviteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(entryByInviteEndpoint)
		entryByInviteEndpoint = zipkin.TraceEndpoint(trace, "EntryByInvite")(entryByInviteEndpoint)

		entryByInviteEndpoint = client.WithMeta()(entryByInviteEndpoint)
	}

	return Endpoints{
		EntryByAPEndpoint:           entryByAPEndpoint,
		EntryByFaceEndpoint:         entryByFaceEndpoint,
		EntryByOAuthEndpoint:        entryByOAuthEndpoint,
		EntryByQRCodeEndpoint:       entryByQRCodeEndpoint,
		EntryByValidateCodeEndpoint: entryByValidateCodeEndpoint,
		EntryByInviteEndpoint:       entryByInviteEndpoint,
	}

}

func MakeEntryByInviteEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByInvite(ctx, request.(*fs_entry_login.EntryByInviteRequest))
	}
}

func MakeEntryByAPEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByAP(ctx, request.(*fs_entry_login.EntryByAPRequest))
	}
}

func MakeEntryByFaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByFace(ctx, request.(*fs_entry_login.EntryByFaceRequest))
	}
}

func MakeEntryByOAuthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByOAuth(ctx, request.(*fs_entry_login.EntryByOAuthRequest))
	}
}

func MakeEntryByQRCodeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByQRCode(ctx, request.(*fs_entry_login.EntryByQRCodeRequest))
	}
}

func MakeEntryByValidateCodeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.EntryByValidateCode(ctx, request.(*fs_entry_login.EntryByValidateCodeRequest))
	}
}
