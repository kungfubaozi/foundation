package validatemw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/openzipkin/zipkin-go"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/cmd/validatecli"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/ref"
)

type Client struct {
	FunctionCli function.Service
	ValidateCli validate.Service
	FaceCli     face.Service
}

func NewClient(tracer *zipkin.Tracer) Client {
	return Client{
		ValidateCli: validatecli.NewClient(osenv.GetConsulAddr(), tracer),
	}
}

func Middleware(client Client) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			mr := ref.GetMetaInfo(request)
			fr, _ := client.FunctionCli.Get(context.Background(), &fs_base_function.GetRequest{
				Api: meta.Api,
			})
			if !fr.State.Ok {
				return errno.ErrResponse(fr.State)
			}
			ctx = context.WithValue(ctx, "func", fr.Func)
			if !fr.Func.Verification {
				return next(ctx, request)
			}
			//face validate
			if fr.Func.Fcv == names.F_VALIDATE_FACE {
				if len(mr.Face) == 0 {
					return errno.ErrResponse(errno.ErrMetaValidate)
				}
				fcr, _ := client.FaceCli.Compare(context.Background(), &fs_base_face.CompareRequest{
					UserId:     meta.UserId,
					Base64Face: mr.Face,
				})
				if !fcr.State.Ok {
					return errno.ErrResponse(fcr.State)
				}
			} else {
				if len(mr.Id) == 0 && len(mr.Validate) == 0 {
					return errno.ErrResponse(errno.ErrMetaValidate)
				}
				vcr, _ := client.ValidateCli.Verification(context.Background(), &fs_base_validate.VerificationRequest{
					Metadata: meta,
					VerId:    mr.Id,
					Code:     mr.Validate,
					Func:     fr.Func.Func,
				})
				if !vcr.State.Ok {
					return errno.ErrResponse(vcr.State)
				}
			}
			return next(ctx, request)
		}
	}
}
