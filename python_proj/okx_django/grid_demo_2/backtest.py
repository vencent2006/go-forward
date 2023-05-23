from typing import Tuple

from pandas import DataFrame

from okx_sdk import market
from decimal import Decimal
import time
import history
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker


def backtest_grid_trading(df: pd.DataFrame) -> dict:
    """
    回测网格交易数据
    :param df: 行情数据
    :return:
    """

    # 3. 变量设定
    max = df['high'].max()  # high的max
    min = df['low'].min()  # low的min
    benchmark = float(Decimal((max + min) / 2 * 1.2).quantize(Decimal('0.000')))  # 基准价
    print('high.max=%.2f, low.min=%.2f, benchmark=%.2f' % (max, min, benchmark))
    print('high.max的那一行:\n', df.loc[df['high'].idxmax()])  # 查看是哪一行是最大
    print('low.min的那一行:\n', df.loc[df['low'].idxmin()])  # 查看是哪一行是最大

    grid = 0.05  # 网格大小比例
    count = 1  # 一手买多少
    max_consume_money = Decimal(0)  # 花费的钱的最大值，也就是你要准备的最大资金
    consume_money = Decimal(0)  # 花费的钱, buy会增加，sell会减少
    profit = 0  # 利润
    opt = []  # 操作记录，[日期, 价格, -1/1] 用于画点
    opt_b = []  # operation buy
    opt_s = []  # operation sell

    # 4. 建仓
    first = df.iloc[0]  # iloc是从0开始计数
    while benchmark * (1 - grid) > first['open']:  # 比较开盘价
        # todo 要不要统一成close收盘价呢
        # 一手买入 或者 倍数买入
        # 买入
        # 基准价变更
        benchmark = float(Decimal(benchmark * (1 - grid)).quantize(Decimal('0.000')))
        print(df.index[0], "建仓买入", benchmark)

        # 计算的操作
        consume_money += Decimal(benchmark) * Decimal(count)

        # 添加记录
        h = history.History(instId, 1, benchmark, count)
        opt.append(h)
        opt_b.append([df.index[0], benchmark, 1])

    max_consume_money = consume_money

    # 5. 根据行情进行交易
    for day_up_down in df.index:  # 获取index，再使用loc进行遍历
        open = df.loc[day_up_down].values[0]
        high = df.loc[day_up_down].values[1]
        low = df.loc[day_up_down].values[2]
        # close = data.loc[day_up_down].values[3]

        # 盘前
        # 如果 opt 为空，没有任何操作， 基准价 > 开盘价，触发买入
        # if len(opt) == 0 and benchmark > open:
        if benchmark * (1 - grid) > open:  # 比较开盘价
            # 一手买入 或者 倍数买入
            # 买入
            # 基准价变更
            benchmark = open
            print(day_up_down, "开盘买入", benchmark)

            # 计算的操作
            consume_money += Decimal(benchmark) * Decimal(count)
            if consume_money.compare(max_consume_money) > 0:  # 更新最大花费值
                max_consume_money = consume_money

            # 添加记录
            h = history.History(instId, 1, open, count)
            opt.append(h)
            opt_b.append([day_up_down, benchmark, 1])
        elif benchmark * (1 + grid) <= open:  # 比较开盘价
            if len(opt) > 0:
                # 卖出
                # 基准价变更
                benchmark = open

                # 计算的操作
                # 利润
                temp = float(
                    ((Decimal(benchmark) - Decimal(opt[len(opt) - 1].price)) * count).quantize(Decimal('0.00')))
                profit += temp
                consume_money -= Decimal(benchmark) * Decimal(count)

                print(day_up_down, "开盘卖出", benchmark, opt[len(opt) - 1].price, "收益", temp)
                # 修改记录
                h = history.History(instId, -1, benchmark, count)
                opt.pop()  # todo 为啥要pop呢，是要保证有买才能卖吗？
                opt_s.append([day_up_down, benchmark, -1])

        while benchmark * (1 - grid) >= low:  # 比较最低价
            # 盘中
            # 一手买入 或者 倍数买入
            # 买入
            # 基准价变更
            benchmark = float(Decimal(benchmark * (1 - grid)).quantize(Decimal('0.000')))
            print(day_up_down, "盘中买入", benchmark)

            # 计算的操作
            consume_money += Decimal(benchmark) * Decimal(count)
            if consume_money.compare(max_consume_money) > 0:
                max_consume_money = consume_money

            # 添加记录
            h = history.History(instId, 1, benchmark, count)
            opt.append(h)
            opt_b.append([day_up_down, benchmark, 1])

        # open = high开盘价就是最高价的情况 到时候再触发 low 会多买，这是一个假收益
        # 不会那么巧吧开盘价跟最高价一样
        if len(opt) > 0 and open != high:
            while len(opt) > 0 and benchmark * (1 + grid) <= high:  # 比较最高价
                # 卖出
                # 基准价变更
                benchmark = float(Decimal(benchmark * (1 + grid)).quantize(Decimal('0.000')))

                # 计算的操作
                temp = float(
                    ((Decimal(benchmark) - Decimal(opt[len(opt) - 1].price)) * count)
                    .quantize(Decimal('0.00')))
                profit += temp
                consume_money -= Decimal(benchmark) * Decimal(count)

                print(day_up_down, "盘中卖出", benchmark, opt[len(opt) - 1].price, "收益", temp)
                # 修改记录
                h = history.History(instId, -1, benchmark, count)
                opt.pop()
                opt_s.append([day_up_down, benchmark, -1])
    # buy df
    df_b = pd.DataFrame(opt_b, columns=['timestamp', 'price', 'opt'])
    df_b.set_index('timestamp', drop=True, inplace=True)
    df_b.sort_index()
    # sell df
    df_s = pd.DataFrame(opt_s, columns=['timestamp', 'price', 'opt'])
    df_s.set_index('timestamp', drop=True, inplace=True)
    df_s.sort_index()

    return {'df_b': df_b, 'df_s': df_s, 'profit': profit, 'max_consume_money': max_consume_money}


