package strategy

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/constants"
)

type repository interface {
	Close()

	Get(session string) (*fs_base.Strategy, error)

	Upsert(s *fs_base.Strategy) error

	Size() int
}

type strategyRepository struct {
	session *mgo.Session
	conn    redis.Conn
}

func (repo *strategyRepository) Close() {
	repo.session.Close()
}

var name = "root_strategy"

func (repo *strategyRepository) Get(session string) (*fs_base.Strategy, error) {
	p := &fs_base.Strategy{}

	s, err := redis.Bytes(repo.conn.Do("get", name))
	if err != nil && err == redis.ErrNil {
		err := repo.collection().Find(bson.M{"session": session}).One(p)
		if err != nil {
			return p, err
		}
		err = repo.addToCache(p)
	}

	if err != nil {
		return p, err
	}

	if len(s) > 0 {
		err = msgpack.Unmarshal(s, p)
	}

	return p, err
}

func (repo *strategyRepository) Size() int {
	i, _ := repo.collection().Count()
	return i
}

func (repo *strategyRepository) addToCache(s *fs_base.Strategy) error {
	b, err := msgpack.Marshal(s)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("set", name, b)

	return err
}

func (repo *strategyRepository) Upsert(s *fs_base.Strategy) error {
	_, err := repo.collection().Upsert(bson.M{"session": s.Session}, s)
	err = repo.addToCache(s)
	return err
}

func (repo *strategyRepository) collection() *mgo.Collection {
	return repo.session.DB(fs_constants.DB_BASE).C("strategy")
}
