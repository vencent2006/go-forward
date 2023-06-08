package cache

type BuyStack struct {
	Strategy string
}

func NewBuyStack(strategy string) *BuyStack {
	return &BuyStack{Strategy: strategy}
}

func (b *BuyStack) Top() int {
	// b.
	// 用strategy区分不同的策略id
	// todo redis读取逻辑
	return 0
}

func (b *BuyStack) Push() error {
	// 用strategy区分不同的策略id
	// todo redis更改逻辑
	return nil
}

func (b *BuyStack) Pop() interface{} {
	// 用strategy区分不同的策略id
	// todo redis更改逻辑
	return nil
}
