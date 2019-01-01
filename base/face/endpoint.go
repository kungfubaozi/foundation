package face

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	UpdateEndpoint     endpoint.Endpoint
	CompareEndpoint    endpoint.Endpoint
	SearchEndpoint     endpoint.Endpoint
	AddFaceEndpoint    endpoint.Endpoint
	RemoveFaceEndpoint endpoint.Endpoint
}

func NewEndpoints() Endpoints {

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

	}

}

func MakeUpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Update(ctx,request.(*))
	}
}
