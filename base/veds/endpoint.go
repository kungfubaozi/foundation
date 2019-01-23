package veds

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/veds/pb"
)

type Endpoints struct {
	EncryptEndpoint endpoint.Endpoint
	DecryptEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var encryptEndpoint endpoint.Endpoint
	{
		encryptEndpoint = MakeEncryptEndpoint(svc)
		encryptEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(encryptEndpoint)
		encryptEndpoint = zipkin.TraceEndpoint(trace, "Encrypt")(encryptEndpoint)
	}

	var decryptEndpoint endpoint.Endpoint
	{
		decryptEndpoint = MakeDecryptEndpoint(svc)
		decryptEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(decryptEndpoint)
		decryptEndpoint = zipkin.TraceEndpoint(trace, "Decrypt")(decryptEndpoint)
	}

	return Endpoints{
		EncryptEndpoint: encryptEndpoint,
		DecryptEndpoint: decryptEndpoint,
	}

}

func (g Endpoints) Encrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	resp, err := g.EncryptEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_veds.CryptResponse), nil
}

func (g Endpoints) Decrypt(ctx context.Context, in *fs_base_veds.CryptRequest) (*fs_base_veds.CryptResponse, error) {
	resp, err := g.DecryptEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_veds.CryptResponse), nil
}

func MakeEncryptEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Encrypt(ctx, request.(*fs_base_veds.CryptRequest))
	}
}

func MakeDecryptEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Decrypt(ctx, request.(*fs_base_veds.CryptRequest))
	}
}
