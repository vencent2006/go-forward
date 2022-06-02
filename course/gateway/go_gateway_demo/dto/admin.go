/**
 * @Author: vincent
 * @Description:
 * @File:  admin
 * @Version: 1.0.0
 * @Date: 2022/5/31 21:51
 */

package dto

import (
	"go-examples/course/gateway/go_gateway_demo/public"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminInfoOutput struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"` //密码
}

func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
