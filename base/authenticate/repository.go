package authenticate

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"zskparker.com/foundation/base/authenticate/pb"
)

type repository interface {
	Add(auth *fs_base_authenticate.Authorize) error

	Get(userId, clientId, relation string) (*fs_base_authenticate.Authorize, error)

	SizeOf(userId string) ([]interface{}, error)

	Del(userId, key string) error

	DelAll(userId string) error

	Close()
}

type authenticateRepository struct {
	conn redis.Conn
}

func (repo *authenticateRepository) Del(userId, key string) error {
	_, err := repo.conn.Do("HDEL", "auth."+userId, key)
	return err
}

func (repo *authenticateRepository) DelAll(userId string) error {
	_, err := repo.conn.Do("DEL", "auth."+userId)
	return err
}

func (repo *authenticateRepository) Add(auth *fs_base_authenticate.Authorize) error {
	b, err := msgpack.Marshal(auth)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("hset", "auth."+auth.UserId, fmt.Sprintf("%s.%s", auth.ClientId, auth.Relation), b)
	return err
}

func (repo *authenticateRepository) Get(userId, clientId, relation string) (*fs_base_authenticate.Authorize, error) {
	v, err := redis.Bytes(repo.conn.Do("hget", "auth."+userId, fmt.Sprintf("%s.%s", clientId, relation)))
	if err != nil {
		return nil, err
	}
	a := &fs_base_authenticate.Authorize{}
	err = msgpack.Unmarshal(v, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (repo *authenticateRepository) SizeOf(userId string) ([]interface{}, error) {
	v, err := redis.Values(repo.conn.Do("hkeys", "auth."+userId))
	if err == redis.ErrNil {
		err = nil
	}
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, nil
	}
	return v, nil
}

func (repo *authenticateRepository) Close() {
	repo.conn.Close()
}
