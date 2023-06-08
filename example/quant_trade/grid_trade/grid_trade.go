package grid_trade

import (
	"example/quant_trade/cache"
	"example/quant_trade/sdk"
	"example/quant_trade/sdk/okx"
	"fmt"
	"time"
)

const (
	SLEEP_SECONDS = 10 * time.Second
	MIN_GRID_NUM  = 2 // 最小网格数量
)

type StackItem struct {
	ClientOrderId string
	LastIndex     int
}

type GridTrade struct {
	Strategy string
	InstId   string  // 交易对, 如"BTC-USDT"
	Highest  float64 // 最高价
	Lowest   float64 // 最低价
	GridNum  int     // 网格数量
	Budget   float64 // 预算

	PercentLevels []float64 // 持仓比例区间
	PriceLevels   []float64 // 价格区间

	client         *okx.Client
	lastIndexCache *cache.LastIndex
	buyStackCache  *cache.BuyStack
}

func CalPercentLevels(gridNum int) []float64 {
	if gridNum <= 0 {
		panic(fmt.Sprintf("invalid gridNum(%d)", gridNum))
	}
	total := 1.0
	grid := total / float64(gridNum)
	var levels []float64
	for i := 0; i < gridNum; i++ {
		levels = append(levels, float64(i)*grid)
	}
	levels = append(levels, 1.0)
	return levels
}

func CalPriceLevels(highest float64, lowest float64, gridNum int) []float64 {
	if gridNum <= 0 {
		panic(fmt.Sprintf("invalid gridNum(%d)", gridNum))
	}
	if lowest < 0 || highest <= lowest {
		panic(fmt.Sprintf("invalid highest(%.2f) lowest(%.2f)", highest, lowest))
	}
	total := highest - lowest
	grid := total / float64(gridNum)
	var levels []float64
	for i := 0; i < gridNum; i++ {
		levels = append(levels, highest-float64(i)*grid)
	}
	levels = append(levels, lowest)
	return levels
}

func CalLastIndex() int {
	return 0
}

func NewGridTrade(strategy string, instId string, highest float64, lowest float64, gridNum int, budget float64) *GridTrade {
	if gridNum < MIN_GRID_NUM {
		panic(fmt.Sprintf("gridNum(%d) is invalid", gridNum))
	}
	// todo 重启后budget怎么算呢
	g := &GridTrade{
		Strategy:      strategy,
		InstId:        instId,
		Highest:       highest,
		Lowest:        lowest,
		GridNum:       gridNum,
		Budget:        budget,
		PriceLevels:   CalPriceLevels(highest, lowest, gridNum),
		PercentLevels: CalPercentLevels(gridNum),
	}
	g.client = okx.NewOkxClient()
	g.lastIndexCache = cache.NewLastIndex(strategy)
	g.buyStackCache = cache.NewBuyStack(strategy)

	return g
}

func (g *GridTrade) GetTickerWithSave() sdk.Ticker {
	// get close
	ticker := g.client.GetLatestClose()

	// todo save close
	return ticker
}

func (g *GridTrade) RealtimeTrade() error {

	for {
		// 1. get ticker
		ticker := g.GetTickerWithSave()

		// sleep
		time.Sleep(SLEEP_SECONDS)

		lastIndex := g.lastIndexCache.Get()
		if lastIndex == 0 {
			// 准备建仓
			g.OpenPosition(&ticker)
		} else {
			// 已经建仓了，处理不同的ticker
			g.ProcessOneTicker(&ticker)
		}
	}

}

func (g *GridTrade) OpenPosition(ticker *sdk.Ticker) error {
	return nil
}

func (g *GridTrade) ProcessOneTicker(ticker *sdk.Ticker) error {
	//var upper float64
	//var lower float64
	//oldIndex := g.lastIndexCache.Get()
	//for {
	//	last_index := g.lastIndexCache.Get()
	//
	//	if last_index > 0 {
	//
	//	}
	//}
	return nil
}

func GetIndex(close float64, priceLevels []float64) int {
	if len(priceLevels) < 2 {
		// 都不够一个网格
		return -1
	}
	index := len(priceLevels) - 1
	for i, price := range priceLevels {
		if close > price {
			index = i
			break
		}
	}
	return index
}

// 回测
func (g *GridTrade) BackTest(filename string) error {
	return nil
}
