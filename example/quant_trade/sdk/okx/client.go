package okx

import (
	"example/quant_trade/sdk"
)

type Client struct {
}

func NewOkxClient() *Client {
	return &Client{}
}

func (c *Client) GetLatestClose() (sdk.Ticker, error) {
	// todo 补充逻辑
	return sdk.Ticker{}, nil
}

func (c *Client) BuyLimit() (*PlaceOrderData, error) {
	//	if res.Code != "0" {
	//		return fmt.Errorf("code:%s, msg:%s", res.Code, res.Msg)
	//	} else if res.Data[0].SCode != "0" {
	//		return fmt.Errorf("scode:%s, smsg:%s", res.Data[0].SCode, res.Data[0].SMsg)
	//	}
	return nil, nil
}

func (c *Client) SellLimit() (*PlaceOrderData, error) {
	//	if res.Code != "0" {
	//		return fmt.Errorf("code:%s, msg:%s", res.Code, res.Msg)
	//	} else if res.Data[0].SCode != "0" {
	//		return fmt.Errorf("scode:%s, smsg:%s", res.Data[0].SCode, res.Data[0].SMsg)
	//	}
	return nil, nil
}

func (c *Client) GetOrderByClientOrderId(clientOrderId string) (*GetOrderData, error) {
	//	if res.Code != "0" {
	//		return fmt.Errorf("code:%s, msg:%s", res.Code, res.Msg)
	//	} else if res.Data[0].SCode != "0" {
	//		return fmt.Errorf("scode:%s, smsg:%s", res.Data[0].SCode, res.Data[0].SMsg)
	//	}
	return nil, nil
}

func (c *Client) CancelOrderByClientOrderId(clientOrderId string) (*CancelOrderData, error) {
	//	if res.Code != "0" {
	//		return fmt.Errorf("code:%s, msg:%s", res.Code, res.Msg)
	//	} else if res.Data[0].SCode != "0" {
	//		return fmt.Errorf("scode:%s, smsg:%s", res.Data[0].SCode, res.Data[0].SMsg)
	//	}
	return nil, nil
}
