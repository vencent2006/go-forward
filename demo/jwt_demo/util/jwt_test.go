/**
 * @Author: vincent
 * @Description:
 * @File:  jwt_test
 * @Version: 1.0.0
 * @Date: 2021/8/23 19:07
 */

package util

import (
	"fmt"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	token, _ := GenerateToken("xingye", "123456")
	fmt.Println("生成的token:", token)
	claim, err := ParseToken(token)
	if err != nil {
		t.Error("解析token出现错误：", err)
	} else if time.Now().Unix() > claim.ExpiresAt {
		t.Log("时间超时")
	} else {
		t.Log("username:", claim.Username)
		t.Log("password:", claim.Password)
	}

}
