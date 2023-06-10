package okx

import (
	"example/quant_trade/sdk"
	"example/quant_trade/sdk/okx/rest"
	"example/quant_trade/sdk/okx/rest/request/market"
	"example/quant_trade/sdk/okx/rest/request/trade"
	"fmt"
)

type Client struct {
	restClient *rest.ClientRest
}

func NewOkxClient() *Client {
	return &Client{restClient: rest.NewClient(apiKey, secretKey, passphrase, restURL, isSimulated)}
}

func (c *Client) GetName() string {
	return "okx"
}

func (c *Client) GetTicker(r *sdk.GetTickerReq) (*sdk.GetTickerRes, error) {
	// todo 补充逻辑
	var req market.GetTicker
	req.InstId = r.InstId
	res, err := c.restClient.Market.GetTicker(req)
	if err != nil {
		return nil, err
	}
	if len(res.Tickers) != 1 {
		return nil, fmt.Errorf("len(res.Tickers) = %d, not 1", len(res.Tickers))
	}
	resp := &sdk.GetTickerRes{
		InstId: r.InstId,
		Ts:     res.Tickers[0].TS,
		Close:  res.Tickers[0].Last,
	}
	return resp, nil
}

func (c *Client) PlaceOrder(r *sdk.PlaceOrderReq) (*sdk.PlaceOrderRes, error) {
	if r.Side != sdk.SIDE_BUY && r.Side != sdk.SIDE_SELL {
		panic(fmt.Sprintf("invalid side %s", r.Side))
	}

	// https://www.okx.com/docs-v5/zh/#rest-api-trade-place-order
	var req trade.PlaceOrder
	req.InstID = r.InstId
	req.TdMode = "cash"
	req.ClOrdID = r.ClOrdId
	req.Side = r.Side
	req.OrdType = "limit"
	req.Px = r.Price
	req.Sz = r.Size
	res, err := c.restClient.Trade.PlaceOrder(req)
	if err != nil {
		return nil, err
	}
	if len(res.PlaceOrders) != 1 {
		return nil, fmt.Errorf("len(res.PlaceOrders) = %d, not 1", len(res.PlaceOrders))
	}

	if res.PlaceOrders[0].SCode != "0" {
		return nil, fmt.Errorf("scode:%s, smsg:%s", res.PlaceOrders[0].SCode, res.PlaceOrders[0].SMsg)
	}
	if r.ClOrdId != res.PlaceOrders[0].ClOrdID {
		panic(fmt.Errorf("r.ClOrderId(%s) != res.PlaceOrders[0].ClOrdID(%s)", r.ClOrdId, res.PlaceOrders[0].ClOrdID))
	}
	resp := &sdk.PlaceOrderRes{
		Side:    r.Side,
		ClOrdId: res.PlaceOrders[0].ClOrdID,
		OrdId:   res.PlaceOrders[0].OrdID,
	}
	return resp, nil
}

func (c *Client) GetOrderByClOrdId(r *sdk.GetOrderByClOrdIdReq) (*sdk.GetOrderRes, error) {

	var req trade.OrderDetails
	req.InstID = r.InstId
	req.ClOrdID = r.ClOrdId
	res, err := c.restClient.Trade.GetOrderDetail(req)
	if err != nil {
		return nil, err
	}
	if len(res.Orders) != 1 {
		return nil, fmt.Errorf("len(res.PlaceOrders) = %d, not 1", len(res.Orders))
	}
	resp := &sdk.GetOrderRes{
		InstId:  res.Orders[0].InstID,
		OrdId:   res.Orders[0].OrdID,
		ClOrdId: res.Orders[0].ClOrdID,
		Status:  res.Orders[0].State,
	}
	return resp, nil
}

func (c *Client) GetOrderByOrdId(r *sdk.GetOrderByOrdIdReq) (*sdk.GetOrderRes, error) {

	var req trade.OrderDetails
	req.InstID = r.InstId
	req.ClOrdID = r.OrdId
	res, err := c.restClient.Trade.GetOrderDetail(req)
	if err != nil {
		return nil, err
	}
	if len(res.Orders) != 1 {
		return nil, fmt.Errorf("len(res.PlaceOrders) = %d, not 1", len(res.Orders))
	}
	resp := &sdk.GetOrderRes{
		InstId:  res.Orders[0].InstID,
		OrdId:   res.Orders[0].OrdID,
		ClOrdId: res.Orders[0].ClOrdID,
		Status:  res.Orders[0].State,
	}
	return resp, nil
}

func (c *Client) CancelOrderByClOrdId(r *sdk.CancelOrderByClOrdIdReq) (*sdk.CancelOrderRes, error) {

	var req trade.CancelOrder
	req.InstID = r.InstId
	req.ClOrdID = r.ClOrdId
	res, err := c.restClient.Trade.CandleOrder(req)
	if err != nil {
		return nil, err
	}
	if len(res.CancelOrders) != 1 {
		return nil, fmt.Errorf("len(res.CancelOrders) = %d, not 1", len(res.CancelOrders))
	}
	if res.CancelOrders[0].SCode != "0" {
		return nil, fmt.Errorf("scode:%s, smsg:%s", res.CancelOrders[0].SCode, res.CancelOrders[0].SMsg)
	}
	if r.ClOrdId != res.CancelOrders[0].ClOrdID {
		panic(fmt.Errorf("CancelOrderByClOrdId | r.ClOrderId(%s) != res.CancelOrders[0].ClOrdID(%s)",
			r.ClOrdId, res.CancelOrders[0].ClOrdID))
	}

	resp := &sdk.CancelOrderRes{
		ClOrdId: res.CancelOrders[0].ClOrdID,
		OrdId:   res.CancelOrders[0].OrdID,
	}
	return resp, nil
}

func (c *Client) CancelOrderByOrdId(r *sdk.CancelOrderByOrdIdReq) (*sdk.CancelOrderRes, error) {

	var req trade.CancelOrder
	req.InstID = r.InstId
	req.ClOrdID = r.OrdId
	res, err := c.restClient.Trade.CandleOrder(req)
	if err != nil {
		return nil, err
	}
	if len(res.CancelOrders) != 1 {
		return nil, fmt.Errorf("len(res.CancelOrders) = %d, not 1", len(res.CancelOrders))
	}
	if res.CancelOrders[0].SCode != "0" {
		return nil, fmt.Errorf("scode:%s, smsg:%s", res.CancelOrders[0].SCode, res.CancelOrders[0].SMsg)
	}
	if r.OrdId != res.CancelOrders[0].OrdID {
		panic(fmt.Errorf("CancelOrderByOrdId | r.ClOrderId(%s) != res.CancelOrders[0].OrdID(%s)",
			r.OrdId, res.CancelOrders[0].OrdID))
	}

	resp := &sdk.CancelOrderRes{
		ClOrdId: res.CancelOrders[0].ClOrdID,
		OrdId:   res.CancelOrders[0].OrdID,
	}
	return resp, nil
}
