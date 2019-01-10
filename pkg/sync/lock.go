package fs_redisync

import (
	"github.com/garyburd/redigo/redis"
	"zskparker.com/foundation/base/pb"
	"zskparker.com/foundation/pkg/errno"
	"zskparker.com/foundation/pkg/utils"
)

type Redisync struct {
	pool *redis.Pool
}

func Create(pool *redis.Pool) *Redisync {
	return &Redisync{pool: pool}
}

var delScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

func (l *Redisync) Unlock(function, tag string) {
	delScript.Do(l.pool.Get(), utils.Md5(function+tag), 1)
}

func (l *Redisync) Lock(function, tag string, timeout int64) *fs_base.State {
	lockReply, err := l.pool.Get().Do("SET", utils.Md5(function+tag), 1, "ex", timeout, "nx")
	if err != nil {
		return errno.ErrSystem
	}
	if lockReply == "OK" {
		return nil
	} else {
		return errno.ErrBusy
	}
}
