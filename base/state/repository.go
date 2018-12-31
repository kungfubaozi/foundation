package state

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type repository interface {
	Close()

	Upset(tag string, status int64) error

	Get(tag string) (int64, error)
}

type stateRepository struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *stateRepository) Upset(tag string, status int64) error {
	err := make(chan error, 2)
	go func() {
		err <- repo.addCacheStore(tag, status)
	}()

	go func() {
		err <- repo.addToDataStore(tag, status)
	}()

	e := <-err
	if e != nil {
		return e
	}
	return nil
}

func (repo *stateRepository) Get(tag string) (int64, error) {
	i, err := redis.Int64(repo.conn.Do("hmget", "state_manage", "a2."+tag))
	if i == 0 {
		status := &store{}
		err = repo.collection().Find(bson.M{"tag": tag}).One(status)
		if err != nil {
			if err == mgo.ErrNotFound {
				return -1, mgo.ErrNotFound
			}
		}
		err = repo.addCacheStore(status.Tag, status.Status)
		if err != nil {
			return -1, err
		}
		return status.Status, nil
	}
	return i, nil

}

func (repo *stateRepository) addToDataStore(tag string, status int64) error {
	s := &store{
		Tag:      tag,
		Status:   status,
		CreateAt: time.Now().UnixNano(),
		ModifyAt: time.Now().UnixNano(),
	}
	e := repo.collection().Update(bson.M{"tag": tag}, bson.M{"$set": bson.M{"modify_at": time.Now().UnixNano(), "status": status}})
	if e != nil && e == mgo.ErrNotFound {
		e = repo.collection().Insert(s)
	}
	return e
}

func (repo *stateRepository) addCacheStore(tag string, status int64) error {
	_, err := repo.conn.Do("hmset", "state_manage", "a2."+tag, status)
	return err
}

func (repo *stateRepository) collection() *mgo.Collection {
	return repo.session.DB("foundation").C("stores")
}

func (repo *stateRepository) Close() {
	repo.session.Close()
	repo.conn.Close()
}

type store struct {
	Tag      string `json:"tag" bson:"tag"`
	Status   int64  `json:"status" bson:"status"`
	CreateAt int64  `json:"create_at" bson:"create_at"`
	ModifyAt int64  `json:"modify_at" bson:"modify_at"`
}
