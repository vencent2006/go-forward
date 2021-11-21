/**
 * @Author: vincent
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2021/11/21 16:18
 */

package services

import (
	"context"
	"errors"
	"go-examples/course/handwriting-web-inf/code_28/framework"
	"go-examples/course/handwriting-web-inf/code_28/framework/contract"
	"go-examples/course/handwriting-web-inf/code_28/framework/provider/redis"
	"sync"
	"time"

	redisv8 "github.com/go-redis/redis/v8"
)

// RedisCache 代表Redis缓存
type RedisCache struct {
	container framework.Container
	client    *redisv8.Client
	lock      sync.RWMutex
}

// NewRedisCache 初始化redis服务
func NewRedisCache(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	if !container.IsBind(contract.RedisKey) {
		err := container.Bind(&redis.RedisProvider{})
		if err != nil {
			return nil, err
		}
	}

	// 获取redis服务配置，并且实例化redis.Client
	redisService := container.MustMake(contract.RedisKey).(contract.RedisService)
	client, err := redisService.GetClient(redis.WithConfigPath("cache"))
	if err != nil {
		return nil, err
	}

	// 返回RedisCache实例
	obj := &RedisCache{
		container: container,
		client:    client,
		lock:      sync.RWMutex{}, // todo: 这个lock不用是pointer吗
	}

	return obj, nil
}

// GetObj 获取某个key对应的对象，对象必须实现 https://pkg.go.dev/encoding#BinaryUnMarshaler
func (r *RedisCache) GetObj(ctx context.Context, key string, model interface{}) error {
	cmd := r.client.Get(ctx, key)
	if errors.Is(cmd.Err(), redisv8.Nil) {
		return ErrKeyNotFound
	}

	err := cmd.Scan(model)
	if err != nil {
		return err
	}

	return nil
}

// GetMany 获取某些key对应的值
func (r *RedisCache) GetMany(ctx context.Context, keys []string) (map[string]string, error) {
	pipeline := r.client.Pipeline()
	vals := make(map[string]string)
	cmds := make([]*redisv8.StringCmd, 0, len(keys))

	for _, key := range keys {
		cmds = append(cmds, pipeline.Get(ctx, key))
	}

	_, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}

	errs := make([]string, 0, len(keys))
	for _, cmd := range cmds {
		val, err := cmd.Result()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		key := cmd.Args()[1].(string)
		vals[key] = val
	}
	return vals, nil
}

// Set 设置某个key和值到缓存，带超时时间
func (r *RedisCache) Set(ctx context.Context, key string, val string, timout time.Duration) error {
	return r.client.Set(ctx, key, val, timout).Err()
}

// SetObj 设置某个key和对象到缓存，对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler
func (r *RedisCache) SetObj(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	return r.client.Set(ctx, key, val, timeout).Err()
}

// SetMany 设置多个key和值到缓存
func (r *RedisCache) SetMany(ctx context.Context, data map[string]string, timeout time.Duration) error {
	pipeline := r.client.Pipeline()
	cmds := make([]*redisv8.StatusCmd, 0, len(data))
	for k, v := range data {
		cmds = append(cmds, pipeline.Set(ctx, k, v, timeout))
	}
	_, err := pipeline.Exec(ctx)
	return err
}

// SetForever 设置某个key和值到缓存，不带超时时间1
func (r *RedisCache) SetForever(ctx context.Context, key string, val string) error {
	return r.client.Set(ctx, key, val, NoneDuration).Err()
}

// SetForeverObj 设置某个key和对象到缓存，不带超时时间，对象必须实现 https://pkg.go.dev/encoding#BinaryMarshaler
func (r *RedisCache) SetForeverObj(ctx context.Context, key string, val interface{}) error {
	return r.client.Set(ctx, key, val, NoneDuration).Err()
}

// SetTTL 设置某个key的超时时间
func (r *RedisCache) SetTTL(ctx context.Context, key string, timeout time.Duration) error {
	return r.client.Expire(ctx, key, timeout).Err()
}

// GetTTL 获取某个key的超时时间
func (r *RedisCache) GetTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.TTL(ctx, key).Result()
}

// Remember
func (r *RedisCache) Remember(ctx context.Context, key string, timeout time.Duration, rememberFunc contract.RememberFunc, obj interface{}) error {
	err := r.GetObj(ctx, key, obj)
	if err == nil {
		return nil
	}

	if !errors.Is(err, ErrKeyNotFound) {
		return err
	}

	// key not found
	objNew, err := rememberFunc(ctx, r.container)
	if err != nil {
		return err
	}

	// todo: SetObj后，里面就GetObj，是个什么操作呢，而且如果是主从分离，也不一定能读到
	if err := r.SetObj(ctx, key, objNew, timeout); err != nil {
		return err
	}
	if err := r.GetObj(ctx, key, obj); err != nil {
		return err
	}
	return nil
}
func (r *RedisCache) Calc(ctx context.Context, key string, step int64) (int64, error) {
	return r.client.IncrBy(ctx, key, step).Result()
}

func (r *RedisCache) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.IncrBy(ctx, key, 1).Result()
}

func (r *RedisCache) Decrement(ctx context.Context, key string) (int64, error) {
	return r.client.IncrBy(ctx, key, -1).Result()
}

func (r *RedisCache) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *RedisCache) DelMany(ctx context.Context, keys []string) error {
	pipline := r.client.Pipeline()
	cmds := make([]*redisv8.IntCmd, 0, len(keys))
	for _, key := range keys {
		cmds = append(cmds, pipline.Del(ctx, key))
	}
	_, err := pipline.Exec(ctx)
	return err
}
