package model

import (
	"example/quant_trade/sdk"
	"fmt"
	"time"
)

type OrderModel struct {
	Id         string    `gorm:"column:id"`
	Channel    string    `gorm:"channel"`
	StrategyId int       `gorm:"column:strategy_id"`
	InstId     string    `gorm:"column:instId"`
	OrdId      string    `gorm:"column:ord_id"`
	ClOrdId    string    `gorm:"column:cl_ord_id"`
	price      float64   `gorm:"column:price"`
	size       float64   `gorm:"column:size"`
	Status     int       `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func NewOrderModel() *OrderModel {
	return &OrderModel{}
}

func (o *OrderModel) Insert(strategy int, orderRes *sdk.GetOrderRes) error {
	orderModel := OrderModel{
		StrategyId: strategy,
		InstId:     orderRes.InstId,
		OrdId:      orderRes.OrdId,
		ClOrdId:    orderRes.ClOrdId,
		Status:     orderRes.Status,
		CreateTime: time.Now().UTC(),
		UpdateTime: time.Now().UTC(),
	}
	// lastInsertId, 在orderModel里
	return db.Create(&orderModel).Error
}

func (o *OrderModel) GetByClOrdId(clOrdId string) (*OrderModel, error) {
	model := &OrderModel{}
	err := db.Where(&OrderModel{ClOrdId: clOrdId}).First(model).Error
	return model, err
}

func (o *OrderModel) GetByOrdId(ordId string) (*OrderModel, error) {
	model := &OrderModel{}
	err := db.Where(&OrderModel{OrdId: ordId}).First(model).Error
	return model, err
}

func (o *OrderModel) UpdateStatusByClOrdId(clOrdId string, oldStatus, NewStatus int) error {
	res := db.Model(&OrderModel{}).Where("cl_ord_id = ? and status = ?", clOrdId, oldStatus).Update("status", NewStatus)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		panic(fmt.Errorf("res.RowsAffected(%d) != 1", res.RowsAffected))
	}
	return nil
}

func (o *OrderModel) UpdateStatusByOrdId(ordId string, oldStatus, NewStatus int) error {
	res := db.Model(&OrderModel{}).Where("ord_id = ? and status = ?", ordId, oldStatus).Update("status", NewStatus)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		panic(fmt.Errorf("res.RowsAffected(%d) != 1", res.RowsAffected))
	}
	return nil
}
