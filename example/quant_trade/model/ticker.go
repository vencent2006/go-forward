package model

import (
	"example/quant_trade/sdk"
	"time"
)

type TickerModel struct {
	TickerId    int       `gorm:"column:ticker_id"`
	Channel     string    `gorm:"column:channel"`
	InstId      string    `gorm:"column:instId"`
	Price       float64   `gorm:"column:price"`
	ChannelTime time.Time `gorm:"column:channel_time"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

func NewTickerModel() *TickerModel {
	return &TickerModel{}
}

func (t *TickerModel) TableName() string {
	return "ticker"
}

func (t *TickerModel) Insert(ticker *sdk.GetTickerRes) error {
	m := TickerModel{
		Channel:     ticker.Channel,
		InstId:      ticker.InstId,
		Price:       ticker.Close,
		ChannelTime: ticker.Ts,
		CreateTime:  time.Now().UTC(),
		UpdateTime:  time.Now().UTC(),
	}
	return db.Create(&m).Error
}

func (t *TickerModel) GetLast(channel string) (*TickerModel, error) {
	m := TickerModel{}
	res := db.Where("channel = ?", channel).Last(&m)
	return &m, res.Error
}
