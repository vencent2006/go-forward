package trade

import (
	"time"
)

type (
	PlaceOrder struct {
		ClOrdID string `json:"clOrdId"`
		Tag     string `json:"tag"`
		SMsg    string `json:"sMsg"`
		SCode   string `json:"sCode"`
		OrdID   string `json:"ordId"`
	}
	CancelOrder struct {
		OrdID   string `json:"ordId"`
		ClOrdID string `json:"clOrdId"`
		SMsg    string `json:"sMsg"`
		SCode   string `json:"sCode"`
	}
	AmendOrder struct {
		OrdID   string `json:"ordId"`
		ClOrdID string `json:"clOrdId"`
		ReqID   string `json:"reqId"`
		SMsg    string `json:"sMsg"`
		SCode   string `json:"sCode"`
	}
	ClosePosition struct {
		InstID  string `json:"instId"`
		PosSide string `json:"posSide"`
	}
	Order struct {
		InstID      string    `json:"instId"`
		Ccy         string    `json:"ccy"`
		OrdID       string    `json:"ordId"`
		ClOrdID     string    `json:"clOrdId"`
		TradeID     string    `json:"tradeId"`
		Tag         string    `json:"tag"`
		Category    string    `json:"category"`
		FeeCcy      string    `json:"feeCcy"`
		RebateCcy   string    `json:"rebateCcy"`
		Px          float64   `json:"px"`
		Sz          float64   `json:"sz"`
		Pnl         float64   `json:"pnl"`
		AccFillSz   float64   `json:"accFillSz"`
		FillPx      float64   `json:"fillPx"`
		FillSz      float64   `json:"fillSz"`
		FillTime    float64   `json:"fillTime"`
		AvgPx       float64   `json:"avgPx"`
		Lever       float64   `json:"lever"`
		TpTriggerPx float64   `json:"tpTriggerPx"`
		TpOrdPx     float64   `json:"tpOrdPx"`
		SlTriggerPx float64   `json:"slTriggerPx"`
		SlOrdPx     float64   `json:"slOrdPx"`
		Fee         float64   `json:"fee"`
		Rebate      float64   `json:"rebate"`
		State       string    `json:"state"`
		TdMode      string    `json:"tdMode"`
		PosSide     string    `json:"posSide"`
		Side        string    `json:"side"`
		OrdType     string    `json:"ordType"`
		InstType    string    `json:"instType"`
		TgtCcy      string    `json:"tgtCcy"`
		UTime       time.Time `json:"uTime"`
		CTime       time.Time `json:"cTime"`
	}
	TransactionDetail struct {
		InstID   string    `json:"instId"`
		OrdID    string    `json:"ordId"`
		TradeID  string    `json:"tradeId"`
		ClOrdID  string    `json:"clOrdId"`
		BillID   string    `json:"billId"`
		Tag      float64   `json:"tag"`
		FillPx   float64   `json:"fillPx"`
		FillSz   float64   `json:"fillSz"`
		FeeCcy   float64   `json:"feeCcy"`
		Fee      float64   `json:"fee"`
		InstType string    `json:"instType"`
		Side     string    `json:"side"`
		PosSide  string    `json:"posSide"`
		ExecType string    `json:"execType"`
		TS       time.Time `json:"ts"`
	}
	PlaceAlgoOrder struct {
		AlgoID string `json:"algoId"`
		SMsg   string `json:"sMsg"`
		SCode  int64  `json:"sCode"`
	}
	CancelAlgoOrder struct {
		AlgoID string `json:"algoId"`
		SMsg   string `json:"sMsg"`
		SCode  int64  `json:"sCode"`
	}
	AlgoOrder struct {
		InstID       string    `json:"instId"`
		Ccy          string    `json:"ccy"`
		OrdID        string    `json:"ordId"`
		AlgoID       string    `json:"algoId"`
		ClOrdID      string    `json:"clOrdId"`
		TradeID      string    `json:"tradeId"`
		Tag          string    `json:"tag"`
		Category     string    `json:"category"`
		FeeCcy       string    `json:"feeCcy"`
		RebateCcy    string    `json:"rebateCcy"`
		TimeInterval string    `json:"timeInterval"`
		Px           float64   `json:"px"`
		PxVar        float64   `json:"pxVar"`
		PxSpread     float64   `json:"pxSpread"`
		PxLimit      float64   `json:"pxLimit"`
		Sz           float64   `json:"sz"`
		SzLimit      float64   `json:"szLimit"`
		ActualSz     float64   `json:"actualSz"`
		ActualPx     float64   `json:"actualPx"`
		Pnl          float64   `json:"pnl"`
		AccFillSz    float64   `json:"accFillSz"`
		FillPx       float64   `json:"fillPx"`
		FillSz       float64   `json:"fillSz"`
		FillTime     float64   `json:"fillTime"`
		AvgPx        float64   `json:"avgPx"`
		Lever        float64   `json:"lever"`
		TpTriggerPx  float64   `json:"tpTriggerPx"`
		TpOrdPx      float64   `json:"tpOrdPx"`
		SlTriggerPx  float64   `json:"slTriggerPx"`
		SlOrdPx      float64   `json:"slOrdPx"`
		OrdPx        float64   `json:"ordPx"`
		Fee          float64   `json:"fee"`
		Rebate       float64   `json:"rebate"`
		State        string    `json:"state"`
		TdMode       string    `json:"tdMode"`
		ActualSide   string    `json:"actualSide"`
		PosSide      string    `json:"posSide"`
		Side         string    `json:"side"`
		OrdType      string    `json:"ordType"`
		InstType     string    `json:"instType"`
		TgtCcy       string    `json:"tgtCcy"`
		CTime        time.Time `json:"cTime"`
		TriggerTime  time.Time `json:"triggerTime"`
	}
)
