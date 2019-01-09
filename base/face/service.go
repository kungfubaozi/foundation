package face

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/base/user"
	"zskparker.com/foundation/base/user/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/osenv"
)

type Service interface {
	Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error)

	Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error)

	Upsert(ctx context.Context, in *fs_base_face.UpsertRequest) (*fs_base.Response, error)

	RemoveFace(ctx context.Context, in *fs_base_face.RemoveFaceRequest) (*fs_base.Response, error)
}

//独立模块
type faceService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
	pool        *redis.Pool
	usercli     user.Service
}

func (svc *faceService) Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	fs, err := repo.Get(in.UserId)
	if err != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}

	token, err := repo.GetToken()
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	values, err := faceCompare(in.Base64Face, fs.FaceToken, token)
	if err != nil {
		return errno.ErrResponse(errno.ErrFaceProcess)
	}
	if values["error_code"].(float64) != 0 {
		return errno.ErrResponse(errno.ErrSystem)
	}
	//分数大于环境设置分数（默认为80.0）
	if values["result"].(map[string]interface{})["score"].(float64) > osenv.GetFaceCompareScore() {
		//对比成功
		return errno.ErrResponse(errno.Ok)
	}
	return errno.ErrResponse(errno.ErrFaceCompare)
}

func (svc *faceService) Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	token, err := repo.GetToken()

	resp := func(state *fs_base.State) (*fs_base_face.SearchResponse, error) {
		return &fs_base_face.SearchResponse{State: state}, nil
	}

	if err != nil {
		return resp(errno.ErrSystem)
	}
	values, err := faceSearch(in.Base64Face, "user", token)
	if err != nil {
		return resp(errno.ErrSystem)
	}
	if values["error_code"].(float64) != 0 {
		return resp(errno.ErrSystem)
	}

	userList := values["result"].(map[string]interface{})["user_list"].([]interface{})
	if len(userList) > 0 {
		face := userList[0].(map[string]interface{})
		if face["score"].(float64) > osenv.GetFaceCompareScore() {
			userId := face["user_id"].(string)
			//查找用户
			ur, err := svc.usercli.FindByUserId(context.Background(), &fs_base_user.FindRequest{
				Value: userId,
			})
			if err != nil {
				return resp(errno.ErrSystem)
			}
			if !ur.State.Ok {
				return resp(ur.State)
			}
			return &fs_base_face.SearchResponse{
				State:  errno.Ok,
				UserId: userId,
				Level:  ur.Level,
			}, nil
		}
	}
	return resp(errno.ErrInvalidFace)
}

func (svc *faceService) Upsert(ctx context.Context, in *fs_base_face.UpsertRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	fs := &faceset{
		UserId:   in.UserId,
		CreateAt: time.Now().UnixNano(),
	}

	token, err := repo.GetToken()
	if err != nil {
		fmt.Println("token", err)
		return errno.ErrResponse(errno.ErrSystem)
	}

	values, err := RegisterFace(in.Base64Face, in.UserId, "user", token)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	code := values["error_code"].(float64)
	if code == 223105 { //用户存在则更新用户
		//update face
		values, err = UpdateFace(in.Base64Face, in.UserId, "user", token)
		if err != nil {
			return errno.ErrResponse(errno.ErrSystem)
		}
		code = values["error_code"].(float64)
		if code != 0 {
			return errno.ErrResponse(errno.ErrSystem)
		}
	}

	if code != 0 {
		return errno.ErrResponse(errno.ErrRequest)
	} else {
		fs.FaceToken = values["result"].(map[string]interface{})["face_token"].(string)
		err = repo.Upsert(fs)
		if err != nil {
			return errno.ErrResponse(errno.ErrRequest)
		}
	}

	return errno.ErrResponse(errno.Ok)
}

func (svc *faceService) RemoveFace(ctx context.Context, in *fs_base_face.RemoveFaceRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	err := repo.Delete(in.UserId)
	if err != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}

	token, err := repo.GetToken()
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}

	values, err := DeleteFace(in.UserId, "user", token)
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	code := values["error_code"].(float64)
	if code != 0 && code != 223103 { //223103 用户不存在
		return errno.ErrResponse(errno.ErrRequest)
	}
	return errno.ErrResponse(errno.Ok)
}

func (svc *faceService) GetRepo() repository {
	return &faceRepository{session: svc.session.Clone(), conn: svc.pool.Get()}
}

func NewService(session *mgo.Session, reportercli reportercli.Channel, pool *redis.Pool, usercli user.Service) Service {
	var svc Service
	{
		svc = &faceService{session: session, reportercli: reportercli, pool: pool, usercli: usercli}
	}
	return svc
}
