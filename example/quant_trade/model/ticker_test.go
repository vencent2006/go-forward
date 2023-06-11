package model

import (
	"example/quant_trade/sdk"
	"testing"
	"time"
)

func TestTickerModel_GetLast(t *testing.T) {

	channel := "okx"
	res, err := NewTickerModel().GetLast(channel)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("GetLast | %+v", res)
}

func TestTickerModel_Insert(t *testing.T) {

	req := sdk.GetTickerRes{
		Channel: "okx",
		InstId:  "BTC-USDT",
		Ts:      time.Now().UTC(),
		Close:   111.333,
	}
	err := NewTickerModel().Insert(&req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Insert req succ| %+v", req)
}
