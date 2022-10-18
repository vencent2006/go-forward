package logic

import (
	"liwenzhou/bluebell/dao/mysql"
	"liwenzhou/bluebell/models"
	"liwenzhou/bluebell/pkg/jwt"
	"liwenzhou/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) error {
	// 1. 判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (string, error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是user指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 生成JWT
	return jwt.GenToken(user.UserID, user.Username)
}
