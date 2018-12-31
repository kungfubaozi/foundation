package validate

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/state/pb"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/states"
	"zskparker.com/foundation/pkg/utils"
)

type Service interface {
	Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base.Response, error)

	Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error)
}

type validateService struct {
	session *mgo.Session
	channel messagecli.MessageChannel
	state   state.Service
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

	//验证凭证(通过操作有时间等限制)
	voucher := in.Metadata.Ip + ";" + strconv.FormatInt(in.Do, 10)

	//有用户ID设置为凭证
	if len(in.Metadata.UserId) > 0 {
		voucher = in.Metadata.UserId + ";" + strconv.FormatInt(in.Do, 10)
	}

	//查找最后一次同个操作的时间
	vl, err := repo.FindLast(voucher)
	if err == mgo.ErrNotFound {
		vl = &verification{
			CreateAt: time.Now().UnixNano() - in.OnVerification.ExpiredTime*10e9,
			Do:       in.Do,
			Voucher:  voucher,
		}
	}
	if err != nil {
		return &fs_base_validate.CreateResponse{
			State: errno.ErrSystem,
		}, nil
	}

	//限制时间
	if time.Now().UnixNano()-vl.CreateAt <= in.OnVerification.ExpiredTime*10e9 {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		var code string
		if in.OnVerification.CombinationMode == 1 {
			code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
		} else {
			code = strings.ToLower(utils.GetRandomString())
		}
		b, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
		if err != nil {
			return &fs_base_validate.CreateResponse{
				State: errno.ErrSystem,
			}, nil
		}

		vl = &verification{
			VerId:    utils.RandomMD5(voucher),
			CreateAt: time.Now().UnixNano(),
			Code:     string(b),
		}

		//send code
		switch in.Mode {
		case 1: //phone
			svc.channel.SendSMS(&fs_base.DirectMessage{
				To:      in.To,
				Content: fmt.Sprintf(osenv.GetValidateTemplate(), code, in.OnVerification.EffectiveTime),
			})
			break
		case 2: //email
			svc.channel.SendEmail(&fs_base.DirectMessage{
				To:      in.To,
				Content: fmt.Sprintf(osenv.GetValidateTemplate(), code, in.OnVerification.EffectiveTime),
			})
			break
		case 3: //face 这里不做操作
			break
		default:
			return &fs_base_validate.CreateResponse{
				State: errno.ErrSupport,
			}, nil
		}

		//添加到状态管理里
		resp, err := svc.state.Upsert(context.Background(), &fs_base_state.UpsertRequest{
			Key:    vl.Voucher + "-" + vl.VerId,
			Status: states.F_STATE_WAITING,
		})

		if err != nil {
			return &fs_base_validate.CreateResponse{
				State: errno.ErrSystem,
			}, nil
		}

		if !resp.State.Ok {
			return &fs_base_validate.CreateResponse{
				State: resp.State,
			}, nil
		}

		return &fs_base_validate.CreateResponse{
			VerId: vl.VerId,
			State: errno.Ok,
		}, nil
	}

	return &fs_base_validate.CreateResponse{
		State: errno.ErrBusy,
	}, nil
}

func (svc *validateService) Verification(ctx context.Context, in *fs_base_validate.VerificationRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	if len(in.Code) == 0 && in.Do < 10100 {
		return errno.ErrResponse(errno.ErrRequest)
	}

	vl, err := repo.Get(in.VerId)
	if err != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}

	//检查时间和操作
	if time.Now().UnixNano()-vl.CreateAt < in.OnVerification.EffectiveTime*60*1e9 && in.Do == vl.Do {

		//验证状态
		resp, err := svc.state.Get(context.Background(), &fs_base_state.GetRequest{
			Key: vl.Voucher + "-" + vl.VerId,
		})
		if err != nil {
			return errno.ErrResponse(errno.ErrSystem)
		}
		if !resp.State.Ok {
			return errno.ErrResponse(resp.State)
		}

		//需要在等待验证状态
		if resp.Status == states.F_STATE_WAITING {
			code := strings.ToLower(in.Code)
			if bcrypt.CompareHashAndPassword([]byte(vl.Code), []byte(code)) != nil {
				//更新状态
				resp, err := svc.state.Upsert(context.Background(), &fs_base_state.UpsertRequest{
					Key:    vl.Voucher + "-" + vl.VerId,
					Status: states.F_STATE_OK,
				})
				if err != nil {
					return errno.ErrResponse(errno.ErrSystem)
				}
				return errno.ErrResponse(resp.State)
			} else {
				return errno.ErrResponse(errno.ErrInvalid)
			}
		}
	}

	return errno.ErrResponse(errno.ErrExpired)
}

func NewService(session *mgo.Session, channel messagecli.MessageChannel, state state.Service) Service {
	var svc Service
	{
		svc = &validateService{session: session, channel: channel, state: state}
	}
	return svc
}
