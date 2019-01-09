package validate

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strings"
	"time"
	"zskparker.com/foundation/base/message/cmd/messagecli"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/state"
	"zskparker.com/foundation/base/validate/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/osenv"
	"zskparker.com/foundation/pkg/states"
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
}

func (svc *validateService) GetRepo() repository {
	return &validateRepository{session: svc.session.Clone()}
}

//TODO 未加入频繁操作限制
func (svc *validateService) Create(ctx context.Context, in *fs_base_validate.CreateRequest) (*fs_base_validate.CreateResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	if len(in.Metadata.Ip) == 0 {
		return &fs_base_validate.CreateResponse{
			State: errno.ErrRequest,
		}, nil
	}

	//验证凭证(通过操作有时间等限制)
	voucher := in.Metadata.Ip + ";" + in.Func
	//有用户ID设置为凭证
	if len(in.Metadata.UserId) > 0 {
		voucher = in.Metadata.UserId + ";" + in.Func
	}
	voucher = utils.Md5(voucher)
	//查找最后一次同个操作的时间
	vl, err := repo.FindLast(voucher)
	if err == mgo.ErrNotFound {
		vl = &verification{
			CreateAt: time.Now().UnixNano() - 61*1e9,
			Func:     in.Func,
			Voucher:  voucher,
		}
		err = nil
	}

	if err != nil {
		return &fs_base_validate.CreateResponse{
			State: errno.ErrSystem,
		}, nil
	}

	//限制发送间隔时间
	if time.Now().UnixNano()-vl.CreateAt >= in.OnVerification.VoucherDuration*1e9 {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		var code string
		if in.OnVerification.CombinationMode == 1 {
			code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
		} else {
			code = strings.ToUpper(utils.GetRandomString())[0:8]
		}
		b, err := bcrypt.GenerateFromPassword([]byte(code), bcrypt.DefaultCost)
		if err != nil {
			return &fs_base_validate.CreateResponse{
				State: errno.ErrSystem,
			}, nil
		}

		vl.VerId = bson.NewObjectId()
		vl.CreateAt = time.Now().UnixNano()
		vl.Code = string(b)
		vl.State = states.F_STATE_WAITING

		err = repo.Create(vl)
		if err != nil {
			return &fs_base_validate.CreateResponse{State: errno.ErrSystem}, nil
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
			return &fs_base_validate.CreateResponse{
				State: errno.ErrSupport,
			}, nil
		}

		return &fs_base_validate.CreateResponse{
			VerId: vl.VerId.Hex(),
			State: errno.Ok,
		}, nil
	}

	return &fs_base_validate.CreateResponse{
		State: errno.ErrBusy,
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
		fmt.Println("err1", err)
		return &fs_base_validate.VerificationResponse{State: errno.ErrRequest}, nil
	}

	b := time.Now().UnixNano()-vl.CreateAt <= in.OnVerification.EffectiveTime*60*1e9

	fmt.Println("effective", in.OnVerification.EffectiveTime)
	fmt.Println("create", vl.CreateAt)
	fmt.Println("eff", b)

	//检查时间和操作
	if b {
		voucher := in.Metadata.Ip + ";" + in.Func
		if len(in.Metadata.UserId) > 16 {
			voucher = in.Metadata.UserId + ";" + in.Func
		}
		md := utils.Md5(voucher)
		if md != vl.Voucher {
			return &fs_base_validate.VerificationResponse{State: errno.ErrRequest}, nil
		}
		//需要在等待验证状态
		if vl.State == states.F_STATE_WAITING {
			code := strings.ToLower(in.Code)
			if bcrypt.CompareHashAndPassword([]byte(vl.Code), []byte(code)) == nil {

				err = repo.Update(in.VerId, states.F_STATE_OK)
				if err != nil {
					return &fs_base_validate.VerificationResponse{State: errno.ErrSystem}, nil
				}

				return &fs_base_validate.VerificationResponse{State: errno.Ok, To: vl.To}, nil
			} else {
				return &fs_base_validate.VerificationResponse{State: errno.ErrValidateCode}, nil
			}
		}
	}

	fmt.Println("err3")

	return &fs_base_validate.VerificationResponse{State: errno.ErrExpired}, nil
}

func NewService(session *mgo.Session, messagecli messagecli.Channel, state state.Service) Service {
	var svc Service
	{
		svc = &validateService{session: session, messagecli: messagecli, state: state}
	}
	return svc
}
