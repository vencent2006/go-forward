package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// PostVoteHandler 投票
func PostVoteHandler(c *gin.Context) {
	// 1. 获取参数和校验参数
	p := new(models.ParamVoteData)
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

	// 2. 处理逻辑
	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost failed",
			zap.Int64("userID", userID),
			zap.Any("ParamVoteData", p),
			zap.Error(err),
		)
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}
