package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// CreatePostHandler 发布帖子
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Debug("c.ShouldBindJson(p) error", zap.Any("err", err))
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		// 并使用removeTopStruct函数去除字段名中的结构体名称标识
		zap.L().Debug("c.ShouldBindJson(p) error", zap.Any("err", err))
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// community是否存在
	_, err := logic.GetCommunityDetail(p.CommunityId)
	if err != nil {
		if errors.Is(err, mysql.ErrorInvalidID) {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 2. 创建帖子
	uid, err := getCurrentUser(c) // 获取uid为AuthorId做准备
	if err != nil {
		zap.L().Error("cannot get uid", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorId = uid
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
	return
}
