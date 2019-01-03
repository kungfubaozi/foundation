package function

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Add(function *function) error

	Update(function *function) error

	Remove(value, key string) error

	Get(api string) (*function, error)

	FindByFunc(function string) (*function, error)

	Close()
}

type function struct {
	Func         string `bson:"func"`
	API          string `bson:"api"`
	CreateAt     string `bson:"create_at"`
	ZH           string `bson:"zh"`
	EN           string `bson:"en"`
	Verification bool   `bson:"verification"` //是否需要验证
	Fcv          int64  `bson:"fcv"`          //首选验证模式
	Level        int64  `bson:"level"`
}

type functionRepository struct {
	session *mgo.Session
}

func (repo *functionRepository) Update(function *function) error {

}

func (repo *functionRepository) Add(function *function) error {

}

func (repo *functionRepository) Remove(value, key string) error {
	return repo.collection().Remove(bson.M{key: value})
}

func (repo *functionRepository) Get(api string) (*function, error) {
	f := &function{}
	err := repo.collection().Find(bson.M{"api": api}).One(f)
	return f, err
}

func (repo *functionRepository) FindByFunc(fu string) (*function, error) {
	f := &function{}
	err := repo.collection().Find(bson.M{"func": fu}).One(f)
	return f, err
}

func (repo *functionRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("functions")
}

func (repo *functionRepository) Close() {
	repo.session.Close()
}
