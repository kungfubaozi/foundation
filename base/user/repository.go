package user

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/pkg/constants"
	"zskparker.com/foundation/pkg/errno"
)

type repository interface {
	Add(user *User) error

	Get(value, key string) (*User, error)

	Close()

	UpdatePassword(userId, password string) error

	UpdatePhone(userId, phone string) error

	UpdateEnterprise(userId, enterprise string) error

	UpdateEmail(userId, email string) error

	FindAdmin() int

	FindSame(phone, email, enterprise string) error
}

type userRepository struct {
	session *mgo.Session
}

func (repo *userRepository) Close() {
	repo.session.Close()
}

func (repo *userRepository) FindAdmin() int {
	i, err := repo.collection().Find(bson.M{"level": fs_constants.LEVEL_ADMIN}).Count()
	if err != nil && err == mgo.ErrNotFound {
		i = 0
		err = nil
	}
	if err != nil {
		return -1
	}
	return i
}

func (repo *userRepository) FindSame(phone, email, enterprise string) error {
	var b []bson.M
	if len(phone) > 0 {
		b = append(b, bson.M{"phone": phone})
	}
	if len(email) > 0 {
		b = append(b, bson.M{"email": email})
	}
	if len(enterprise) > 0 {
		b = append(b, bson.M{"enterprise": enterprise})
	}
	i, err := repo.collection().Find(bson.M{"$or": b}).Count()
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	if i > 0 {
		return errno.ERROR
	}
	return nil
}

func (repo *userRepository) Add(user *User) error {
	return repo.collection().Insert(user)
}

func (repo *userRepository) Get(value, key string) (*User, error) {
	user := &User{}
	var err error
	if key == "_id" {
		err = repo.collection().Find(bson.M{key: bson.ObjectIdHex(value)}).One(user)
	} else {
		err = repo.collection().Find(bson.M{key: value}).One(user)
	}

	return user, err
}

func (repo *userRepository) UpdatePassword(id, password string) error {
	_, e := repo.collection().Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"password": password}})
	return e
}

func (repo *userRepository) UpdatePhone(id, phone string) error {
	_, e := repo.collection().Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"phone": phone}})
	return e
}

func (repo *userRepository) UpdateEnterprise(id, enterprise string) error {
	_, e := repo.collection().Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"enterprise": enterprise}})
	return e
}

func (repo *userRepository) UpdateEmail(id, email string) error {
	_, e := repo.collection().Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"email": email}})
	return e
}

func (repo *userRepository) collection() *mgo.Collection {
	return repo.session.DB(fs_constants.DB_USER).C("user")
}

type User struct {
	UserId        bson.ObjectId `bson:"_id"`
	Username      string        `bson:"username"`
	RealName      string        `bson:"real_name"`
	Password      string        `bson:"password"`
	Enterprise    string        `bson:"enterprise"`
	CreateAt      int64         `bson:"create_at"`
	Phone         string        `bson:"phone"`
	Email         string        `bson:"email"`
	Level         int64         `bson:"level"`
	FromProjectId string        `bson:"from_project_id"`
	FromClientId  string        `bson:"from_client_id"`
	Type          int64         `bson:"type"`
}
