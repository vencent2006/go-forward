package main

import (
	"go-examples/example/chat_demo_im/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyChargeInfo struct {
	RecordId       int    `json:"record_id"`
	ToAddr         string `json:"to_addr"`
	TxHash         string `json:"txHash"`
	Time           int64  `json:"time"`
	Confirmations  int64  `json:"confirmations"`
	Value          string `json:"value"`
	ChainName      string `json:"chainName"`
	BlockNumber    int64  `json:"block_number"`
	ContactAddr    string `json:"contactAddr"`
	TokenSymbol    string `json:"tokenSymbol"`
	TokenValue     string `json:"tokenValue"`
	UserId         int    `json:"userId"`
	AssetAccountId int    `json:"assetAccountId"`
	AssetId        int    `json:"assetId"`
	AssetTokenId   int    `json:"assetTokenId"`
}

type NotifyChargeRes struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	r := gin.Default()

	response := &NotifyChargeRes{
		Code:    10000,
		Message: "succ",
	}
	r.POST("/wallet/notify/chain/charge", func(c *gin.Context) {
		//body, _ := ioutil.ReadAll(c.Request.Body)
		//if body != nil {
		//	logger.Printf("body is %+v", string(body))
		//}
		info := &NotifyChargeInfo{}
		err := c.BindJSON(info)
		if err != nil {
			logger.Errorf("BindJSON err:%v", err)
			c.String(http.StatusBadRequest, "")
			return
		}

		logger.Infof("req is %+v", info)
		c.JSON(http.StatusOK, response)
	})

	r.Run(":8080")
}
