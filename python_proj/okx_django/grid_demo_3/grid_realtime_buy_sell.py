import time
from datetime import datetime
from decimal import Decimal

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from matplotlib import ticker

import constant
import okx_sdk
import utils
from okx_sdk import get_ticker


class History:
    """买入操作历史类"""
    # 操作 1 买入 -1 卖出
    stock_code = ""
    count = 0
    price = 0.0
    opt_type = 0
    money = 0.0  # 花费的钱数

    def __init__(self, stock_code, opt_type, price, count, money):
        self.stock_code = stock_code
        self.opt_type = opt_type
        self.price = price
        self.count = count
        self.money = money

    def __str__(self):
        return str(self.__class__) + ":" + str(self.__dict__)

    def __repr__(self):
        return str(self.__class__) + ":" + str(self.__dict__)


def get_latest_close(instId: str):
    """
    返回最新的close结构{'timestamp':ts, 'close': close}
    :param instId:
    :return: fail, return False; success, {'timestamp':ts, 'close': close}
    """
    res = get_ticker(instId)
    if res["code"] != '0':
        return False
    else:
        """
        {
          "code": "0",
          "msg": "",
          "data": [
            {
              "instType": "SPOT",
              "instId": "BTC-USDT",
              "last": "27145.9",
              "lastSz": "158.13880974",
              "askPx": "27147.5",
              "askSz": "118.87250831",
              "bidPx": "27147.4",
              "bidSz": "90.57158482",
              "open24h": "27011.5",
              "high24h": "27495.3",
              "low24h": "26980.7",
              "volCcy24h": "80250621069.41386445",
              "vol24h": "2942226.83285096",
              "ts": "1684893330510",
              "sodUtc0": "27225.3",
              "sodUtc8": "27324.5"
            }
        ]
        """
        timestamp = float(float(res['data'][0]['ts']) / 1000)
        time_array = time.localtime(timestamp)
        ts = time.strftime("%Y-%m-%d %H:%M:%S", time_array)
        close = float(res['data'][0]['last'])
        res = {'ts': ts, 'close': close, 'local_time': datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')[:-3]}
        print('get_ticker', instId, res)
        return res


def get_balance(instId: str):
    ccy = 'USDT'
    res = okx_sdk.get_one_account_balance(ccy=ccy)
    if res["code"] != '0':
        return False
    else:
        detail = res['data']['details'][0]
        avail_balance = detail['availBal']  # 可用余额
        cash_balance = detail['cashBal']  # 可用余额
        return {}


def gen_client_order_id():
    return utils.get_uuid()


def try_buy(instId: str, price, size, client_order_id):
    """
    尝试着去buy, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """
    return place_order('buy', instId=instId, price=price, size=size, client_order_id=client_order_id)


def try_sell(instId: str, price, size, client_order_id):
    """
    尝试着去sell, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """
    return place_order('sell', instId=instId, price=price, size=size, client_order_id=client_order_id)


def place_order(side: str, instId: str, price, size, client_order_id):
    """
    尝试着去buy, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """

    # 1. 下单，并获取下单结果
    if side == 'buy':
        res = okx_sdk.buy_limit_order(instId=instId, price=price, size=size, clientOrderId=client_order_id)
    elif side == 'sell':
        res = okx_sdk.sell_limit_order(instId=instId, price=price, size=size, clientOrderId=client_order_id)
    else:
        raise NameError("invalid side(%s)" % side)

    # 判定结果
    if res["code"] != '0':
        # sdk返回错误
        print('buy_limit_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return code', res['code'], 'return msg', res['msg'])
        return False
    elif res['data']['sCode'] != '0':
        # sdk返回错误
        print('buy_limit_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return sCode', res['data']['sCode'], 'return sMsg', res['data']['sMsg'])
        return False

    # place buy limit order, successfully
    exchange_order_id = res['data']['ordId']
    print('buy_limit_order successfully |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
          client_order_id, 'return exchange_order_id', exchange_order_id)

    # 2. 每隔5s，获取订单结果，最长等待5min，
    time_sleep_cnt = 0
    interval = 5
    while time_sleep_cnt < 60 * interval:
        time.sleep(interval)
        time_sleep_cnt += interval
        res = okx_sdk.get_order_by_client_order(instId, client_order_id)
        if res["code"] != '0':
            # sdk 返回异常
            print('get_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
            continue
        else:
            data = res['data'][0]
            state = data['state']
            if state == constant.ORDER_STATUS_FILLED:
                # 已经成交，终态
                return exchange_order_id
            elif state == constant.ORDER_STATUS_CANCELED:
                # 已经取消，终态
                return False
            else:
                # 非终态，继续查询
                continue

    # 如果还没交易成功，那么就取消订单，并返回false
    res = okx_sdk.cancel_order_by_client_order(instId=instId, client_order_id=client_order_id)
    if res["code"] != '0':
        print('cancel_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
    elif res['data']['sCode'] != '0':
        # sdk返回错误
        print('cancel_order_by_client_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return sCode', res['data']['sCode'], 'return sMsg', res['data']['sMsg'])
        return False

    # 取消可能成功, 再查询下
    res = okx_sdk.get_order_by_client_order(instId=instId, client_order_id=client_order_id)
    if res["code"] != '0':
        # sdk 返回异常
        print('after cancel order | get_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
    else:
        data = res['data'][0]
        state = data['state']
        if state == constant.ORDER_STATUS_FILLED:
            # 已经成交，终态
            return exchange_order_id
        elif state == constant.ORDER_STATUS_CANCELED:
            # 已经取消，终态
            return False
        else:
            # 非终态，继续查询
            print(
                'after cancel order | invalid state | get_order_by_client_order return state = %s(not final state)' % state)
            return False


def realtime_grid_trading(instId: str, highest: float, lowest: float, grid: float, grid_num: int,
                          balance: float) -> dict:
    """
    网格策略
    :param instId: 交易对
    :param highest: 最高上限
    :param lowest: 最低下限
    :param grid: 网格大小(小数)
    :param grid_num: 网格数量
    :return: {'df_ticker': df_ticker, 'df_b': df_b, 'df_s': df_s, 'profit': profit}
    """
    print('**** param ****')
    print('instId', instId)
    print('highest', highest)
    print('lowest', lowest)
    print('grid', grid)
    print('grid_num', grid_num)

    # 1. 变量
    mid = (highest + lowest) / 2
    half_grid_num = grid_num // 2  # 网格数目的1/2
    print('highest:%.2f, lowest:%.2f, mid:%.2f' % (highest, lowest, mid))
    # 百分比区间
    perc_levels = [x for x in np.arange(
        1 + grid * half_grid_num, 1 - grid * half_grid_num - grid / 2, -grid)]
    print('perc_levels', perc_levels)
    i = 0
    for p in perc_levels:
        print('%d %.2f' % (i, p))
        i += 1
    # 价格区间
    price_levels = [mid * x for x in perc_levels]
    print('price_levels', price_levels)
    i = 0
    for p in price_levels:
        print('%d %.1f' % (i, p))
        i += 1
    # 记录上一次穿越的网格
    last_price_index = None
    # 目标仓位
    target = 0
    # 操作
    opt_b_stack = []  # 操作栈, 主要是用来配对buy和sell
    opt_b = []  # operation buy
    opt_s = []  # operation sell
    # 收益
    profit = 0

    # 2. main loop
    ticker_cnt = 0  # ticker 计数
    max_ticker_cnt = 5
    tickers = []  # 记录所有的ticker
    while True:
        # ticker_cnt 计数
        ticker_cnt += 1
        if ticker_cnt > max_ticker_cnt:
            print('reach %s times, so stop' % ticker_cnt)
            break

        # sleep
        time.sleep(10)

        # 实时获取 close
        ticker = get_latest_close(instId)
        close = ticker['close']
        ts = ticker['ts']
        tickers.append([ts, close])

        if last_price_index == None:
            print('建仓过程 | processing %s | close %.2f | last_price_index=%s, target=%.2f' % (
                ts, close, last_price_index, target))
            # 2. 开仓
            for i in range(len(price_levels)):
                if close > price_levels[i]:
                    last_price_index = i
                    target = i / (len(price_levels) - 1)
                    print('\tstart with i = ', i, ', target = ', target)
                    break
            if target is not None and target > 0:
                # try buy
                money = balance * target
                size = float(money / close)
                client_order_id = gen_client_order_id()
                res = try_buy(instId=instId, price=close, size=size, client_order_id=client_order_id)
                if res == False:
                    # 建仓失败
                    print("建仓失败 | try_buy failed | ")
                    raise NameError("建仓失败")
                else:
                    # 建仓成功
                    exchange_order_id = res  # 交易所的订单id
                    print('\t建仓: close=%.2f, buy, percent=%.2f, target=%.2f, last_price_index=%d' % (
                        close, target, target, last_price_index))
                    # 添加记录
                    h = History(instId, 1, close, size, money)
                    opt_b_stack.append(h)
                    opt_b.append([ts, close, 1])
                    balance -= money
                    continue
            # 没有建到仓
            print('建仓过程 | 达不到建仓的位置')
            continue
        else:
            print('普通过程 | processing %s | close %.2f | last_price_index=%s, price=%.2f, target=%.2f' % (
                ts, close, last_price_index, price_levels[last_price_index], target))
            # 3. 实时行情
            signal = False
            while True:
                upper = None
                lower = None
                if last_price_index > 0:
                    upper = price_levels[last_price_index - 1]
                if last_price_index < len(price_levels) - 1:
                    lower = price_levels[last_price_index + 1]
                # 还不是最轻仓，继续涨，就再卖一档
                if upper != None and close > upper:
                    if len(opt_b_stack) > 0:  # 必须得先buy才能卖

                        # 游标变化
                        last_price_index = last_price_index - 1
                        target = last_price_index / (len(price_levels) - 1)
                        signal = True
                        # sell的逻辑
                        money = balance * target
                        size = float(money / close)
                        client_order_id = gen_client_order_id()
                        res = try_sell(instId=instId, price=close, size=size, client_order_id=client_order_id)
                        if res == False:
                            # sell失败
                            print(
                                '\ttry_sell failed | upper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, '
                                'target=%.2f' % (
                                    upper, close, price_levels[last_price_index], grid, target))
                            break
                        else:
                            # sell成功
                            exchange_order_id = res  # 交易所的订单id
                            # todo 如为了避免滑点，可以分别sell
                            print('\tupper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, target=%.2f' % (
                                upper, close, price_levels[last_price_index], grid, target))
                            # 计算profit
                            count = opt_b_stack[len(opt_b_stack) - 1].count
                            temp = float(
                                ((Decimal(close) - Decimal(opt_b_stack[len(opt_b_stack) - 1].price)) * Decimal(
                                    count)).quantize(
                                    Decimal('0.00')))
                            profit += temp
                            money = float(Decimal(close) * Decimal(count))
                            balance += money
                            opt_s.append([ts, price_levels[last_price_index], -1])

                    continue
                # 还不是最重仓，继续跌，再买一档
                if lower != None and close < lower:
                    # 游标变化
                    last_price_index = last_price_index + 1
                    signal = True
                    # buy的逻辑
                    # try buy
                    target = last_price_index / (len(price_levels) - 1)
                    money = balance * target
                    size = float(money / close)
                    client_order_id = gen_client_order_id()
                    res = try_buy(instId=instId, price=close, size=size, client_order_id=client_order_id)
                    if res == False:
                        # 建仓失败
                        print(
                            '\ttry_buy failed | upper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, '
                            'target=%.2f' % (
                                upper, close, price_levels[last_price_index], grid, target))
                    else:
                        # buy成功
                        exchange_order_id = res  # 交易所的订单id
                        print('\tclose=%.2f, buy, percent=%.2f, target=%.2f, last_price_index=%d' % (
                            close, target, target, last_price_index))
                        # 添加记录
                        h = History(instId, 1, close, size, money)
                        opt_b_stack.append(h)
                        opt_b.append([ts, close, 1])
                        balance -= money

                    continue

                # 没到要交易的区间，退出循环
                break

    # ticker df
    df_ticker = pd.DataFrame(tickers, columns=['timestamp', 'close'])
    df_ticker.set_index('timestamp', drop=True, inplace=True)
    df_ticker.sort_index()

    # buy df
    df_b = pd.DataFrame(opt_b, columns=['timestamp', 'price', 'opt'])
    df_b.set_index('timestamp', drop=True, inplace=True)
    df_b.sort_index()
    # sell df
    df_s = pd.DataFrame(opt_s, columns=['timestamp', 'price', 'opt'])
    df_s.set_index('timestamp', drop=True, inplace=True)
    df_s.sort_index()

    return {'df_ticker': df_ticker, 'df_b': df_b, 'df_s': df_s, 'profit': profit}


def draw_pic(df_ticker: pd.DataFrame, df_b: pd.DataFrame, df_s: pd.DataFrame, title: str):
    plt.plot(df_ticker.index, df_ticker['close'])  # 按照close画图
    plt.xticks(rotation=40, size=7)
    # 设置标签
    plt.xlabel('day')
    plt.ylabel('close')  # 画的是收盘价
    plt.title(title)
    plt.gca().xaxis.set_major_locator(ticker.MultipleLocator(len(df_ticker) // 2))
    # plt.show()

    # 6. 作图
    plt.plot(df_b.index, df_b['price'], 'or', label='buy')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
    plt.plot(df_s.index, df_s['price'], 'og', label='sell')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
    # plt.scatter(x1, y1, marker='o', label="circle")
    # plt.scatter(x2, y2, marker='^', label="triangle")
    plt.legend()  # 加上label
    plt.tight_layout()
    plt.ticklabel_format(axis="y", style='plain')  # plain, 关闭坐标轴科学计数法; sci，打开坐标轴科学计数法
    plt.show()


if __name__ == '__main__':
    # 1. var
    instId = 'BTC-USDT'
    lowest = 25000
    highest = 30000
    grid = 0.005
    grid_num = 10
    balance = 100000.0  # 余额

    # 2. call realtime_grid_trading
    dic = realtime_grid_trading(instId=instId, highest=highest, lowest=lowest, grid=0.005, grid_num=10,
                                balance=balance)
    df_ticker = dic['df_ticker']
    df_b = dic['df_b']
    df_s = dic['df_s']
    profit = dic['profit']

    # print
    print("***** df_ticker ******")
    print(df_ticker)
    print("***** df_b ******")
    print(df_b)
    print("***** df_s ******")
    print(df_s)
    print("***** profit ******")
    print('profit: %.2f' % profit)

    # 3. plot
    draw_pic(df_ticker, df_b, df_s, instId)
