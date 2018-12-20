package db

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"os"
)

//token过期时间
var TOKEN_EX_TIME string

func CreatePool() *redis.Pool {

	host := os.Getenv("REDIS_HOST")

	pool := &redis.Pool{
		MaxIdle: 20,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)

			if err != nil {
				return nil, err
			}

			return c, err
		},
	}

	return pool
}

// 建立与mongo的连接
func CreateSession(addr string) (*mgo.Session, error) {

	s, err := mgo.Dial(addr)

	if err != nil {
		fmt.Println("connect to mongo error:", err)
		return nil, err
	}

	s.SetMode(mgo.Monotonic, true)

	return s, err

}
