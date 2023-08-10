package cache

import (
	"crac55/common/clog"
	"crac55/common/conf"

	goredis "github.com/go-redis/redis"
)

var RedisClient *goredis.Client

// Init 初始化redis
func Init() {
	c := conf.AppConf.Redis
	RedisClient = goredis.NewClient(&goredis.Options{
		Addr:     c.HostPort,
		Password: c.Password,
		DB:       c.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		clog.Log().Fatalln(err.Error())
	}
}
