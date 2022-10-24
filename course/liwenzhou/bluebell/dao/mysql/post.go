package mysql

import (
	"bluebell/models"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) error {
	// 执行SQL语句入库
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}

// GetPostDetailByID 根据ID查询帖子数据
func GetPostDetailByID(id int64) (p *models.Post, err error) {
	p = new(models.Post)
	sqlStr := `select 
				post_id, title, content, author_id, community_id, status, create_time, update_time 
				from post 
				where post_id = ?`
	err = db.Get(p, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("post not exist", zap.Int64("post_id", id))
			err = ErrorInvalidID
		}
		return
	}
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page int64, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
				post_id, title, content, author_id, community_id, status, create_time, update_time 
				from post
				ORDER BY create_time DESC
				limit ?,?`
	if err = db.Select(&postList, sqlStr, (page-1)*size, size); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no post in db")
			err = nil
			return
		}
	}
	return
}

// GetPostListByIDs 根据指定的ID列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	// FIND_IN_SET 保证顺序
	sqlStr := `select 
				post_id, title, content, author_id, community_id, status, create_time, update_time 
				from post 
				where post_id in (?)
				order by FIND_IN_SET(post_id, ?)
				`
	// 学习参考： https://www.liwenzhou.com/posts/Go/sqlx/
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		zap.L().Error("sqlx.In failed", zap.Error(err))
		return nil, err
	}
	query = db.Rebind(query)
	// postList要加&
	err = db.Select(&postList, query, args...) // args要加...
	return
}

// GetPostListByCommunityID 根据指定的CommunityID列表查询帖子列表
func GetPostListByCommunityID(ids []string) (postList []*models.Post, err error) {
	// FIND_IN_SET 保证顺序
	sqlStr := `select 
				post_id, title, content, author_id, community_id, status, create_time, update_time 
				from post 
				where community_id = ?
				order by FIND_IN_SET(post_id, ?)
				`
	// 学习参考： https://www.liwenzhou.com/posts/Go/sqlx/
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		zap.L().Error("sqlx.In failed", zap.Error(err))
		return nil, err
	}
	query = db.Rebind(query)
	// postList要加&
	err = db.Select(&postList, query, args...) // args要加...
	return
}
