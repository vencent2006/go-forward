package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"strconv"

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

// PostDetailHandler 获取帖子详情
func PostDetailHandler(c *gin.Context) {
	// 1. 获取参数和校验参数
	idStr := c.Param("pid")
	postId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Warn("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 逻辑处理
	post, err := logic.GetPostDetailByID(postId)
	if err != nil {
		if err == mysql.ErrorInvalidID {
			zap.L().Warn("get post detail with invalid param", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		zap.L().Error("logic.GetPostDetailByID(postId) failed", zap.Int64("postId", postId), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回影响
	ResponseSuccess(c, post)
}

func GetPostListSimpleHandler(c *gin.Context) {
	// 1. 获取参数和检查参数
	page, size := getPageInfo(c)

	// 2. 获取数据
	postList, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList()", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, postList)
}

// GetPostListHandler 升级版帖子列表接口
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照 分数排序
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息
func GetPostListHandler(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 1. 获取参数和检查参数
	p := &models.ParamPostList{
		Page:        1,
		Size:        10,
		Order:       models.OrderTime,
		CommunityID: 0,
	}
	// c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(&p); err != nil {
		zap.L().Error("PostListHandler2 with invalid params", zap.Error(err), zap.Any("ParamPostList", p))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 获取数据
	postList, err := logic.GetPostListNew(p)
	if err != nil {
		zap.L().Error("logic.GetPostListNew", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, postList)
}

//// GetCommunityPostListHandler 根据社区去查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
//	// 1. 获取参数和检查参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//		CommunityID: 0,
//	}
//	// c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
//	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
//	if err := c.ShouldBindQuery(&p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err), zap.Any("ParamPostList", p))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 2. 获取数据
//	postList, err := logic.GetPostList3(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList()", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//
//	// 3. 返回响应
//	ResponseSuccess(c, postList)
//}
