package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

/*
投票的几种情况：
direction=1时：
	1. 之前没有投过票，现在投：赞成票 --> 更新分数和投票记录 差值的绝对值: 1 +432
	2. 之前投反对票，现在改投：赞成票 --> 更新分数和投票记录 差值的绝对值: 2 +432*2
direction=0时
	1. 之前投赞成票，现在要：取消投票 --> 更新分数和投票记录 差值的绝对值: 1 -432
	2. 之前投反对票，现在要：取消投票 --> 更新分数和投票记录 差值的绝对值: 1 +432
direction=-1时
	1. 之前没有投过票，现在投：反对票 --> 更新分数和投票记录 差值的绝对值: 1 -432
	2. 之前投赞成票，现在改投：反对票 --> 更新分数和投票记录 差值的绝对值: 2 -432*2

投票的限制：
	1. 每个帖子子发表之日起一个星期之内允许用户投票，超过一个兴趣就不允许再投票了
	2. 到期之后将redis中保存的赞成票数和反对票数存储到mysql表中
	3. 到期之后删除哪个 KeyPostVotedZSetPrefix
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600 // 一周的秒数
	scorePerVote     = 432           // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票的限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 2和3需要放到一个pipeline事务中操作
	// 2. 更新帖子分数
	// 先查当前用户给当前帖子的投票记录
	oldValue := client.ZScore(getRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	var op float64
	if value > oldValue {
		op = 1
	} else {
		op = -1
	}
	absDiff := math.Abs(oldValue - value) // 计算两次投票的差值

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*absDiff*scorePerVote, postID)

	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value, // 当前用户投的赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
