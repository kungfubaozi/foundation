package blacklist

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
)

const (
	USER    = 1
	IP      = 2
	ACCOUNT = 3
	DEVICE  = 4
)

type repository interface {
	Add(m *model, t int) error

	Get(value string, t int) error

	Close()
}

type blacklistRepository struct {
	session *mgo.Session
}

func (repo *blacklistRepository) Get(value string, t int) error {
	m := &model{}
	err := repo.collection(value, t).Find(bson.M{"tag": value}).One(m)
	if err != nil && err == mgo.ErrNotFound {
		return nil
	}
	if len(m.Tag) > 1 {
		return errno.ERROR
	}
	return err
}

func (repo *blacklistRepository) collection(data string, t int) *mgo.Collection {
	var c string
	switch t {
	case USER:
		c = "user"
		break
	case IP:
		c = "ip"
		break
	case ACCOUNT:
		c = "account"
		break
	case DEVICE:
		c = "device"
		break
	}
	return repo.session.DB(fs_constants.DB_BASE).C(fmt.Sprintf("%s_%s_%s", "blacklist", c, data[0:1]))
}

func (repo *blacklistRepository) Add(m *model, t int) error {
	_, err := repo.collection(m.Tag, t).Upsert(bson.M{"tag": m.Tag}, m)
	return err
}

func (repo *blacklistRepository) Close() {
	repo.session.Close()
}

type model struct {
	Tag     string `bson:"tag"`
	Creator string `bson:"creator"`
}
