package review

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
)

type model struct {
	Id        string `bson:"_id"`
	Creator   string `bson:"creator"`
	CreateAt  int64  `bson:"create_at"`
	ProjectId string `bson:"project_id"`
	Version   int64  `bson:"version"`
	Pairs     []Pair `bson:"pairs"`

	/*
		pair里的value是作为key的类型
		支持以下几种:
		String,Int,Float,ImageUrl
	*/
}

type Pair struct {
	Key   string `bson:"key"`
	Value string `bson:"value"`
	Force bool   `bson:"force"` // 强制
}

//用户审核是对应的项目中的审核，因为项目如果不开启审核是不会经过审核流程的
type UserReview struct {
	UserId    string `bson:"user_id"`
	Version   int64  `bson:"version"` //审核的版本
	Pairs     []Pair `bson:"pairs"`
	CreateAt  int64  `bson:"create_at"`
	State     int64  `bson:"state"` //审核状态
	ProjectId string `bson:"project_id"`
	ReviewId  string `bson:"review_id"`  //对应审核的id
	FailedMsg string `bson:"failed_msg"` //审核不通过的消息
}

type reviewRepository struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *reviewRepository) IsReviewed(userId string) bool {
	panic("implement me")
}

func (repo *reviewRepository) Get(projectId string) (*model, error) {
	panic("implement me")
}

func (repo *reviewRepository) New(m *model) error {
	panic("implement me")
}

func (repo *reviewRepository) GetUserReviewDetail(userId string) (*UserReview, error) {
	panic("implement me")
}

func (repo *reviewRepository) Set(review *UserReview) error {
	panic("implement me")
}

func (repo *reviewRepository) collection() {

}

func (repo *reviewRepository) userReviewCollection(userId string) *mgo.Collection {
	return repo.session.DB("fds_review").C(fmt.Sprintf("%s_%s", "review", userId[len(userId)-2:]))
}

type repository interface {
	IsReviewed(userId string) bool

	Get(projectId string) (*model, error)

	New(m *model) error

	GetUserReviewDetail(userId string) (*UserReview, error)

	Set(review *UserReview) error
}
