package state

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
	"zskparker.com/foundation/pkg/constants"
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
	var err error

	wg := sync.WaitGroup{}

	errc := func(e error) {
		if err == nil {
			err = e
		}
		wg.Done()
	}

	wg.Add(1)
	go func() {
		errc(repo.addCacheStore(tag, status))
	}()

	wg.Add(1)
	go func() {
		errc(repo.addToDataStore(tag, status))
	}()

	wg.Wait()

	if err != nil {
		return err
	}
	return nil
}

func (repo *stateRepository) Get(tag string) (int64, error) {
	i, err := redis.Int64(repo.conn.Do("hget", "state_manage", "a2."+tag))
	if i == 0 {
		status := &store{}
		err = repo.collection(tag).Find(bson.M{"tag": tag}).One(status)
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

	e := repo.collection(tag).Update(bson.M{"tag": tag}, bson.M{"$set": bson.M{"modify_at": time.Now().UnixNano(), "status": status}})
	if e != nil && e == mgo.ErrNotFound {
		e = repo.collection(tag).Insert(s)
	}
	return e
}

func (repo *stateRepository) addCacheStore(tag string, status int64) error {
	_, err := repo.conn.Do("hset", "state_manage", "a2."+tag, status)
	return err
}

func (repo *stateRepository) collection(tag string) *mgo.Collection {
	return repo.session.DB(fs_constants.DB_BASE).C("stores_" + tag[len(tag)-1:])
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
