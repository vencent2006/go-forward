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
	"go-examples/course/gateway/go_gateway_demo/dao"
	"go-examples/course/gateway/go_gateway_demo/dto"
	"go-examples/course/gateway/go_gateway_demo/middleware"
	"go-examples/course/gateway/go_gateway_demo/public"
	"time"

	"github.com/gin-gonic/contrib/sessions"

	"go-examples/course/gateway/go_gateway_demo/golang_common/lib"

	"github.com/gin-gonic/gin"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLogout)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept json
// @Produce json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	// step 1: params.Username 取得管理员信息 admininfo
	// step 2: adminInfo.salt + params.Password sha256 => saltPassword
	// step 3: saltPassword == adminInfo.password
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.UserName}

	middleware.ResponseSuccess(c, out)
}

//AdminLogout godoc
//@Summary 管理员退出
//@Description 管理员退出
//@Tags 管理员接口
//@ID /admin_login/logout
//@Accept json
//@Produce json
//@Success 200 {object} middleware.Response{data=string} "success"
//@Router /admin_login/logout [get]
func (adminLogin *AdminLoginController) AdminLogout(c *gin.Context) {

	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()

	middleware.ResponseSuccess(c, "")
}
