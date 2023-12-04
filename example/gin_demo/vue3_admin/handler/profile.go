package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleInfo struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type PermissionInfo struct {
	Menus []string `json:"menus"`
}

type Profile struct {
	Avatar     string         `json:"avatar"`
	Id         string         `json:"id"`
	Permission PermissionInfo `json:"permission"`
	Role       []RoleInfo     `json:"role"`
	Title      string         `json:"title"`
	Username   string         `json:"username"`
	Idx        string         `json:"_id"`
}

func HandleProfile(c *gin.Context) {
	r := &Response{
		Code: CODE_SUCC,
		Data: Profile{
			Avatar: "https://m.imooc.com/static/wap/static/common/img/logo-small@2x.png",
			Id:     "61270a9ec87aa543c9c3420",
			Permission: PermissionInfo{
				Menus: []string{"userManage", "roleList", "permissionList", "articleRanking", "articleCreate"},
			},
			Role: []RoleInfo{
				{
					Id:    "1",
					Title: "超级管理员",
				},
			},
			Title:    "超级管理员",
			Username: "super-admin",
			Idx:      "612ba1cfdeac5b28b43243e",
		},
		Message: "获取资料成功",
		Success: BOOL_SUCC,
	}
	c.JSON(http.StatusOK, r)
}
