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
