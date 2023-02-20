package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/httplib"
)

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
		UserId:         123,
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

func Test_PrintBool(t *testing.T) {
	var a bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%t\n", a)
}
