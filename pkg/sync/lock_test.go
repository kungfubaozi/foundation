package fs_redisync

import (
	"fmt"
	"testing"
	"zskparker.com/foundation/pkg/db"
)

func TestCreate(t *testing.T) {
	pool := db.CreatePool("192.168.2.60:6379")
	redisync := Create(pool)
	//defer func() {
	//	redisync.Unlock("173e2bb3f601","2312235435")
	//}()

	err := redisync.Lock("173e2bb3f601", "2312235435", 60)
	fmt.Println(err)
	//err = redisync.Lock("173e2bb3f601", "2312235435", 60)
	//fmt.Println(err)
}
