package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocalEntReq struct {
	Uid      int64 `json:"uid" form:"uid"`
	CityId   int   `json:"city_id" form:"city_id" binding:"required"`
	PageNo   int   `json:"page_no" form:"page_no" binding:"required"`
	PageSize int   `json:"page_size" form:"page_size" binding:"required"`
}

type LocalEntRes struct {
	CityId      int      `json:"city_id"`
	List        []string `json:"list"`
	TotalNumber int      `json:"total_number"`
	PageNo      int      `json:"page_no"`
	PageSize    int      `json:"page_size"`
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	r.GET("/interface/local_ent", func(c *gin.Context) {
		req := &LocalEntReq{}
		resp := &response{}
		err := c.BindQuery(req)
		if err != nil {
			resp.Code = 1001
			resp.Msg = err.Error()
			resp.Data = nil
			c.JSON(http.StatusOK, resp)
		}

		//c.JSON(http.StatusOK, "ok")
		c.JSON(http.StatusBadGateway, "return 502")
	})

	r.Run(":8081")
}
