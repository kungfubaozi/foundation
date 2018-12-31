package face

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	Get(userId string) (*faceset, error)

	Upsert(faceset *faceset) error

	Delete(userId string) error

	Close()
}

type faceRepository struct {
	session *mgo.Session
}

func (repo *faceRepository) Get(userId string) (*faceset, error) {
	coll := repo.collection(userId)
	fs := &faceset{}
	err := coll.Find(bson.M{"user_id": userId}).One(fs)
	return fs, err
}

func (repo *faceRepository) Upsert(faceset *faceset) error {
	coll := repo.collection(faceset.UserId)
	_, err := coll.Upsert(bson.M{"user_id": faceset.UserId}, faceset)
	return err
}

func (repo *faceRepository) Delete(userId string) error {
	coll := repo.collection(userId)

	return coll.Remove(bson.M{"user_id": userId})
}

func (repo *faceRepository) Close() {
	repo.session.Close()
}

func (repo *faceRepository) collection(userId string) *mgo.Collection {
	return repo.session.DB("foundation").C("faceset_" + userId[0:1])
}
