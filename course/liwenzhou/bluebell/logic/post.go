package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	// 1. 生成post_id
	p.PostId = snowflake.GenID()

	// 3. 保存进数据库
	return mysql.InsertPost(p)
}

func GetPostDetailByID(id int64) (p *models.Post, err error) {
	return mysql.GetPostDetailByID(id)
}
