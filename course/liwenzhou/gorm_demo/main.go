package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Post 为内存对齐对字段顺序进行了调整
type Post struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	PostId      int64     `json:"post_id" gorm:"column:post_id"`
	AuthorId    int64     `json:"author_id" gorm:"column:author_id"`
	CommunityId int64     `json:"community_id" gorm:"column:community_id"`
	Status      int32     `json:"status" gorm:"column:status"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content" gorm:"column:content"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
	UpdateTime  time.Time `json:"update_time" gorm:"column:update_time"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/liwenzhou?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 查询
	var u = new(Post)
	db.First(u)
	fmt.Printf("%#v\n", u)
}
