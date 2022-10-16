package redis

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init() error {
	// 下面不能写rdb:=，因为要初始化全局变量
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})

	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
