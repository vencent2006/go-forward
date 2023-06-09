package grid_trade

import (
	"example/quant_trade/cache"
	"example/quant_trade/sdk"
	"example/quant_trade/sdk/okx"
	"example/quant_trade/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	SLEEP_SECONDS = 10 * time.Second
	MIN_GRID_NUM  = 2 // 最小网格数量

	// side def
	SIDE_BUY  = "buy"
	SIDE_SELL = "sell"

	TIME_WAIT_ORDER_RESULT  = 5 * 60 * time.Second
	TIME_SLEET_ORDER_RESULT = 2 * time.Second
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
		PriceLevels:   utils.CalPriceLevels(highest, lowest, gridNum),
		PercentLevels: utils.CalPercentLevels(gridNum),
	}
	g.client = okx.NewOkxClient()
	g.lastIndexCache = cache.NewLastIndex(strategy)
	g.buyStackCache = cache.NewBuyStack(strategy)

	return g
}

func (g *GridTrade) GetTickerWithSave() (*sdk.Ticker, error) {
	// get close
	ticker, err := g.client.GetLatestClose()
	if err != nil {
		logrus.Errorf("client.GetLatestClose() failed | err:%v", err)
		return nil, err
	}

	// todo save close
	return &ticker, nil
}

func (g *GridTrade) RealtimeTrade() error {
	var (
		ticker *sdk.Ticker
		err    error
	)
	for {
		// sleep
		time.Sleep(SLEEP_SECONDS)

		// 1. get ticker
		ticker, err = g.GetTickerWithSave()
		if err != nil {
			logrus.Errorf("GetTickerWithSave failed | err:%v", err)
			continue
		}

		// 2. ProcessOneTicker
		if err = g.ProcessOneTicker(ticker); err != nil {
			// todo 如果err的话，要停止循环吗
			logrus.Errorf("ProcessOneTicker(ticker:%+v) failed | err:%v", ticker, err)
			continue
		}

	}

}

func (g *GridTrade) OpenPosition(ticker *sdk.Ticker) error {
	return nil
}

func (g *GridTrade) ProcessOneTicker(ticker *sdk.Ticker) error {

	newIndex := GetIndexByPrice(ticker.Close, g.PriceLevels)
	oldIndex := g.lastIndexCache.Get()
	return g.SetPositionLevel(oldIndex, newIndex, ticker.Close)
}

// 设定持仓位置
func (g *GridTrade) SetPositionLevel(oldIndex, newIndex int, closePrice float64) error {

	if newIndex == oldIndex {
		// no operation
		return nil
	} else if newIndex > oldIndex {
		// sell
		money := g.Budget * float64((newIndex-oldIndex)/g.GridNum)
		size := money / closePrice
		price := g.PriceLevels[newIndex] // 取较地的价格 sell
		clientOrderId := utils.GetUUID()
		return g.TrySell(g.InstId, price, size, clientOrderId)
	} else {
		// buy, newIndex < oldIndex
		money := g.Budget * float64((oldIndex-newIndex)/g.GridNum)
		size := money / closePrice
		price := g.PriceLevels[newIndex-1] // 取较高的价格 buy
		clientOrderId := utils.GetUUID()
		return g.TryBuy(g.InstId, price, size, clientOrderId)
	}
}

func GetIndexByPrice(close float64, priceLevels []float64) int {
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

func (g *GridTrade) TryBuy(instId string, price float64, size float64, clientOrderId string) error {
	return g.TryPlaceOrder(SIDE_BUY, instId, price, size, clientOrderId)
}

func (g *GridTrade) TrySell(instId string, price float64, size float64, clientOrderId string) error {
	return g.TryPlaceOrder(SIDE_SELL, instId, price, size, clientOrderId)
}

func (g *GridTrade) TryPlaceOrder(side string, instId string, price float64, size float64, clientOrderId string) error {

	var res *okx.PlaceOrderData
	var err error
	var exchangeOrderId string

	switch side {
	case SIDE_BUY:
		res, err = g.client.BuyLimit()
	case SIDE_SELL:
		res, err = g.client.SellLimit()
	default:
		panic(fmt.Sprintf("invalid side %s", side))
	}

	// 1. 判定结果
	if err != nil {
		fmt.Printf("%s | instId:%s | price:%f | size:%f | clientOrderId:%s | place order failed | err:%v",
			side, instId, price, size, clientOrderId, err)
		return err
	}

	// place order successfully
	exchangeOrderId = res.OrdId
	fmt.Printf("%s | instId:%s | price:%f | size:%f | clientOrderId:%s | exchangeOrderId:%s | place order successful",
		side, instId, price, size, clientOrderId, exchangeOrderId)
	// todo 记录到数据库中

	// 2. 查询订单
	deadline := time.Now().Add(TIME_WAIT_ORDER_RESULT)
	var orderData *okx.GetOrderData
	for time.Now().Before(deadline) {

		// 先 sleep
		time.Sleep(TIME_SLEET_ORDER_RESULT)

		// 获取订单信息
		orderData, err = g.client.GetOrderByClientOrderId(clientOrderId)
		if err != nil {
			fmt.Printf("GetOrderByClientOrderId(clientOrderId: %s) failed | err:%v", clientOrderId, err)
			continue
		}

		// 取得了订单信息
		switch orderData.State {
		case okx.ORDER_STATUS_FILLED:
			// 已经成交，终态
			// todo 更新数据库状态
			return nil
		case okx.ORDER_STATUS_CANCELED:
			// 已经取消，终态
			// todo 更新数据库状态
			return fmt.Errorf("order(clientOrderId:%s, exchangeOrderId:%s) already canceled", clientOrderId, exchangeOrderId)
		default:
			// 非终态
			continue
		}
	}

	// 如果还没交易成功，那么就取消订单，并返回false
	_, err = g.client.CancelOrderByClientOrderId(clientOrderId)
	if err != nil {
		fmt.Printf("GetOrderByClientOrderId(clientOrderId: %s) failed | err:%v", clientOrderId, err)
	}

	// 获取订单信息
	for i := 0; i < 3; i++ {
		orderData, err = g.client.GetOrderByClientOrderId(clientOrderId)
		if err != nil {
			fmt.Printf("GetOrderByClientOrderId(clientOrderId: %s) failed | err:%v", clientOrderId, err)
		}
	}
	if err != nil {
		// todo 暂时不知道怎么处理
		panic(err)
	}

	// 取得了订单信息
	switch orderData.State {
	case okx.ORDER_STATUS_FILLED:
		// 已经成交，终态
		return nil
	case okx.ORDER_STATUS_CANCELED:
		// 已经取消，终态
		return fmt.Errorf("clientOrderId(%s) already canceled", clientOrderId)
	default:
		// 非终态
		// todo 暂时不知道怎么处理
		// todo 更新数据库状态
		panic(fmt.Sprintf("status = %s | dont know how to handle"))
	}

}

// 回测
func (g *GridTrade) BackTest(filename string) error {
	return nil
}
