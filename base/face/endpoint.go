package face

import (
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/sony/gobreaker"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
)

type Endpoints struct {
	CompareEndpoint    endpoint.Endpoint
	SearchEndpoint     endpoint.Endpoint
	UpsertEndpoint     endpoint.Endpoint
	RemoveFaceEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service, trace *stdzipkin.Tracer, logger log.Logger) Endpoints {

	var compareEndpoint endpoint.Endpoint
	{
		compareEndpoint = MakeCompareEndpoint(svc)
		compareEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(compareEndpoint)
		compareEndpoint = zipkin.TraceEndpoint(trace, "Compare")(compareEndpoint)
	}

	var searchEndpoint endpoint.Endpoint
	{
		searchEndpoint = MakeSearchEndpoint(svc)
		searchEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(searchEndpoint)
		searchEndpoint = zipkin.TraceEndpoint(trace, "Search")(searchEndpoint)
	}

	var upsertEndpoint endpoint.Endpoint
	{
		upsertEndpoint = MakeUpsertEndpoint(svc)
		upsertEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(upsertEndpoint)
		upsertEndpoint = zipkin.TraceEndpoint(trace, "Upsert")(upsertEndpoint)
	}

	var removeFaceEndpoint endpoint.Endpoint
	{
		removeFaceEndpoint = MakeRemoveFaceEndpoint(svc)
		removeFaceEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(removeFaceEndpoint)
		removeFaceEndpoint = zipkin.TraceEndpoint(trace, "Remove")(removeFaceEndpoint)
	}

	return Endpoints{
		CompareEndpoint:    compareEndpoint,
		SearchEndpoint:     searchEndpoint,
		UpsertEndpoint:     upsertEndpoint,
		RemoveFaceEndpoint: removeFaceEndpoint,
	}

}

func (g Endpoints) Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error) {
	resp, err := g.CompareEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error) {
	resp, err := g.SearchEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base_face.SearchResponse), nil
}

func (g Endpoints) Upsert(ctx context.Context, in *fs_base_face.UpsertRequest) (*fs_base.Response, error) {
	resp, err := g.UpsertEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func (g Endpoints) RemoveFace(ctx context.Context, in *fs_base_face.RemoveFaceRequest) (*fs_base.Response, error) {
	resp, err := g.RemoveFaceEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*fs_base.Response), nil
}

func MakeCompareEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Compare(ctx, request.(*fs_base_face.CompareRequest))
	}
}

func MakeSearchEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Search(ctx, request.(*fs_base_face.SearchRequest))
	}
}

func MakeUpsertEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Upsert(ctx, request.(*fs_base_face.UpsertRequest))
	}
}

func MakeRemoveFaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.RemoveFace(ctx, request.(*fs_base_face.RemoveFaceRequest))
	}
}
