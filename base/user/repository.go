package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Add(user *user) error

	Get(value, key string) (*user, error)

	Close()

	UpdatePassword(userId, password string) error

	UpdatePhone(userId, phone string) error

	UpdateEnterprise(userId, enterprise string) error

	UpdateEmail(userId, email string) error
}

type userRepository struct {
	session *mgo.Session
}

func (repo *userRepository) Close() {
	repo.session.Close()
}

func (repo *userRepository) Add(user *user) error {
	return repo.collection().Insert(user)
}

func (repo *userRepository) Get(value, key string) (*user, error) {
	user := &user{}
	err := repo.collection().Find(bson.M{key: value}).One(user)
	return user, err
}

func (repo *userRepository) UpdatePassword(userId, password string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"password": password}})
	return e
}

func (repo *userRepository) UpdatePhone(userId, phone string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"phone": phone}})
	return e
}

func (repo *userRepository) UpdateEnterprise(userId, enterprise string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"enterprise": enterprise}})
	return e
}

func (repo *userRepository) UpdateEmail(userId, email string) error {
	_, e := repo.collection().Upsert(bson.M{"user_id": userId}, bson.M{"$set": bson.M{"email": email}})
	return e
}

func (repo *userRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("user")
}

type user struct {
	UserId        string `bson:"user_id"`
	Username      string `bson:"username"`
	Password      string `bson:"password"`
	Enterprise    string `bson:"enterprise"`
	CreateAt      int64  `bson:"create_at"`
	Phone         string `bson:"phone"`
	Email         string `bson:"email"`
	Level         int64  `bson:"level"`
	FromProjectId string `bson:"from_project_id"`
	FromAppId     string `bson:"from_app_id"`
}
