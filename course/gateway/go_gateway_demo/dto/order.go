package dto

import (
	"go-examples/course/gateway/go_gateway_demo/public"

	"github.com/gin-gonic/gin"
)

type OrderStatInput struct {
	StatType  int    `json:"stat_type" form:"stat_type" comment:"统计类型" example:"1" validate:"required"`
	StartDate string `json:"start_date" form:"start_date" comment:"开始日期" example:"20201001" validate:"required"`
	EndDate   string `json:"end_date" form:"end_date" comment:"结束日期" example:"20201002" validate:"required"`
}

func (param *OrderStatInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type OrderStatOutput struct {
	StatType    int    `json:"stat_type" form:"stat_type" comment:"统计类型" example:"1" validate:""`
	StartDate   string `json:"start_date" form:"start_date" comment:"开始日期" example:"20201001" validate:""`
	EndDate     string `json:"end_date" form:"end_date" comment:"结束日期" example:"20201002" validate:""`
	ItemNew     int    `json:"item_new" form:"item_new" comment:"新客" example:"1" validate:""`
	ItemRenew   int    `json:"item_renew" form:"item_renew" comment:"续费" example:"2" validate:""`
	ItemPackage int    `json:"item_package" form:"item_package" comment:"增值服务" example:"3" validate:""`
	ItemTotal   int    `json:"item_total" form:"item_total" comment:"总和" example:"6" validate:""`
}
