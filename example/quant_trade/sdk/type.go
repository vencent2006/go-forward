package sdk

import "time"

const (
	// side def
	SIDE_BUY  = "buy"
	SIDE_SELL = "sell"

	// status
	ORDER_STATUS_LIVE        = 0 // 等待成交
	ORDER_STATUS_PART_FILLED = 1 // 部分成交
	ORDER_STATUS_FILLED      = 2 // 完全成交
	ORDER_STATUS_CANCELED    = 3 // 撤单成功
)

// todo 加个interface，可以接收不同交易所
type Exchange interface {
	GetName() string
	GetTicker(*GetTickerReq) (*GetTickerRes, error)
	PlaceOrder(*PlaceOrderReq) (*PlaceOrderRes, error)
	GetOrderByClOrdId(*GetOrderByClOrdIdReq) (*GetOrderRes, error)
	GetOrderByOrdId(*GetOrderByOrdIdReq) (*GetOrderRes, error)
	CancelOrderByClOrdId(*CancelOrderByClOrdIdReq) (*CancelOrderRes, error)
	CancelOrderByOrdId(*CancelOrderByOrdIdReq) (*CancelOrderRes, error)
}

type GetTickerReq struct {
	InstId string `json:"instId"`
}

type GetTickerRes struct {
	Channel string    `json:"channel"`
	InstId  string    `json:"instId"`
	Ts      time.Time `json:"ts"`
	Close   float64   `json:"close"`
}

type PlaceOrderReq struct {
	Side    string
	InstId  string
	ClOrdId string
	Price   float64
	Size    float64
}

type PlaceOrderRes struct {
	Side    string
	ClOrdId string
	OrdId   string
}

type GetOrderByClOrdIdReq struct {
	InstId  string
	ClOrdId string
}

type GetOrderByOrdIdReq struct {
	InstId string
	OrdId  string
}

type GetOrderRes struct {
	InstId  string
	OrdId   string
	ClOrdId string
	Status  int
}

type CancelOrderByClOrdIdReq struct {
	InstId  string
	ClOrdId string
}

type CancelOrderByOrdIdReq struct {
	InstId string
	OrdId  string
}

type CancelOrderRes struct {
	ClOrdId string
	OrdId   string
}
