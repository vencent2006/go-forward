package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"liwenzhou/bluebell/models"
)

// 把每一步数据库操作封装成函数

const secret = "liwenzhou.com"

// CheckUserExist 检查指定用户是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) error {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err := db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

func encryptPassword(originPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(originPassword)))
}
