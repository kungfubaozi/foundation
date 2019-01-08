package fs_metadata_transport

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc/metadata"
	stdhttp "net/http"
	"strconv"
	"strings"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
)

var (
	MetadataTransportKey = "meta"
)

func GRPCToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, mds metadata.MD) context.Context {
		header, ok := mds["authorization"]
		if !ok {
			ctx.Err()
			return ctx
		}
		meta := &fs_base.Metadata{}
		meta.ClientId = header[0]
		meta.Ip = header[1]
		meta.UserAgent = header[2]
		meta.Api = header[3]
		meta.Token = header[4]
		meta.Device = header[5]
		meta.UserId = header[6]
		i, e := strconv.ParseInt(header[7], 10, 64)
		if e != nil {
			i = 0
		}
		meta.Level = i
		return context.WithValue(ctx, MetadataTransportKey, meta)
	}
}

func HTTPToContext() http.RequestFunc {
	return func(ctx context.Context, request *stdhttp.Request) context.Context {
		meta := &fs_base.Metadata{}
		meta.Device = request.Header.Get("X-User-Device")
		meta.ClientId = request.Header.Get("X-Client-Id")
		meta.Ip = request.Header.Get("X-Real-IP")
		meta.UserAgent = request.Header.Get("User-Agent")
		meta.Api = uri(request.RequestURI)
		meta.Token = request.Header.Get("Authorization")
		return context.WithValue(ctx, MetadataTransportKey, meta)
	}
}

func uri(uri string) string {
	if i := strings.Index(uri, "?"); i != -1 {
		return uri[:i]
	}
	return uri
}

func ContextToGRPC() grpc.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		metadata, ok := ctx.Value(MetadataTransportKey).(*fs_base.Metadata)
		if ok {
			// capital "Key" is illegal in HTTP/2.
			(*md)["authorization"] = []string{
				metadata.ClientId,
				metadata.Ip,
				metadata.UserAgent,
				metadata.Api,
				metadata.Token,
				metadata.Device,
				metadata.UserId,
				strconv.FormatInt(metadata.Level, 10),
			}
		}
		return ctx
	}
}

func GetResponseState(err error, resp interface{}) *fs_base.State {
	fmt.Println("response state")
	if err == errno.ERROR {
		if resp == nil {
			return errno.ErrSystem
		}
		return resp.(*fs_base.State)
	}
	return errno.ErrSystem
}
