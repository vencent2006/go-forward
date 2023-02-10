package main

import (
	"testing"
	"time"

	"github.com/beego/beego/v2/client/httplib"
)

type NotifyChargeRes struct {
	Message string `json:"message"`
}

func post(url string, info *NotifyChargeInfo) (*NotifyChargeRes, error) {
	var err error

	timeout := time.Second
	req := httplib.Post(url).SetTimeout(timeout, timeout)
	_, err = req.JSONBody(info)
	if err != nil {
		return nil, err
	}
	res := &NotifyChargeRes{}
	err = req.ToJSON(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TestPost(t *testing.T) {

	notify := &NotifyChargeInfo{
		UserId:         "123",
		AssetAccountId: 1,
		AssetId:        2,
		AssetTokenId:   3,
	}

	url := "http://localhost:8080/post"
	res, err := post(url, notify)
	if err != nil {
		t.Errorf("PostNotifyChargeToPay failed, err = %s", err.Error())
		return
	}

	t.Logf("res is %+v", res)
}
