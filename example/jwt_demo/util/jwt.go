/**
 * @Author: vincent
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2021/8/23 18:43
 */

package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("aaabbbccceddf")
var issuser = "gin-blog"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) // 3小时过期

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			Issuer:    issuser,           // 签名的发行者
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
