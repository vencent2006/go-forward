package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	// 1. 生成post_id
	p.PostId = snowflake.GenID()

	// 3. 保存进数据库
	if err := mysql.CreatePost(p); err != nil {
		return err
	}
	if err := redis.CreatePost(p.PostId, p.CommunityId); err != nil {
		return err
	}
	return nil
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

func GetPostListWithoutCommunityID(p *models.ParamPostList) (postList []*models.ApiPostDetail, err error) {
	// 2.去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder(p) failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 data")
		return
	}
	// 3.根据id去数据库查询帖子详细信息
	// 返回的数据还要按照给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids)",
			zap.Error(err), zap.Any("ids", ids))
		return
	}

	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
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
		detail := new(models.ApiPostDetail)
		detail.AuthorName = authorDetail.Username
		detail.CommunityDetail = communityDetail
		detail.Post = post

		postList = append(postList, detail)
	}

	return
}

// GetPostListWithCommunityID 跟进Community来获取帖子列表
func GetPostListWithCommunityID(p *models.ParamPostList) (postList []*models.ApiPostDetail, err error) {
	zap.L().Debug("models.ParamCommunityPostList is ", zap.Any("ParamCommunityPostList", p))
	// 2.去redis查询id列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetCommunityPostIDsInOrder(p) failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetCommunityPostIDsInOrder return 0 data",
			zap.Any("models.ParamCommunityPostList", p))
		return
	}
	// 3.根据id去数据库查询帖子详细信息
	// 返回的数据还要按照给定的id的顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids)",
			zap.Error(err), zap.Any("ids", ids))
		return
	}

	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
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
		detail := new(models.ApiPostDetail)
		detail.AuthorName = authorDetail.Username
		detail.CommunityDetail = communityDetail
		detail.Post = post

		postList = append(postList, detail)
	}

	return
}

func GetPostListNew(p *models.ParamPostList) (postList []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		// 无配置社区id
		return GetPostListWithoutCommunityID(p)
	} else {
		// 有配置社区id
		return GetPostListWithCommunityID(p)
	}
}
