package face

import (
	"context"
	"gopkg.in/mgo.v2"
	"time"
	"zskparker.com/foundation/base/face/pb"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/base/reporter/cmd/reportercli"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/osenv"
)

type Service interface {
	Update(ctx context.Context, in *fs_base_face.UpdateRequest) (*fs_base.Response, error)

	Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error)

	Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error)

	AddFace(ctx context.Context, in *fs_base_face.AddFaceRequest) (*fs_base.Response, error)

	RemoveFace(ctx context.Context, in *fs_base_face.RemoveFaceRequest) (*fs_base.Response, error)
}

//独立模块
type faceService struct {
	session     *mgo.Session
	reportercli reportercli.Channel
}

func (svc *faceService) Update(ctx context.Context, in *fs_base_face.UpdateRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	panic(errno.ERROR)
}

func (svc *faceService) Compare(ctx context.Context, in *fs_base_face.CompareRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	fs, err := repo.Get(in.UserId)
	if err != nil {
		return errno.ErrResponse(errno.ErrRequest)
	}

	values, err := faceCompare(in.Base64Face, fs.FaceToken)
	if err != nil {
		return errno.ErrResponse(errno.ErrFaceProcess)
	}
	if values["error_code"].(float64) != 0 {
		return errno.ErrResponse(errno.ErrSystem)
	}
	//分数大于环境设置分数（默认为80.0）
	if values["score"].(float64) > osenv.GetFaceCompareScore() {
		//对比成功
		return errno.ErrResponse(errno.Ok)
	}
	return errno.ErrResponse(errno.ErrFaceCompare)
}

func (svc *faceService) Search(ctx context.Context, in *fs_base_face.SearchRequest) (*fs_base_face.SearchResponse, error) {
	repo := svc.GetRepo()
	defer repo.Close()
	panic(errno.ERROR)
}

func (svc *faceService) AddFace(ctx context.Context, in *fs_base_face.AddFaceRequest) (*fs_base.Response, error) {
	repo := svc.GetRepo()
	defer repo.Close()

	fs := &faceset{
		UserId:   in.UserId,
		CreateAt: time.Now().UnixNano(),
	}

	values, err := RegisterFace(in.Base64Face, in.UserId, "user")
	if err != nil {
		return errno.ErrResponse(errno.ErrSystem)
	}
	code := values["error_code"].(float64)
	if code == 223105 { //用户存在则更新用户
		//update face
		values, err = UpdateFace(in.Base64Face, in.UserId, "user")
		if err != nil {
			return errno.ErrResponse(errno.ErrSystem)
		}
		code = values["error_code"].(float64)
		if code != 0 {
			return errno.ErrResponse(errno.ErrSystem)
		}
		fs.FaceToken = values["face_token"].(string)
		err = repo.Upsert(fs)
		if err != nil {
			return errno.ErrResponse(errno.ErrRequest)
		}
	}
	if code != 0 {
		return errno.ErrResponse(errno.ErrRequest)
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

	values, err := DeleteFace(in.UserId, "user")
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
	return &faceRepository{session: svc.session.Clone()}
}

func NewService(session *mgo.Session, reportercli reportercli.Channel) Service {
	var svc Service
	{
		svc = &faceService{session: session, reportercli: reportercli}
	}
	return svc
}
