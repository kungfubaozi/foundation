package refresh

import (
	"context"
	"github.com/go-kit/kit/log"
	"zskparker.com/foundation/base/authenticate"
	"zskparker.com/foundation/base/authenticate/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/refresh/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/pkg/tags"
	"zskparker.com/foundation/pkg/transport"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Auth(ctx context.Context, in *fs_base_refresh.AuthRequest) (*fs_base_refresh.AuthResponse, error)
}

type refreshService struct {
	authenticatecli authenticate.Service
	reportercli     reportercli.Channel
	logger          log.Logger
	statecli        state.Service
	usercli         user.Service
	redisync        *fs_redisync.Redisync
}

func (svc *refreshService) Auth(ctx context.Context, in *fs_base_refresh.AuthRequest) (*fs_base_refresh.AuthResponse, error) {
	resp := func(state *fs_base.State) (*fs_base_refresh.AuthResponse, error) {
		return &fs_base_refresh.AuthResponse{State: state}, nil
	}
	if len(in.RefreshToken) < 32 {
		svc.logger.Log("step")
		return resp(errno.ErrRequest)
	}
	refreshTokenClaims, err := authenticate.DecodeToken(in.RefreshToken)
	svc.logger.Log(refreshTokenClaims.ExpiresAt)
	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrToken)
	}
	if refreshTokenClaims.Valid() != nil {
		svc.logger.Log("step")
		return resp(errno.ErrTokenExpired)
	}
	token := refreshTokenClaims.Token
	if token.Access {
		svc.logger.Log("step")
		return resp(errno.ErrToken)
	}

	//锁定30s
	if s := svc.redisync.Lock(fs_function_tags.GetAuthenticateRefreshTag(), token.Relation, 30); s.Ok {
		return resp(s)
	}

	meta := fs_metadata_transport.ContextToMeta(ctx)

	authr, err := svc.authenticatecli.Get(context.Background(), &fs_base_authenticate.GetRequest{
		UserId:   meta.UserId,
		ClientId: meta.ClientId,
		Relation: token.Relation,
	})
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if !authr.State.Ok {
		return resp(authr.State)
	}

	auth := authr.Auth

	node := utils.NodeGenerate()
	auth.AccessTokenUUID = node.Generate().Base64()
	var accessToken string

	//!!!验证操作已经在拦截器中设置完成（状态检测，黑名单，用户依存等）

	//可能是跳转的token
	if auth.ClientId != meta.ClientId {
		svc.logger.Log("step")
		//需要通过web端才可跨项目访问
		if auth.Platform == fs_constants.PLATFORM_WEB && meta.Platform == fs_constants.PLATFORM_WEB {
			accessToken, err = authenticate.EncodeAccessToken(&fs_base_authenticate.Authorize{
				ClientId:        meta.ClientId,
				UserId:          auth.UserId,
				AccessTokenUUID: auth.AccessTokenUUID,
				Relation:        auth.Relation,
			})
			if err != nil {
				svc.logger.Log("err")
				return resp(errno.ErrSystem)
			}

			//由于在缓存里取出的是遵循 userId,clientId,relationId
			//当出现使用refreshToken刷新别的web client时，会新建一个以当前用户的基本信息+client生成token
			//公用一个refreshToken的Relation，
			auth.ClientId = meta.ClientId
		} else { //除了web端其他的客户端不支持token公用
			svc.logger.Log("step")
			return resp(errno.ErrSupport)
		}
	} else {
		svc.logger.Log("step")
		accessToken, err = authenticate.EncodeAccessToken(auth)
		if err != nil {
			svc.logger.Log("step")
			return resp(errno.ErrSystem)
		}
	}

	//覆盖原有的token

	ar, err := svc.authenticatecli.ReplaceAuth(context.Background(), &fs_base_authenticate.ReplaceAuthRequest{
		Auth: auth,
	})
	if err != nil {
		svc.logger.Log("step")
		return resp(errno.ErrSystem)
	}
	if !ar.State.Ok {
		return resp(ar.State)
	}

	return &fs_base_refresh.AuthResponse{
		State:       errno.Ok,
		AccessToken: accessToken,
	}, nil
}

func NewService(authenticatecli authenticate.Service, reportercli reportercli.Channel, logger log.Logger,
	statecli state.Service, redisync *fs_redisync.Redisync, usercli user.Service) Service {
	var svc Service
	{
		svc = &refreshService{authenticatecli: authenticatecli, reportercli: reportercli, logger: logger,
			statecli: statecli, redisync: redisync, usercli: usercli}
	}
	return svc
}
