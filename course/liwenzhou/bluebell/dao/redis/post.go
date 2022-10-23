package redis

import (
	"bluebell/models"
	"errors"

	"go.uber.org/zap"
)

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

	// 2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3. ZRevRange 按分数从大到小（时间从近要远，得分从高到低）查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
