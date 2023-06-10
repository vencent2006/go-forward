package okx

const (
	ORDER_STATUS_CANCELED    = "canceled"         // 撤单成功
	ORDER_STATUS_LIVE        = "live"             // 等待成交
	ORDER_STATUS_PART_FILLED = "partially_filled" // 部分成交
	ORDER_STATUS_FILLED      = "filled"           // 完全成交

)

// place order
type PlaceOrderRes struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data []PlaceOrderData `json:"data"`
}

type PlaceOrderData struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

// get order
type GetOrderRes struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []GetOrderData `json:"data"`
}

type GetOrderData struct {
	InstType           string `json:"instType"`
	InstId             string `json:"instId"`
	Ccy                string `json:"ccy"`
	OrdId              string `json:"ordId"`
	ClOrdId            string `json:"clOrdId"`
	Tag                string `json:"tag"`
	Px                 string `json:"px"`
	Sz                 string `json:"sz"`
	Pnl                string `json:"pnl"`
	OrdType            string `json:"ordType"`
	Side               string `json:"side"`
	PosSide            string `json:"posSide"`
	TdMode             string `json:"tdMode"`
	AccFillSz          string `json:"accFillSz"`
	FillPx             string `json:"fillPx"`
	TradeId            string `json:"tradeId"`
	FillSz             string `json:"fillSz"`
	FillTime           string `json:"fillTime"`
	Source             string `json:"source"`
	State              string `json:"state"`
	AvgPx              string `json:"avgPx"`
	Lever              string `json:"lever"`
	TpTriggerPx        string `json:"tpTriggerPx"`
	TpTriggerPxType    string `json:"tpTriggerPxType"`
	TpOrdPx            string `json:"tpOrdPx"`
	SlTriggerPx        string `json:"slTriggerPx"`
	SlTriggerPxType    string `json:"slTriggerPxType"`
	SlOrdPx            string `json:"slOrdPx"`
	StpId              string `json:"stpId"`
	StpMode            string `json:"stpMode"`
	FeeCcy             string `json:"feeCcy"`
	Fee                string `json:"fee"`
	RebateCcy          string `json:"rebateCcy"`
	Rebate             string `json:"rebate"`
	TgtCcy             string `json:"tgtCcy"`
	Category           string `json:"category"`
	ReduceOnly         string `json:"reduceOnly"`
	CancelSource       string `json:"cancelSource"`
	CancelSourceReason string `json:"cancelSourceReason"`
	QuickMgnType       string `json:"quickMgnType"`
	AlgoClOrdId        string `json:"algoClOrdId"`
	AlgoId             string `json:"algoId"`
	UTime              string `json:"uTime"`
	CTime              string `json:"cTime"`
}

// cancel order
type CancelOrderRes struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Data []CancelOrderData `json:"data"`
}

type CancelOrderData struct {
	ClOrdId string `json:"clOrdId"`
	OrdId   string `json:"ordId"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}
