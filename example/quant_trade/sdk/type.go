package sdk

import "time"

const (
	// side def
	SIDE_BUY  = "buy"
	SIDE_SELL = "sell"

	// status
	ORDER_STATUS_CANCELED    = "canceled"         // 撤单成功
	ORDER_STATUS_LIVE        = "live"             // 等待成交
	ORDER_STATUS_PART_FILLED = "partially_filled" // 部分成交
	ORDER_STATUS_FILLED      = "filled"           // 完全成交
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
	InstId string    `json:"instId"`
	Ts     time.Time `json:"ts"`
	Close  float64   `json:"close"`
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
	Status  string
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
