package strategy

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/base/pb"
)

type repository interface {
	Close()

	Get(projectId string) (*fs_base.ProjectStrategy, error)

	Upsert(s *fs_base.ProjectStrategy) error

	Size() int
}

type strategyRepository struct {
	session *mgo.Session
}

func (repo *strategyRepository) Close() {
	repo.session.Close()
}

func (repo *strategyRepository) Get(projectId string) (*fs_base.ProjectStrategy, error) {
	p := &fs_base.ProjectStrategy{}
	err := repo.collection().Find(bson.M{"projectid": projectId}).One(p)
	return p, err
}

func (repo *strategyRepository) Size() int {
	i, _ := repo.collection().Count()
	return i
}

func (repo *strategyRepository) Upsert(s *fs_base.ProjectStrategy) error {

	_, err := repo.collection().Upsert(bson.M{"projectid": s.ProjectId}, s)
	return err
}

func (repo *strategyRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("strategy")
}
