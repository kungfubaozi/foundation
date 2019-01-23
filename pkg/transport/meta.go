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
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/functions"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/tool/encrypt"
)

var (
	MetadataTransportKey = "meta"
	StrategyTransportKey = "strategy"
	ProjectTransportKey  = "project"
	ValidateTransportKey = "validate_account"
)

func CheckValidateAccount(ctx context.Context, account string) *fs_base.State {
	if ctx.Value(ValidateTransportKey) == nil {
		return errno.ErrRequest
	}
	to := ContextToValidateAccount(ctx)
	if to != fs_tools_encrypt.SHA256(account) {
		return errno.ErrRequest
	}
	return errno.Ok
}

func MetaToReporter(reportercli reportercli.Channel, ctx context.Context, who string, status int64) {
	meta := ContextToMeta(ctx)
	reportercli.Write(meta.FuncTag, who, fmt.Sprintf("%s;%s;%s;%s;%s", meta.Ip, meta.ProjectId, meta.ClientId, meta.UserAgent, meta.Device), status)
}

func MetaToReporterByMetadata(reportercli reportercli.Channel, meta *fs_base.Metadata, who, tag string, status int64) {
	reportercli.Write(tag, who, fmt.Sprintf("%s;%s;%s;%s;%s", meta.Ip, meta.ProjectId, meta.ClientId, meta.UserAgent, meta.Device), status)
}

func MetaToReporterByTag(reportercli reportercli.Channel, ctx context.Context, who string, tag string, status int64) {
	meta := ContextToMeta(ctx)
	reportercli.Write(tag, who, fmt.Sprintf("%s;%s;%s;%s;%s", meta.Ip, meta.ProjectId, meta.ClientId, meta.UserAgent, meta.Device), status)
}

func ContextToValidateAccount(ctx context.Context) string {
	return ctx.Value(ValidateTransportKey).(string)
}

func ContextToStrategy(ctx context.Context) *fs_base.Strategy {
	return ctx.Value(StrategyTransportKey).(*fs_base.Strategy)
}

func ContextToMeta(ctx context.Context) *fs_base.Metadata {
	return ctx.Value(MetadataTransportKey).(*fs_base.Metadata)
}

func ContextToProject(ctx context.Context) *fs_base_project.ProjectInfo {
	return ctx.Value(ProjectTransportKey).(*fs_base_project.ProjectInfo)
}

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
		if e != nil || i == 0 {
			i = fs_constants.LEVEL_TOURIST
		}
		meta.Level = i
		meta.Session = header[8]
		meta.InitSession = header[9]

		return context.WithValue(ctx, MetadataTransportKey, meta)
	}
}

func HTTPToContext() http.RequestFunc {
	//检查session
	session := osenv.GetInitializeProjectSession()
	return func(ctx context.Context, request *stdhttp.Request) context.Context {
		meta := &fs_base.Metadata{}
		meta.Device = request.Header.Get("X-User-Device")
		meta.ClientId = request.Header.Get("X-Client-Id")

		meta.Ip = request.Header.Get("X-Real-IP")
		meta.UserAgent = request.Header.Get("User-Agent")
		meta.Api = uri(request.RequestURI)
		forward := request.Header.Get("X-Forward-URI")
		f := fs_functions.GetInterceptFunc()
		meta.InitSession = session
		meta.Session = session
		if len(forward) > 2 && meta.Api == fmt.Sprintf("%s%s", f.Prefix, f.Infix) { //只允许拦截器设置URI和session
			meta.Api = forward
			meta.Session = request.Header.Get("X-Server-Session")
		}
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
		m, ok := ctx.Value(MetadataTransportKey).(*fs_base.Metadata)
		if ok {
			// capital "Key" is illegal in HTTP/2.
			(*md)["authorization"] = []string{
				m.ClientId,
				m.Ip,
				m.UserAgent,
				m.Api,
				m.Token,
				m.Device,
				m.UserId,
				strconv.FormatInt(m.Level, 10),
				m.Session,
				m.InitSession,
			}
		}
		return ctx
	}
}

func GetResponseState(err error, resp interface{}) *fs_base.State {
	if err == errno.ERROR {
		if resp == nil {
			return errno.ErrRequest
		}
		return resp.(*fs_base.State)
	}
	return errno.ErrSystem
}
