package invite

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Add(m *model) error

	Get(key, user string) (*model, error)

	Close()

	FindInvite(code string) (*model, error)
}

type inviteRepository struct {
	session *mgo.Session
}

func (repo *inviteRepository) Get(key, user string) (*model, error) {
	m := &model{}
	err := repo.collection().Find(bson.M{key: user}).One(m)
	return m, err
}

func (repo *inviteRepository) FindInvite(code string) (*model, error) {
	m := &model{}
	err := repo.collection().Find(bson.M{"code": code}).One(m)
	return m, err
}

func (repo *inviteRepository) Add(m *model) error {
	return repo.collection().Insert(m)
}

func (repo *inviteRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("invite")
}

func (repo *inviteRepository) Close() {
	repo.session.Close()
}

type model struct {
	Phone        string        `bson:"phone"`
	Email        string        `bson:"email"`
	Enterprise   string        `bson:"enterprise"`
	Username     string        `bson:"username"`
	RealName     string        `bson:"real_name"`
	Level        int64         `bson:"level"`
	ExpireAt     int64         `bson:"expire_at"` //过期时间
	CreateAt     int64         `bson:"create_at"`
	Code         string        `bson:"code"`
	InviteId     bson.ObjectId `bson:"_id"`
	FromProject  string        `bson:"from_project"`
	FromClientId string        `bson:"from_client_id"`
}
