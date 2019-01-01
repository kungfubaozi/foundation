package project

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Get(clientId string) (*project, error)

	Save(p *project) error

	Enable(projectId string, platform int64, open bool) error

	Close()
}

type project struct {
	Id        string      `bson:"project_id"`
	Logo      string      `bson:"logo"`
	ZH        string      `bson:"zh"`
	EN        string      `bson:"en"`
	Desc      string      `bson:"desc"`
	CreateAt  int64       `bson:"create_at"`
	Platforms []*platform `bson:"platforms"`
}

type platform struct {
	ClientId string `bson:"client_id"`
	Platform int64  `bson:"platform"`
	Enabled  bool   `bson:"enabled"`
}

type projectRepository struct {
	session *mgo.Session
}

func (repo *projectRepository) Close() {
	repo.session.Close()
}

func (repo *projectRepository) Enable(projectId string, platform int64, open bool) error {
	return repo.collection().Update(
		bson.M{"project_id": projectId,
			"platforms": bson.M{"$elemMatch": bson.M{"platform": platform}},
		}, bson.M{"$set": bson.M{"enabled": open}})
}

func (repo *projectRepository) Get(clientId string) (*project, error) {
	p := &project{}
	err := repo.collection().Find(bson.M{"$elemMatch": bson.M{"client_id": clientId}}).One(p)
	return p, err
}

func (repo *projectRepository) Save(p *project) error {
	return repo.collection().Insert(p)
}

func (repo *projectRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("project")
}
