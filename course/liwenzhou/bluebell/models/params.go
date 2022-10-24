package models

// 定义请求的参数结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    int64 `json:"post_id, string" binding:"required"`       // 帖子id
	Direction int   `json:"direction, string" binding:"oneof=1 0 -1"` // 1:赞成，-1:反对, 0取消投票
}

// ParamPostList 获取帖子列表的query string参数
type ParamPostList struct {
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	Order       string `form:"order"`
	CommunityID int64  `form:"community_id"`
}

// ParamCommunityPostList 按照社区获取帖子列表的query string参数
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `form:"community_id"`
}
