package redis

import (
	"bluebell/models"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

func getIDsFromKey(key string, page, size int64) ([]string, error) {
	// 确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1
	// 3. ZRevRange 按分数从大到小（时间从近要远，得分从高到低）查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 从redis获取id
	// 1. 根据用户请求中携带的order参数确定要查询的redis key
	var key string

	if p.Order == models.OrderTime {
		key = getRedisKey(KeyPostTimeZSet)
	} else if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	} else {
		zap.L().Error("invalid order", zap.String("p.Order", p.Order))
		return nil, errors.New("invalid order")
	}

	return getIDsFromKey(key, p.Page, p.Size)
}

// GetCommunityPostIDsInOrder 按社区查询帖子的id列表
func GetCommunityPostIDsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	// ZInterStore 把区分的帖子set与帖子分数的ZSet 生成一个新的ZSet
	// 针对新的ZSet按之前的逻辑取数据

	// 社区的key
	communityKey := getRedisKey(KeyCommunitySetPrefix + strconv.FormatInt(p.CommunityID, 10))
	var orderKey string
	if p.Order == models.OrderTime {
		orderKey = getRedisKey(KeyPostTimeZSet)
	} else if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	} else {
		zap.L().Error("invalid order", zap.String("p.Order", p.Order))
		return nil, errors.New("invalid order")
	}

	// 利用缓存key减少ZInterStore执行的次数
	key := getRedisKey(p.Order + ":" + strconv.FormatInt(p.CommunityID, 10))
	zap.L().Debug("GetCommunityPostIDsInOrder ", zap.String("key", key))
	if client.Exists(key).Val() < 1 {
		// 不存在，需要计算; 这里使用pipeline
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX", // 交集应该只有一个数据，感觉是不是max，应该没啥影响
		}, communityKey, orderKey) // ZInterStore 计算
		pipeline.Expire(key, 60*time.Second) // 设置过期时间60秒过期
		_, err := pipeline.Exec()
		if err != nil {
			zap.L().Error("pipeline exec failed", zap.Error(err))
			return nil, err
		}
	}

	// 存在的话就直接跟进key查询ids
	return getIDsFromKey(key, p.Page, p.Size)
}
