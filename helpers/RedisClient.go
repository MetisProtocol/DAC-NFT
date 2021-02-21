package helpers

import (
	"fmt"
	"github.com/astaxie/beego"
	redigo "github.com/gomodule/redigo/redis"
)

var pool *redigo.Pool

func init() {
	redisHost := beego.AppConfig.String("redis_host")
	redisPort, _ := beego.AppConfig.Int("redis_port")
	poolSize, _ := beego.AppConfig.Int("redis_size")
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
