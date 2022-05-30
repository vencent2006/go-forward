/**
 * @Author: vincent
 * @Description:
 * @File:  admin_login
 * @Version: 1.0.0
 * @Date: 2022/5/31 00:42
 */

package controller

import (
	"go-examples/course/gateway/go_gateway_demo/dto"
	"go-examples/course/gateway/go_gateway_demo/middleware"

	"github.com/gin-gonic/gin"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
