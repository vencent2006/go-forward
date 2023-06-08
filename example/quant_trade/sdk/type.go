package sdk

type Exchange interface {
	BuyLimit() error
	SellLimit() error
	GetOrder() error
}

type Ticker struct {
	Ts    int     `json:"ts"`
	Close float64 `json:"close"`
}
