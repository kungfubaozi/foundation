package functionmw

import (
	"context"
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
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/ref"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
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
	return middleware(logger, mwcli, function, false)
}

func WithMeta(logger log.Logger, mwcli *MWServices) endpoint.Middleware {
	return middleware(logger, mwcli, "", false)
}

//忽略项目等级
func WithIgnoreProjectLevel(logger log.Logger, mwcli *MWServices, function string) endpoint.Middleware {
	return middleware(logger, mwcli, function, true)
}

func middleware(logger log.Logger, mwcli *MWServices, function string, ignoreProjectLevel bool) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			meta := ctx.Value(fs_metadata_transport.MetadataTransportKey).(*fs_base.Metadata)

			mr := ref.GetMetaInfo(request)
			ps := errno.Ok
			var cf *fs_base_function.Func
			var strategy *fs_base.ProjectStrategy
			var p *fs_base_project.ProjectInfo
			var wg sync.WaitGroup

			//必须传入对应的项目session
			if len(meta.Session) < 32 {
				return errno.ErrRequest, errno.ERROR
			}

			errc := func(s *fs_base.State) {
				if !ps.Ok {
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
				p = pr.Info
				meta.Platform = p.Platform.Platform
				meta.ProjectId = pr.Strategy.ProjectId
				strategy = pr.Strategy
				wg.Done()
			}()

			wg.Add(1)
			go func() {
				fr, _ := mwcli.functioncli.Get(context.Background(), &fs_base_function.GetRequest{
					Tag:  utils.Md5(meta.Api + meta.Session),
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

			wg.Wait()

			if !ps.Ok {
				logger.Log("middleware", "state", "!ok", ps)
				return ps, errno.ERROR
			}

			if strategy == nil || p == nil {
				logger.Log("middleware", "check", "strategy|project", "invalid")
				return errno.ErrClient, errno.ERROR
			}

			//项目权限
			if !ignoreProjectLevel && meta.Level >= p.Level {
				return errno.ErrProjectPermission, errno.ERROR
			}

			if meta.Session != p.Session {
				return errno.ErrRequestPermission, errno.ERROR
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
				resp, err := mwcli.validatecli.Verification(context.Background(), &fs_base_validate.VerificationRequest{
					VerId:          mr.Id,
					Code:           mr.Validate,
					Func:           cf.Func,
					OnVerification: strategy.Events.OnVerification,
					Metadata:       meta,
				})
				if err != nil {
					ps = errno.ErrSystem
					return
				}
				if resp == nil {
					ps = errno.ErrSystem
					return
				}
				ps = resp.State
				ctx = context.WithValue(ctx, "validate_to", resp.To)
			}

			authCheck := func() {
				if len(meta.Token) <= 32 {
					ps = errno.ErrToken
					return
				}
				resp, err := mwcli.authenticatecli.Check(context.Background(), &fs_base_authenticate.CheckRequest{
					Metadata: meta,
				})
				if err != nil {
					ps = errno.ErrSystem
					return
				}
				ps = resp.State
				meta.UserId = resp.UserId
				meta.Level = resp.Level
			}

			faceCheck := func() {
				resp, err := mwcli.facecli.Search(context.Background(), &fs_base_face.SearchRequest{
					Base64Face: mr.Face,
				})
				if err != nil {
					ps = errno.ErrSystem
					return
				}
				ps = resp.State
				meta.UserId = resp.UserId
				meta.Level = resp.Level
			}

			//验证
			if cf.Fcv != 0 && cf.Fcv != fs_constants.FCV_NONE {
				logger.Log("middleware", "fcv", "step", "check")
				if cf.Fcv == fs_constants.FCV_AUTH {
					logger.Log("middleware", "fcv", "step", "1-0-1")
					authCheck()
				} else if cf.Fcv == fs_constants.FCV_VALIDATE_CODE {
					logger.Log("middleware", "fcv", "step", "2-0-1")
					if metaCheck(false) {
						logger.Log("middleware", "fcv", "step", "2-1-1")
						validateCheck()
					}
				} else if cf.Fcv == fs_constants.FCV_FACE {
					logger.Log("middleware", "fcv", "step", "3-0-1")
					if metaCheck(false) {
						logger.Log("middleware", "fcv", "step", "3-1-1")
						faceCheck()
					}
				} else if cf.Fcv == fs_constants.FCV_AUTH|fs_constants.FCV_FACE {
					logger.Log("middleware", "fcv", "step", "4-0-1")
					if metaCheck(true) {
						logger.Log("middleware", "fcv", "step", "4-1-1")
						authCheck()
						if !ps.Ok {
							logger.Log("middleware", "fcv", "step", "4-1-2", "err", ps)
							return ps, errno.ERROR
						}
						logger.Log("middleware", "fcv", "step", "4-2-1")
						faceCheck()
					}
				} else if cf.Fcv == fs_constants.FCV_AUTH|fs_constants.FCV_VALIDATE_CODE {
					logger.Log("middleware", "fcv", "step", "5-0-1")
					if metaCheck(false) {
						logger.Log("middleware", "fcv", "step", "5-1-1")
						authCheck()
						if !ps.Ok {
							logger.Log("middleware", "fcv", "step", "5-1-2", "err", ps)
							return ps, errno.ERROR
						}
						logger.Log("middleware", "fcv", "step", "5-2-1")
						validateCheck()
					}
				} else {
					logger.Log("middleware", "fcv", "step", "failed")
					return errno.ErrFunction, errno.ERROR
				}
				if ps == nil {
					return nil, errno.ERROR
				}
				if !ps.Ok {
					logger.Log("middleware", "fcv", "step", "!ok", ps)
					return ps, errno.ERROR
				}
			}

			logger.Log("middleware", "function", "check", "ok")

			ctx = context.WithValue(ctx, fs_metadata_transport.StrategyTransportKey, strategy)
			ctx = context.WithValue(ctx, fs_metadata_transport.ProjectTransportKey, p)

			//check level
			if meta.Level >= cf.Level {
				logger.Log("middleware", "fcv", "step", "next")
				return next(ctx, request)
			}

			return errno.ErrRequest, errno.ERROR
		}
	}
}
