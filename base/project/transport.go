package project

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"zskparker.com/foundation/base/project/pb"
)

type GRPCServer struct {
	new grpctransport.Handler
	get grpctransport.Handler
}

func MakeGRPCServer() fs_base_project.ProjectServer {

}

func MakeGRPCClient() Service {

}
