package fs_sync_lock

import "github.com/garyburd/redigo/redis"

type SyncLock struct {
	pool *redis.Pool
}

func Create(pool *redis.Pool) *SyncLock {
	return &SyncLock{pool: pool}
}

func (l *SyncLock) Unlock(function, tag string) {

}

func (l *SyncLock) Lock(function, tag string) {

}
