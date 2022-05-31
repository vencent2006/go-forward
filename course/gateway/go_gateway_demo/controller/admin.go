/**
 * @Author: vincent
 * @Description:
 * @File:  admin_login
 * @Version: 1.0.0
 * @Date: 2022/5/31 00:42
 */

package controller

import (
	"encoding/json"
	"fmt"
	"go-examples/course/gateway/go_gateway_demo/dto"
	"go-examples/course/gateway/go_gateway_demo/middleware"
	"go-examples/course/gateway/go_gateway_demo/public"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func AdminRegister(group *gin.RouterGroup) {
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept json
// @Produce json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminLogin *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	// 1.读取sessionKey对应json 转换为结构体
	// 2.取出数据然后封装输出结构体

	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}
