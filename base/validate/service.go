package validate

import (
	"context"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strings"
	"time"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/base/veds"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/sync"
	"zskparker.com/foundation/pkg/tool/encrypt"
	"zskparker.com/foundation/pkg/tool/veds"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base_validate.VerificationResponse, error)

	Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error)
}

type validateService struct {
	session    *mgo.Session
	messagecli messagecli.Channel
	state      state.Service
	redisync   *fs_redisync.Redisync
	vedscli    veds.Service
}

func (svc *validateService) GetRepo() repository {
	return &validateRepository{session: svc.session.Clone()}
}

func (svc *validateService) Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	if len(in.Metadata.Ip) == 0 {
		return &fs_base_validate.CreateResponse{
			State: errno.ErrRequest,
		}, nil
	}

	resp := func(s *fs_base.State) (*fs_base_validate.CreateResponse, error) {
		return &fs_base_validate.CreateResponse{State: s}, nil
	}

	//验证凭证(通过操作有时间等限制)
	voucher := in.Metadata.Ip + ";" + in.Func + ";" + in.Metadata.UserAgent
	//有用户ID设置为凭证
	if len(in.Metadata.UserId) > 0 {
		voucher = in.Metadata.UserId + ";" + in.Func + ";" + in.Metadata.UserAgent
	}
	voucher = fs_tools_encrypt.SHA256(voucher)

	//锁一会(默认60秒)
	if s := svc.redisync.Lock(in.Func, voucher, in.OnVerification.VoucherDuration); s != nil {
		return resp(s)
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var code string
	if in.OnVerification.CombinationMode == 1 {
		code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	} else {
		code = strings.ToUpper(utils.GetRandomString())[0:8]
	}

	vl := &verification{
		VerId:    bson.NewObjectId(),
		Func:     in.Func,
		Voucher:  voucher,
		To:       fs_tools_encrypt.SHA256(in.To),
		CreateAt: time.Now().UnixNano(),
		Code:     fs_tools_encrypt.SHA1_256_512(code),
		State:    fs_constants.STATE_WAIT_VALIDATE,
	}

	err := repo.Create(vl)
	if err != nil {
		return resp(errno.ErrSystem)
	}

	//send code
	switch in.Mode {
	case 1: //phone
		svc.messagecli.SendSMS(&fs_base.DirectMessage{
			To:      in.To,
			Content: fmt.Sprintf(osenv.GetValidateTemplate(), code, in.OnVerification.EffectiveTime),
		})
		break
	case 2: //email
		svc.messagecli.SendEmail(&fs_base.DirectMessage{
			To:      in.To,
			Content: fmt.Sprintf(osenv.GetValidateTemplate(), code, in.OnVerification.EffectiveTime),
		})
		break
	case 3: //face 这里不做操作
		break
	default:
		return resp(errno.ErrSupport)
	}

	v := fs_service_veds.Encrypt(svc.vedscli, vl.VerId.Hex())
	if !v.State.Ok {
		return resp(v.State)
	}

	return &fs_base_validate.CreateResponse{
		VerId: v.Values[0],
		State: errno.Ok,
	}, nil

}

func (svc *validateService) Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base_validate.VerificationResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	if len(in.Code) == 0 && len(in.Func) == 0 {
		return &fs_base_validate.VerificationResponse{State: errno.ErrRequest}, nil
	}

	vl, err := repo.Get(in.VerId)
	if err != nil {
		return &fs_base_validate.VerificationResponse{State: errno.ErrRequest}, nil
	}

	b := time.Now().UnixNano()-vl.CreateAt <= in.OnVerification.EffectiveTime*60*1e9

	//检查时间和操作
	if b {
		voucher := in.Metadata.Ip + ";" + in.Func + ";" + in.Metadata.UserAgent
		if len(in.Metadata.UserId) > 16 {
			voucher = in.Metadata.UserId + ";" + in.Func + ";" + in.Metadata.UserAgent
		}
		md := fs_tools_encrypt.SHA256(voucher)
		if md != vl.Voucher {
			return &fs_base_validate.VerificationResponse{State: errno.ErrRequest}, nil
		}
		//需要在等待验证状态
		if vl.State == fs_constants.STATE_WAIT_VALIDATE {
			code := strings.ToLower(in.Code)
			if a := fs_tools_encrypt.SHA1_256_512(code); a == vl.Code {
				return &fs_base_validate.VerificationResponse{State: errno.Ok, To: vl.To}, nil
			}
			return &fs_base_validate.VerificationResponse{State: errno.ErrValidateCode}, nil
		}
	}

	return &fs_base_validate.VerificationResponse{State: errno.ErrExpired}, nil
}

func NewService(session *mgo.Session, messagecli messagecli.Channel, state state.Service, redisync *fs_redisync.Redisync, vedscli veds.Service) Service {
	var svc Service
	{
		svc = &validateService{session: session, messagecli: messagecli, state: state, redisync: redisync, vedscli: vedscli}
	}
	return svc
}
