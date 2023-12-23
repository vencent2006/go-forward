package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TenantInfo struct {
	TenantId    string
	ServiceName string
}

type GetTenantListRes struct {
	TenantList []TenantInfo `json:"tenant_list"`
}

func GetTenantList(c *gin.Context) {
	r := &Response{
		Code: 100000,
		Data: GetTenantListRes{
			TenantList: []TenantInfo{
				{
					TenantId:    "123",
					ServiceName: "service_123",
				},
				{
					TenantId:    "456",
					ServiceName: "service_456",
				},
			}},
		Message: "succ",
		Success: BOOL_SUCC,
	}
	c.JSON(http.StatusOK, r)
}
