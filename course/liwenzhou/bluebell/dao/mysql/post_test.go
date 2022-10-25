package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
	"time"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "root",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := &models.Post{
		PostId:      11,
		AuthorId:    123,
		CommunityId: 1,
		Title:       "test_10",
		Content:     "test_10",
		CreateTime:  time.Time{},
	}
	err := CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost insert record into mysql failed, err:%v\n", err)
	}
	t.Logf("CreatePost insert record into mysql success")
}
