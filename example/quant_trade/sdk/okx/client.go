package okx

import "example/quant_trade/sdk"

type Client struct {
}

func NewOkxClient() *Client {
	return &Client{}
}

func (c *Client) GetLatestClose() sdk.Ticker {
	// todo 补充逻辑
	return sdk.Ticker{}
}
