package helpers

import (
	"fmt"
	redigo "github.com/gomodule/redigo/redis"
)

var pool *redigo.Pool

func init() {
	redisHost := "127.0.0.1"
	redisPort := 6379
	poolSize := 20
	pool = redigo.NewPool(func() (redigo.Conn, error) {
		c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", redisHost, redisPort))
		if err != nil {
			return nil, err
		}
		return c, nil
	}, poolSize)
}

func Get() redigo.Conn {
	return pool.Get()
}
