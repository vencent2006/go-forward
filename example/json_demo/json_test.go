package json_demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

type NotifyChargeInfo struct {
	RecordId int `json:"recordId"`
	//RecordType     string `json:"recordType"`
	ToAddr         string `json:"toAddr"`
	TxHash         string `json:"txHash"`
	Time           int64  `json:"time"`
	Confirmations  int64  `json:"confirmations"`
	Value          string `json:"value"`
	ChainName      string `json:"chainName"`
	BlockNumber    int64  `json:"blockNumber"`
	ContactAddr    string `json:"contactAddr"` // recordType为coin时，该字段为空
	TokenSymbol    string `json:"tokenSymbol"` // recordType为coin时，该字段为空
	TokenValue     string `json:"tokenValue"`  // recordType为coin时，该字段为空
	UserId         int    `json:"userId"`
	AssetAccountId int    `json:"assetAccountId"`
	AssetId        int    `json:"assetId"`
	AssetTokenId   int    `json:"assetTokenId"`
}

func TestJsonOmitEmpty(t *testing.T) {
	info := &NotifyChargeInfo{
		RecordId: 1084,
		//RecordType:     "coin",
		ToAddr:         "0x02c63c10cf1b67c5f8f1e5edf14e1795292d892f",
		TxHash:         "0xe98422c89cbb50da3025f56a2d3a0019b989c7f5288749b9aee85437ddb02501",
		Time:           1676431020,
		Confirmations:  1,
		Value:          "0.0001",
		ChainName:      "ethereum",
		BlockNumber:    8493798,
		ContactAddr:    "",
		TokenSymbol:    "",
		TokenValue:     "",
		UserId:         1000000004,
		AssetAccountId: 1,
		AssetId:        1,
		AssetTokenId:   60,
	}
	bytes, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", bytes)
}
