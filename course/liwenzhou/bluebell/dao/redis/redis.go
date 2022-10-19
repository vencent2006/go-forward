package redis

import (
	"fmt"
	"bluebell/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) error {
	// 下面不能写rdb:=，因为要初始化全局变量
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
