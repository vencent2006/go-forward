package grid_trade

import (
	"example/quant_trade/cache"
	"example/quant_trade/sdk"
	"example/quant_trade/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	SLEEP_SECONDS = 10 * time.Second
	MIN_GRID_NUM  = 2 // 最小网格数量

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

	client         sdk.Exchange
	lastIndexCache *cache.LastIndex
	buyStackCache  *cache.BuyStack
}

func NewGridTrade(strategy string, instId string, highest float64, lowest float64, gridNum int, budget float64, client sdk.Exchange) *GridTrade {
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
	g.client = client
	g.lastIndexCache = cache.NewLastIndex(strategy)
	g.buyStackCache = cache.NewBuyStack(strategy)

	return g
}

func (g *GridTrade) GetTickerWithSave(instId string) (*sdk.GetTickerRes, error) {
	// get close
	req := &sdk.GetTickerReq{InstId: instId}
	ticker, err := g.client.GetTicker(req)
	if err != nil {
		logrus.Errorf("client.GetLatestClose() failed | err:%v", err)
		return nil, err
	}

	// todo save close
	return ticker, nil
}

func (g *GridTrade) RealtimeTrade() error {
	var (
		ticker *sdk.GetTickerRes
		err    error
	)
	for {
		// sleep
		time.Sleep(SLEEP_SECONDS)

		// 1. get ticker
		ticker, err = g.GetTickerWithSave(g.InstId)
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

func (g *GridTrade) ProcessOneTicker(ticker *sdk.GetTickerRes) error {

	newIndex := GetIndexByPrice(ticker.Close, g.PriceLevels)
	oldIndex := g.lastIndexCache.Get()
	return g.SetPositionLevel(oldIndex, newIndex, ticker.Close)
}

// 设定持仓位置
func (g *GridTrade) SetPositionLevel(oldIndex, newIndex int, closePrice float64) error {

	if newIndex == oldIndex {
		// no operation
		return nil
	}

	clientOrderId := utils.GetUUID()
	if newIndex > oldIndex {
		// sell
		money := g.Budget * float64((newIndex-oldIndex)/g.GridNum)
		price := g.PriceLevels[newIndex] // 取较低的价格 sell
		size := money / price
		return g.TrySell(g.InstId, price, size, clientOrderId)
	} else {
		// buy, newIndex < oldIndex
		money := g.Budget * float64((oldIndex-newIndex)/g.GridNum)
		price := g.PriceLevels[newIndex-1] // 取较高的价格 buy
		size := money / closePrice
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
	return g.TryPlaceOrder(sdk.SIDE_BUY, instId, price, size, clientOrderId)
}

func (g *GridTrade) TrySell(instId string, price float64, size float64, clientOrderId string) error {
	return g.TryPlaceOrder(sdk.SIDE_SELL, instId, price, size, clientOrderId)
}

func (g *GridTrade) TryPlaceOrder(side string, instId string, price float64, size float64, clOrdId string) error {

	var res *sdk.PlaceOrderRes
	var err error
	var ordId string

	if side != sdk.SIDE_BUY && side != sdk.SIDE_SELL {
		panic(fmt.Sprintf("invalid side %s", side))
	}

	res, err = g.client.PlaceOrder(&sdk.PlaceOrderReq{
		Side:    side,
		InstId:  g.InstId,
		ClOrdId: clOrdId,
		Price:   price,
		Size:    size,
	})

	// 1. 判定结果
	if err != nil {
		fmt.Printf("%s | instId:%s | price:%f | size:%f | clOrdId:%s | place order failed | err:%v",
			side, instId, price, size, clOrdId, err)
		return err
	}

	// place order successfully
	ordId = res.OrdId
	fmt.Printf("%s | instId:%s | price:%f | size:%f | clOrdId:%s | ordId:%s | place order successful",
		side, instId, price, size, clOrdId, ordId)
	// todo 记录到数据库中

	// 2. 查询订单
	deadline := time.Now().Add(TIME_WAIT_ORDER_RESULT)
	var orderData *sdk.GetOrderRes
	for time.Now().Before(deadline) {

		// 先 sleep
		time.Sleep(TIME_SLEET_ORDER_RESULT)

		// 获取订单信息
		reqGetOrderByClOrdId := &sdk.GetOrderByClOrdIdReq{
			InstId:  g.InstId,
			ClOrdId: clOrdId,
		}
		orderData, err = g.client.GetOrderByClOrdId(reqGetOrderByClOrdId)
		if err != nil {
			fmt.Printf("GetOrderByClOrdId(req: %+v) failed | err:%v", reqGetOrderByClOrdId, err)
			continue
		}

		// 取得了订单信息
		switch orderData.Status {
		case sdk.ORDER_STATUS_FILLED:
			// 已经成交，终态
			// todo 更新数据库状态
			return nil
		case sdk.ORDER_STATUS_CANCELED:
			// 已经取消，终态
			// todo 更新数据库状态
			return fmt.Errorf("order(clOrdId:%s, ordId:%s) already canceled", clOrdId, ordId)
		default:
			// 非终态
			continue
		}
	}

	// 如果还没交易成功，那么就取消订单，并返回false
	reqCancelOrder := &sdk.CancelOrderByClOrdIdReq{
		InstId:  g.InstId,
		ClOrdId: clOrdId,
	}
	_, err = g.client.CancelOrderByClOrdId(reqCancelOrder)
	if err != nil {
		fmt.Printf("CancelOrderByClOrdId(req: %+v) failed | err:%v", reqCancelOrder, err)
	}

	// 获取订单信息
	reqGetOrder := &sdk.GetOrderByClOrdIdReq{
		InstId:  g.InstId,
		ClOrdId: clOrdId,
	}
	for i := 0; i < 3; i++ {
		orderData, err = g.client.GetOrderByClOrdId(reqGetOrder)
		if err != nil {
			fmt.Printf("GetOrderByClOrdId(req: %+v) failed | err:%v", reqGetOrder, err)
		}
	}
	if err != nil {
		// todo 暂时不知道怎么处理
		panic(err)
	}

	// 取得了订单信息
	switch orderData.Status {
	case sdk.ORDER_STATUS_FILLED:
		// 已经成交，终态
		return nil
	case sdk.ORDER_STATUS_CANCELED:
		// 已经取消，终态
		return fmt.Errorf("clOrdId(%s) already canceled", clOrdId)
	default:
		// 非终态
		// todo 暂时不知道怎么处理
		// todo 更新数据库状态
		panic(fmt.Sprintf("status = %d | dont know how to handle", orderData.Status))
	}

}

// 回测
func (g *GridTrade) BackTest(filename string) error {
	return nil
}
