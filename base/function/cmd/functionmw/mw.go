package functionmw

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/openzipkin/zipkin-go"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/cmd/authenticatecli"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/face"
	"zskparker.com/foundation/base/face/cmd/facecli"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/function"
	"zskparker.com/foundation/base/function/cmd/functioncli"
	"zskparker.com/foundation/base/function/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/project"
	"zskparker.com/foundation/base/project/cmd/projectcli"
	"zskparker.com/foundation/base/project/pb"
	"zskparker.com/foundation/base/validate"
	"zskparker.com/foundation/base/validate/cmd/validatecli"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/names"
	"zskparker.com/foundation/pkg/ref"
)

type MWServices struct {
	facecli         face.Service
	authenticatecli authenticate.Service
	validatecli     validate.Service
	projectcli      project.Service
	functioncli     function.Service
}

func NewFunctionMWClient(tracer *zipkin.Tracer) *MWServices {
	return &MWServices{
		authenticatecli: authenticatecli.NewClient(tracer),
		validatecli:     validatecli.NewClient(tracer),
		projectcli:      projectcli.NewClient(tracer),
		facecli:         facecli.NewClient(tracer),
		functioncli:     functioncli.NewClient(tracer),
	}
}

func Middleware(mwcli MWServices) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value("meta").(*fs_base.Metadata)
			mr := ref.GetMetaInfo(request)
			var ps *fs_base.State
			var cf *fs_base_function.Func
			var strategy *fs_base.ProjectStrategy
			var project *fs_base_project.ProjectInfo
			errc := make(chan error, 2)
			go func() {
				pr, _ := mwcli.projectcli.Get(ctx, &fs_base_project.GetRequest{
					ClientId: meta.ClientId,
				})
				if !pr.State.Ok {
					ps = pr.State
					errc <- errno.ERROR
				}
				project = pr.Info
				strategy = pr.Strategy
				errc <- nil
			}()

			go func() {
				fr, _ := mwcli.functioncli.Get(ctx, &fs_base_function.GetRequest{
					Api: meta.Api,
				})
				if !fr.State.Ok {
					ps = fr.State
					errc <- errno.ERROR
				}
				cf = fr.Func
				errc <- nil
			}()

			if err := <-errc; err != nil {
				if ps != nil && ps == errno.ErrFunctionInvalid { //功能未找到
					cf = &fs_base_function.Func{
						Fcv:   names.F_FCV_AUTH,
						Level: 1,
					}
				} else {
					return errno.ErrResponse(errno.ErrSystem)
				}
			}

			metaCheck := func(face bool) bool {
				if face {
					if len(mr.Face) == 0 {
						ps = errno.ErrFaceValidate
						return false
					}
				} else {
					if len(mr.Id) == 0 && len(mr.Validate) == 0 {
						ps = errno.ErrMetaValidate
						return false
					}
				}
				return true
			}

			validateCheck := func() {
				resp, _ := mwcli.validatecli.Verification(context.Background(), &fs_base_validate.VerificationRequest{
					VerId:          mr.Id,
					Code:           mr.Validate,
					Func:           cf.Func,
					OnVerification: strategy.Events.OnVerification,
					Metadata:       meta,
				})
				ps = resp.State
			}

			authCheck := func() {
				if ctx.Value("token") == nil {
					ps = errno.ErrRequest
					return
				}
				resp, _ := mwcli.authenticatecli.Check(context.Background(), &fs_base_authenticate.CheckRequest{
					Metadata: meta,
				})
				ps = resp.State
				meta.UserId = resp.UserId
				meta.Level = resp.Level
			}

			faceCheck := func() {
				resp, _ := mwcli.facecli.Search(context.Background(), &fs_base_face.SearchRequest{
					Base64Face: mr.Face,
				})
				ps = resp.State
				meta.UserId = resp.UserId
				meta.Level = resp.Level
			}

			//验证
			if cf.Fcv != 0 && cf.Fcv != names.F_FCV_NONE {
				if cf.Fcv == names.F_FCV_AUTH {
					authCheck()
				} else if cf.Fcv == names.F_FCV_PHONE && metaCheck(false) {
					validateCheck()
				} else if cf.Fcv == names.F_FCV_EMAIL && metaCheck(false) {
					validateCheck()
				} else if cf.Fcv == names.F_FCV_FACE && metaCheck(true) {
					faceCheck()
				} else if cf.Fcv == names.F_FCV_AUTH|names.F_FCV_FACE && metaCheck(true) {
					authCheck()
					if !ps.Ok {
						return errno.ErrResponse(ps)
					}
					faceCheck()
				} else if cf.Fcv == names.F_FCV_AUTH|names.F_FCV_PHONE && metaCheck(false) {
					authCheck()
					if !ps.Ok {
						return errno.ErrResponse(ps)
					}
					validateCheck()
				} else if cf.Fcv == names.F_FCV_AUTH|names.F_FCV_EMAIL && metaCheck(false) {
					authCheck()
					if !ps.Ok {
						return errno.ErrResponse(ps)
					}
					validateCheck()
				} else {
					return errno.ErrResponse(errno.ErrFunction)
				}
				if !ps.Ok {
					return errno.ErrResponse(ps)
				}
			}

			//check level
			if meta.Level < cf.Level {
				return errno.ErrResponse(errno.ErrRequest)
			}

			return next(ctx, request)
		}
	}
}