def get_df(instId: str, start='', end='') -> pd.DataFrame:
    """
    根据instId获取行情数据的dataframe
    :return: dataframe
    """
    bar = "1Dutc"  # utc的1天
    res = market.get_history_klines(instId, bar=bar, start=start, end=end)
    # print(res)
    data = np.array(res["data"])
    data = data[:, 0:5]
    # print(data)
    for row in data:
        row[0] = time.strftime("%Y-%m-%d", time.localtime(int(row[0]) / 1000))

    df = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close'])
    df.set_index('timestamp', drop=True, inplace=True)
    df = df.sort_index(ascending=True)  # 根据index的值进行排序，这里是要升序
    df[['open', 'high', 'low', 'close']] = df[['open', 'high', 'low', 'close']].apply(pd.to_numeric)
    return df


def draw_pic(df: DataFrame, dic: dict):
    plt.plot(df.index, df['close'])  # 按照close画图
    plt.xticks(rotation=40, size=7)
    # 设置标签
    plt.xlabel('day')
    plt.ylabel('close')  # 画的是收盘价
    plt.title(instId)
    plt.gca().xaxis.set_major_locator(ticker.MultipleLocator(len(df) // 5))
    # plt.show()

    # 6. 作图
    df_b = dic['df_b']
    df_s = dic['df_s']
    plt.plot(df_b.index, df_b['price'], 'or', label='buy')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
    plt.plot(df_s.index, df_s['price'], 'og', label='sell')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
    # plt.scatter(x1, y1, marker='o', label="circle")
    # plt.scatter(x2, y2, marker='^', label="triangle")
    plt.legend()  # 加上label
    plt.tight_layout()
    plt.show()

    profit = dic['profit']
    max_consume_money = dic['max_consume_money']
    print("总收益", Decimal(profit).quantize(Decimal('0.000')))
    # todo 这个总资金消耗，没看懂
    print("总资金消耗", Decimal(max_consume_money).quantize(Decimal("0.000")))


if __name__ == '__main__':
    # 1. 读取sdk数据，并格式化数据
    instId = "BTC-USDT"
    start = "1669824000000"  # 2022-12-01 00:00:00
    end = "1672502400000"  # 2023-01-01 00:00:00
    df = get_df(instId, start, end)
    print(df)
    # print(df.dtypes)  # 打印df的各个column的类型

    # 2. backtest
    dic = backtest_grid_trading(df)

    # 3. plot
    draw_pic(df, dic)
