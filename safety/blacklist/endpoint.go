package blacklist

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"zskparker.com/foundation/safety/blacklist/pb"
)

type Endpoints struct {
	CheckAccountEndpoint endpoint.Endpoint
	CheckMetaEndpoint    endpoint.Endpoint
	AddEndpoint          endpoint.Endpoint
}

func NewEndpoints() Endpoints {

	var checkAccountEndpoint endpoint.Endpoint
	{

	}

	var checkMetaEndpoint endpoint.Endpoint
	{

	}

	var addEndpoint endpoint.Endpoint
	{

	}

	return Endpoints{
		CheckAccountEndpoint: checkAccountEndpoint,
		CheckMetaEndpoint:    checkMetaEndpoint,
		AddEndpoint:          addEndpoint,
	}

}

func MakeCheckAccountEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.CheckAccount(ctx, request.(*fs_safety_blacklist.CheckAccountRequest))
	}
}
