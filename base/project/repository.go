package project

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
)

type repository interface {
	Get(clientId string) (*project, error)

	Save(p *project) error

	Enable(projectId string, platform int64, open bool) error

	Close()

	Exists(name string) error

	Size() int
}

type project struct {
	Id        bson.ObjectId `bson:"_id"`
	Logo      string        `bson:"logo"`
	ZH        string        `bson:"zh"`
	EN        string        `bson:"en"`
	Desc      string        `bson:"desc"`
	CreateAt  int64         `bson:"create_at"`
	Platforms []*platform   `bson:"platforms"`
	Creator   string        `bson:"creator"`
	Session   string        `bson:"session"` //用来服务器与sso通讯用
	Level     int64         `bson:"level"`   // 等级限制
	Website   string        `bson:"website"` //网页 sso会跳转到这里
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

func (repo *projectRepository) Size() int {
	i, _ := repo.collection().Count()
	return i
}

func (repo *projectRepository) Exists(en string) error {
	i, err := repo.collection().Find(bson.M{"en": en}).Count()
	if err != nil && err == mgo.ErrNotFound {
		return nil
	} else {
		return err
	}
	if i > 0 {
		return errno.ERROR
	}
	return nil
}

func (repo *projectRepository) Enable(projectId string, platform int64, open bool) error {
	return repo.collection().Update(
		bson.M{"_id": bson.ObjectIdHex(projectId),
			"platforms": bson.M{"$elemMatch": bson.M{"platform": platform}},
		}, bson.M{"$set": bson.M{"enabled": open}})
}

func (repo *projectRepository) Get(clientId string) (*project, error) {
	p := &project{}
	err := repo.collection().Find(bson.M{"platforms": bson.M{"$elemMatch": bson.M{"client_id": clientId}}}).One(p)
	return p, err
}

func (repo *projectRepository) Save(p *project) error {
	return repo.collection().Insert(p)
}

func (repo *projectRepository) collection() *mgo.Collection {
	return repo.session.DB(fs_constants.DB_BASE).C("project")
}
