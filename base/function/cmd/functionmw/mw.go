package functionmw

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"sync"
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
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/safety/blacklist"
)

type MWServices struct {
	facecli         face.Service
	authenticatecli authenticate.Service
	validatecli     validate.Service
	projectcli      project.Service
	functioncli     function.Service
	blacklistcli    blacklist.Service
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

func WithExpress(logger log.Logger, mwcli *MWServices, function string) endpoint.Middleware {
	return middleware(logger, mwcli, function)
}

func WithMeta(logger log.Logger, mwcli *MWServices) endpoint.Middleware {
	return middleware(logger, mwcli, "")
}

func middleware(logger log.Logger, mwcli *MWServices, function string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value(fs_metadata_transport.MetadataTransportKey).(*fs_base.Metadata)

			mr := ref.GetMetaInfo(request)
			ps := errno.Ok
			var cf *fs_base_function.Func
			var strategy *fs_base.ProjectStrategy
			var project *fs_base_project.ProjectInfo
			var wg sync.WaitGroup

			errc := func(s *fs_base.State) {
				if ps.Ok {
					ps = s
				}
				wg.Done()
			}

			wg.Add(1)
			go func() {
				pr, _ := mwcli.projectcli.Get(context.Background(), &fs_base_project.GetRequest{
					ClientId: meta.ClientId,
				})
				if pr == nil {
					errc(errno.ErrSystem)
					logger.Log("middleware", "function", "err", "find project nil")
					return
				}
				if !pr.State.Ok {
					errc(pr.State)
					logger.Log("middleware", "function", "state", "project", "value", pr)
					return
				}
				project = pr.Info
				meta.ProjectId = pr.Strategy.ProjectId
				strategy = pr.Strategy
				wg.Done()
			}()

			wg.Add(1)
			go func() {
				fr, _ := mwcli.functioncli.Get(context.Background(), &fs_base_function.GetRequest{
					Api:  meta.Api,
					Func: function,
				})
				if fr == nil {
					errc(errno.ErrSystem)
					logger.Log("middleware", "function", "err", "find function nil")
					return
				}
				if !fr.State.Ok {
					errc(fr.State)
					logger.Log("middleware", "function", "state", "function", "value", fr)
					return
				}
				cf = fr.Func
				wg.Done()
			}()

			//wg.Add(1)
			//go func() {
			//	//blacklist
			//	var userId string
			//	if len(meta.Token) > 0 {
			//		s, err := authenticate.DecodeToken(meta.Token)
			//		if err != nil {
			//			errc(errno.ErrToken)
			//			return
			//		}
			//		userId = s.Token.UserId
			//	}
			//
			//	br, _ := mwcli.blacklistcli.Check(context.Background(), &fs_safety_blacklist.CheckRequest{
			//		Ip:     meta.Ip,
			//		UserId: userId,
			//		Device: meta.Device,
			//	})
			//
			//	if br == nil {
			//		errc(errno.ErrSystem)
			//		return
			//	}
			//	if !br.State.Ok {
			//		errc(br.State)
			//		return
			//	}
			//	wg.Done()
			//}()

			wg.Wait()

			if !ps.Ok {
				if ps == errno.ErrFunctionInvalid { //功能未找到
					logger.Log("middleware", "function", "invalid", meta.Api)
					cf = &fs_base_function.Func{
						Fcv:   names.F_FCV_AUTH,
						Level: 1,
					}
				} else {
					fmt.Println("e1-1")
					return ps, errno.ERROR
				}
			}

			if strategy == nil || project == nil {
				logger.Log("middleware", "check", "strategy|project", "invalid")
				return errno.ErrSystem, errno.ERROR
			}

			metaCheck := func(face bool) bool {
				if face {
					if len(mr.Face) == 0 {
						ps = errno.ErrFaceValidate
						return false
					}
				} else {
					if len(mr.Id) == 0 || len(mr.Validate) == 0 {
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
				if resp == nil {
					ps = errno.ErrSystem
					return
				}
				ps = resp.State
				ctx = context.WithValue(ctx, "validate_to", resp.To)
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
				fmt.Println("e2-0")
				if cf.Fcv == names.F_FCV_AUTH {
					authCheck()
				} else if cf.Fcv == names.F_FCV_VALIDATE_CODE {
					if metaCheck(false) {
						validateCheck()
					}
				} else if cf.Fcv == names.F_FCV_FACE {
					if metaCheck(false) {
						faceCheck()
					}
				} else if cf.Fcv == names.F_FCV_AUTH|names.F_FCV_FACE {
					if metaCheck(true) {
						authCheck()
						if !ps.Ok {
							return errno.ErrResponse(ps)
						}
						faceCheck()
					}
				} else if cf.Fcv == names.F_FCV_AUTH|names.F_FCV_VALIDATE_CODE {
					if metaCheck(false) {
						authCheck()
						if !ps.Ok {
							return errno.ErrResponse(ps)
						}
						validateCheck()
					}
				} else {
					fmt.Println("e2")
					return errno.ErrResponse(errno.ErrFunction)
				}
				if !ps.Ok {
					fmt.Println("e3")
					return errno.ErrResponse(ps)
				}
			}

			logger.Log("middleware", "function", "check", "ok")

			ctx = context.WithValue(ctx, fs_metadata_transport.StrategyTransportKey, strategy)
			ctx = context.WithValue(ctx, fs_metadata_transport.ProjectTransportKey, project)

			//check level
			if meta.Level >= cf.Level {
				fmt.Println("next")
				return next(ctx, request)
			}

			return errno.ErrRequest, errno.ERROR
		}
	}
}
