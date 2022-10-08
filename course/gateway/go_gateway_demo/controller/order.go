package controller

import (
	"go-examples/course/gateway/go_gateway_demo/dto"
	"go-examples/course/gateway/go_gateway_demo/middleware"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func OrderRegister(group *gin.RouterGroup) {
	service := &OrderController{}

	group.GET("/order_stat", service.OrderStat)
}

// OrderStat godoc
// @Summary     订单统计
// @Description 订单统计
// @Tags        订单管理
// @ID          /order/order_stat
// @Accept      json
// @Produce     json
// @Param       stat_type  query    int                                           true "请求类型"
// @Param       start_date query    string                                        true "开始日期"
// @Param       end_date   query    string                                        true "结束日期"
// @Success     200        {object} middleware.Response{data=dto.OrderStatOutput} "success"
// @Router      /order/order_stat [get]
func (service *OrderController) OrderStat(c *gin.Context) {
	// curl http://127.0.0.1:8880/order/order_stat?stat_type=1&start_date=20201001&end_date=20201002
	// check input
	params := &dto.OrderStatInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// output
	middleware.ResponseSuccess(c, &dto.OrderStatOutput{
		StatType:    params.StatType,
		StartDate:   params.StartDate,
		EndDate:     params.EndDate,
		ItemNew:     1,
		ItemRenew:   2,
		ItemPackage: 3,
		ItemTotal:   6,
	})
}
