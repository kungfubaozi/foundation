package register

import "gopkg.in/mgo.v2"

type repository interface {
	Close()
}

type registerRepository struct {
	session *mgo.Session
}

func (repo *registerRepository) Close() {
	repo.session.Close()
}
