package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

// InsertPost 向数据库中插入一条新的帖子记录
func InsertPost(post *models.Post) error {
	// 执行SQL语句入库
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}

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

func GetPostList(page int64, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
				post_id, title, content, author_id, community_id, status, create_time, update_time 
				from post
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
