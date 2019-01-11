package function

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/pkg/errno"
)

type repository interface {
	Add(function *FunctionModel) error

	Update(function *FunctionModel) error

	Remove(value, key string) error

	Get(api string) (*FunctionModel, error)

	FindByFunc(function string) (*FunctionModel, error)

	Close()
}

type FunctionModel struct {
	Func        string `bson:"func"`
	API         string `bson:"api"`
	CreateAt    string `bson:"create_at"`
	ZH          string `bson:"zh"`
	EN          string `bson:"en"`
	Fcv         int64  `bson:"fcv"` //首选验证模式
	Level       int64  `bson:"level"`
	FromSession string `bson:"from_session"` //功能对应的项目会话
	//项目集成时会带入session来请求API
	//新加入的功能会加入session防止API一致
	//如果没有注册功能，默认会对比session
	//功能里的fcv检查条件如果没有session检查条件，则不检查session
}

type functionRepository struct {
	session *mgo.Session
}

func (repo *functionRepository) Update(function *FunctionModel) error {
	panic(errno.ERROR)
}

func (repo *functionRepository) Add(function *FunctionModel) error {
	return repo.collection().Insert(function)
}

func (repo *functionRepository) Remove(value, key string) error {
	return repo.collection().Remove(bson.M{key: value})
}

func (repo *functionRepository) Get(api string) (*FunctionModel, error) {
	f := &FunctionModel{}
	err := repo.collection().Find(bson.M{"api": api}).One(f)
	return f, err
}

func (repo *functionRepository) FindByFunc(fu string) (*FunctionModel, error) {
	f := &FunctionModel{}
	err := repo.collection().Find(bson.M{"func": fu}).One(f)
	return f, err
}

func (repo *functionRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("functions")
}

func (repo *functionRepository) Close() {
	repo.session.Close()
}
