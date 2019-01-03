package face

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/base/face/pb"
)

type Endpoints struct {
	UpdateEndpoint     endpoint.Endpoint
	CompareEndpoint    endpoint.Endpoint
	SearchEndpoint     endpoint.Endpoint
	AddFaceEndpoint    endpoint.Endpoint
	RemoveFaceEndpoint endpoint.Endpoint
}

func NewEndpoints(svc Service) Endpoints {

	var updateEndpoint endpoint.Endpoint
	{

	}

	var compareEndpoint endpoint.Endpoint
	{

	}

	var searchEndpoint endpoint.Endpoint
	{

	}

	var addFaceEndpoint endpoint.Endpoint
	{

	}

	var removeFaceEndpoint endpoint.Endpoint
	{

	}

	return Endpoints{
		UpdateEndpoint:     updateEndpoint,
		CompareEndpoint:    compareEndpoint,
		SearchEndpoint:     searchEndpoint,
		AddFaceEndpoint:    addFaceEndpoint,
		RemoveFaceEndpoint: removeFaceEndpoint,
	}

}

func MakeUpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Update(ctx, request.(*fs_base_face.UpdateRequest))
	}
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

func MakeAddFaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.AddFace(ctx, request.(*fs_base_face.AddFaceRequest))
	}
}

func MakeRemoveFaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.RemoveFace(ctx, request.(*fs_base_face.RemoveFaceRequest))
	}
}
