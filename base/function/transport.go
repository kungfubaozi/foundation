package function

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"zskparker.com/foundation/base/function/pb"
)

type GRPCServer struct {
	add    grpctransport.Handler
	remove grpctransport.Handler
	get    grpctransport.Handler
	update grpctransport.Handler
}

func MakeGRPCServer() fs_base_function.FunctionServer {

}

func MakeGRPCClient() Service {

}
