package format

import (
	"context"
)

func GrpcMessage(ctx context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

//
//func GRPCMetadata() grpc.ClientRequestFunc {
//	return func(i context.Context, md *metadata.MD) context.Context {
//		fmt.Println(fmt.Sprintf("%s-%d-%s", "client 2", time.Now().UnixNano(), i.Value("meta")))
//		return i
//	}
//}
//
//func GRPCServerMetadata() grpc.ServerRequestFunc {
//	return func(i context.Context, mds metadata.MD) context.Context {
//		fmt.Println(mds.Len())
//
//		md, ok := metadata.FromIncomingContext(i)
//		if ok {
//			for _, v := range md {
//				fmt.Println("server", v)
//			}
//		}
//
//		fmt.Println(fmt.Sprintf("%s-%d-%s", "server 1", time.Now().UnixNano(), i.Value("meta")))
//		return i
//	}
//}

//client-1546873990712908900-
//client-1546873990713109800-
//server-1546873990714069900-%!s
