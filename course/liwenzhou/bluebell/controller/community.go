package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ------ 社区相关 ------

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端的错误暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 获取社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	idStr := c.Param("cid") // 获取URL参数
	cid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询指定社区详情
	data, err := logic.GetCommunityDetail(cid)
	if err != nil {
		zap.L().Warn("logic.GetCommunityDetail() failed", zap.Error(err), zap.Int64("community_id", cid))
		if errors.Is(err, mysql.ErrorInvalidID) {
			ResponseError(c, CodeCommunityNotExist)
			return
		}
		ResponseError(c, CodeServerBusy) // 不轻易把服务端的错误暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
