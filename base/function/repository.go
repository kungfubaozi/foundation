package function

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
)

type repository interface {
	Add(function *Model) error

	Update(function *Model) error

	Remove(value, key string) error

	Get(tag string) (*Model, error)

	FindByFunc(function string) (*Model, error)

	Close()
}

type Model struct {
	Func     string `bson:"func"`
	Tag      string `bson:"tag"` //api+tag
	CreateAt int64  `bson:"create_at"`
	ZH       string `bson:"zh"`
	EN       string `bson:"en"`
	Fcv      int64  `bson:"fcv"` //首选验证模式
	Level    int64  `bson:"level"`
	//项目集成时会带入session来请求API
	//新加入的功能会加入session防止API一致
	//如果没有注册功能，默认会对比session
	//功能里的fcv检查条件如果没有session检查条件，则不检查session
}

type functionRepository struct {
	session *mgo.Session
}

func (repo *functionRepository) Update(function *Model) error {
	panic(errno.ERROR)
}

func (repo *functionRepository) Add(function *Model) error {
	return repo.collection().Insert(function)
}

func (repo *functionRepository) Remove(value, key string) error {
	return repo.collection().Remove(bson.M{key: value})
}

func (repo *functionRepository) Get(tag string) (*Model, error) {
	f := &Model{}
	err := repo.collection().Find(bson.M{"tag": tag}).One(f)
	return f, err
}

func (repo *functionRepository) FindByFunc(fu string) (*Model, error) {
	f := &Model{}
	err := repo.collection().Find(bson.M{"func": fu}).One(f)
	return f, err
}

func (repo *functionRepository) collection() *mgo.Collection {
	return repo.session.DB(fs_constants.DB_BASE).C("functions")
}

func (repo *functionRepository) Close() {
	repo.session.Close()
}
