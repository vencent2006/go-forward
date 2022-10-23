package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) error {
	// 下面不能写rdb:=，因为要初始化全局变量
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize, // 连接池大小
	})

	_, err := client.Ping().Result()
	return err
}

func Close() {
	_ = client.Close()
}
