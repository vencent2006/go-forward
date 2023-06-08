package cache

type LastIndex struct {
	Strategy string
}

func NewLastIndex(strategy string) *LastIndex {
	return &LastIndex{Strategy: strategy}
}

func (l *LastIndex) Get() int {
	// 用strategy区分不同的策略id
	// todo redis读取逻辑
	return 0
}

func (l *LastIndex) Set(newLastIndex int) error {
	// 用strategy区分不同的策略id
	// todo redis更改逻辑
	return nil
}
