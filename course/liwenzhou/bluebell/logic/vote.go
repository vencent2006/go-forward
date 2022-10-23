package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

// 推荐阅读
// 基于用户投票的相关算法: http://www.ruanyifeng/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分 86400/200 -> 200张赞成票可以给你的帖子续一天 -> 《redis实战》

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(
		strconv.FormatInt(userID, 10),
		strconv.FormatInt(p.PostID, 10),
		float64(p.Direction))
}
