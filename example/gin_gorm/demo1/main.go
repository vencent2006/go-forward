package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BusinessToolOrder struct {
	Id         int64     `json:"id"`
	Uid        int64     `json:"uid"`
	Pid        int       `json:"pid"`
	Oid        int64     `json:"oid"`
	Status     int8      `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	ValidTime  time.Time `json:"valid_time"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	router := gin.Default()
	router.GET("/order", getOrder)
	router.Run(":8081")
}

func getOrder(c *gin.Context) {
	// 不是oid，是business_tool_order的id
	orderId := c.Query("id")
	if orderId == "" {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "参数错误",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
}
