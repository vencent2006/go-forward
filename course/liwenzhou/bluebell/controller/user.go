package controller

import (
	"liwenzhou/bluebell/logic"
	"liwenzhou/bluebell/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		// 并使用removeTopStruct函数去除字段名中的结构体名称标识
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	return
}
