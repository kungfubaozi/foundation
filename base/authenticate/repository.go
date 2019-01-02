package authenticate

import (
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"zskparker.com/foundation/base/authenticate/pb"
)

type repository interface {
	Add(auth *fs_base_authenticate.Authorize) error

	Get(userId, projectId, tokenAb string) (*fs_base_authenticate.Authorize, error)

	Close()
}

type authenticateRepository struct {
	conn redis.Conn
}

func (repo *authenticateRepository) Add(auth *fs_base_authenticate.Authorize) error {
	b, err := msgpack.Marshal(auth)
	if err != nil {
		return err
	}
	_, err = repo.conn.Do("hmset", "auth."+auth.UserId, auth.ClientId+auth.Ab, b)
	return err
}

func (repo *authenticateRepository) Get(userId, projectId, tokenAb string) (*fs_base_authenticate.Authorize, error) {

}

func (repo *authenticateRepository) Close() {
	repo.conn.Close()
}
