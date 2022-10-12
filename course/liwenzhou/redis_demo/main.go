package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() error {
	// 下面不能写rdb:=，因为要初始化全局变量
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})

	_, err := rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	// Result, Val，看业务需要
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		// 值不存在
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}

}

func hgetDemo() {
	v, err := rdb.HGetAll("user").Result()
	if err != nil {
		// redis.Nil
		// 其他错误
		fmt.Printf("hgetall failed, err:%v\n", err)
		return
	}
	fmt.Println(v)

	val, err := rdb.HMGet("user", "name", "age").Result()
	if err != nil {
		fmt.Printf("HMGet failed, err:%v\n", err)
		return
	} else {
		fmt.Println(val)
	}

	v3 := rdb.HGet("user", "age").Val()
	fmt.Println(v3)
}

// zsetDemo 操作zset示例
func zsetDemo() {
	// key
	zsetKey := "language_rank"
	// value
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	//defer cancel()

	// ZADD
	err := rdb.ZAdd(zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func pipelineDemo() {
	pipe := rdb.Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

// 事务
func transactionDemo() {
	pipe := rdb.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)

	_, err := pipe.Exec()
	fmt.Println(incr.Val(), err)
}

// watch
func watchDemo() {
	// 监视 watch_count的值，并在值不变的前提下将其值+1
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			// 业务逻辑
			time.Sleep(time.Second * 5) // 创造机会引起watch冲突
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Printf("tx exec failed, err:%v\n", err)
		return
	}
	fmt.Println("tx exec success")
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Printf("init redis client failed, err:%v\n", err)
		return
	}
	defer rdb.Close()
	fmt.Println("connect redis success...")

	//redisExample()

	//hgetDemo()

	//zsetDemo()

	//pipelineDemo()

	//transactionDemo()

	watchDemo()
}
