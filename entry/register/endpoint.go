package register

import "github.com/go-kit/kit/endpoint"

type Endpoints struct {
	FromAPEndpoint    endpoint.Endpoint
	FromOAuthEndpoint endpoint.Endpoint
}
