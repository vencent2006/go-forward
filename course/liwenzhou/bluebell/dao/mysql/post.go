package mysql

import "bluebell/models"

// InsertPost 向数据库中插入一条新的帖子记录
func InsertPost(post *models.Post) error {
	// 执行SQL语句入库
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?,?,?,?,?)`
	_, err := db.Exec(sqlStr, post.PostId, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}
