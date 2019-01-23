package authorization

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type model struct {
	UserId    string `bson:"user_id"`
	ProjectId string `bson:"project_id"`
	CreateAt  int64  `bson:"create_at"`
}

type repository interface {
	Find(userId, projectId string) (*model, error)

	Sync(userId, projectId string) error
}

type authorizationRepository struct {
	session *mgo.Session
}

func (repo *authorizationRepository) collection(userId, projectId string) *mgo.Collection {
	return repo.session.DB("fds_sync").C(fmt.Sprintf("%s_%s", projectId, userId[len(userId)-1:]))
}

func (repo *authorizationRepository) Find(userId, projectId string) (*model, error) {
	m := &model{}
	err := repo.collection(userId, projectId).Find(bson.M{"user_id": userId, "project_id": projectId}).One(m)
	return m, err
}

func (repo *authorizationRepository) Sync(userId, projectId string) error {
	m := &model{
		UserId:    userId,
		ProjectId: projectId,
		CreateAt:  time.Now().UnixNano(),
	}
	err := repo.collection(userId, projectId).Insert(m)
	return err
}
