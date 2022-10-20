package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 1. 生成post_id
	p.PostId = snowflake.GenID()

	// 3. 保存进数据库
	return mysql.InsertPost(p)
}

func GetPostDetailByID(id int64) (data *models.ApiPostDetail, err error) {

	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(id) failed", zap.Int64("post_id", id), zap.Error(err))
		return
	}

	// 根据作者id查询作者信息
	authorDetail, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("author_id",
			post.AuthorId), zap.Error(err))
		return nil, err
	}

	// 根据社区id查询社区信息
	communityDetail, err := mysql.GetCommunityDetailByID(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityId) failed",
			zap.Int64("community_id", post.CommunityId), zap.Error(err))
		return nil, err
	}

	// 填充ApiPostDetail
	data = new(models.ApiPostDetail)
	data.AuthorName = authorDetail.Username
	data.CommunityDetail = communityDetail
	data.Post = post

	return
}

func GetPostList(page int64, size int64) (postList []*models.ApiPostDetail, err error) {
	// 请求dao
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList failed", zap.Error(err))
		return nil, err
	}

	postList = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		detail, err := GetPostDetailByID(post.PostId)
		if err != nil {
			zap.L().Error("GetPostDetailByID(post.PostId) failed", zap.Int64("post_id", post.PostId), zap.Error(err))
			continue
		}

		postList = append(postList, detail)
	}

	return
}
