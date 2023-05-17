# This is a sample Python script.

import time
from decimal import Decimal

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

import api.xq
import history


def print_hi(name):
    # stock_code = "00700"
    stock_code = "SH512000"
    # stock_code = "SZ000002"

    # 6: 涨跌额 chg
    # 7: 涨跌幅 percent

    data = api.xq.get_data(stock_code, "day")
    data = np.array(data)

    data = np.delete(data[:, 0:6], 1, axis=1)
    new_data = []
    pre_close = 0
    # 遍历 data
    for row in data:
        # 日期转换
        row[0] = time.strftime("%Y-%m-%d", time.localtime(row[0] / 1000))

        # ce 差额
        # ce_up, ce_down = 0, 0

        if pre_close == 0:
            ce_up, ce_down = 0, 0
        else:
            ce_up = (row[2] - pre_close) / pre_close
            ce_down = (pre_close - row[3]) / pre_close
            # 1.1 如果 pre_close < low 说明纯纯上涨
            if pre_close < row[3]:
                ce_down = 0
            # 1.2 如果 pre_close > high 说明纯纯下跌
            if pre_close > row[2]:
                ce_up = 0

        pre_close = row[4]

        # 给 row 添加一列 np.append
        new_data.append(np.append(row, [ce_up, ce_down]))

    data = np.array(new_data)
    print(data)
    # data = np.array(map(lambda row: time.strftime("%Y-%m-%d", time.localtime(row[0] / 1000)), data[:, 0]))

    # 下面这个抄来的 没用上
    # 构建一个列表
    # data = np.array([np.insert(i, 0, np.array([10, 20]), axis=1) for i in data])

    # data 数据处理

    data = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close', 'ce_up', 'ce_down'])
    # data['timestamp'] = pd.to_datetime(data['timestamp'])
    data.set_index('timestamp', drop=True, inplace=True)
    data.sort_index()

    # 时间区间
    # data = data['2015-01-01':]
    data = data['2018-01-01':]
    # data = data['2018-06-01':'2019-06-01']

    data.to_csv(f'%s.csv' % (stock_code), encoding='utf_8_sig')

    print(data)
    # # plt
    # plt.figure(figsize=[10, 10])
    plt.plot(data.index, data['close'])
    # x轴文字倾斜程度
    # plt.xticks(rotation=90, size=7)
    plt.xticks(size=7)
    # 设置标签
    plt.xlabel('day')
    plt.ylabel('close')
    plt.title(stock_code)

    from matplotlib.pyplot import MultipleLocator
    # 把x轴的刻度间隔设置为1，并存在变量里
    x_major_locator = MultipleLocator(40)
    # ax为两条坐标轴的实例
    ax = plt.gca()
    # 把x轴的主刻度设置为1的倍数
    ax.xaxis.set_major_locator(x_major_locator)
    # 把x轴的刻度范围设置为-0.5到11，因为0.5不满一个刻度间隔，所以数字不会显示出来，但是能看到一点空白
    plt.xlim(0, 800)

    # plt.show()

    # 最大值中的最大值
    max = data['high'].max()
    min = data['low'].min()

    # 基准价
    benchmark = float(Decimal((max + min) / 2 * 1.2).quantize(Decimal('0.000')))
    print("最大值", max, "最小值", min, "基准价", benchmark)

    # 网格大小，目前是百分比 2.5%
    # grid = 0.025
    grid = 0.05
    count = 100
    max_consume_money = Decimal(0)
    consume_money = Decimal(0)
    opt = []
    # [日期, 价格, -1/1] 用于画点
    # 历史
    opt_b = []
    opt_s = []
    # 利润
    profit = 0

    # 建仓 todo 如果基准价 较 开盘价相差较大 open: 14.39, benchmark: 32.262
    # 应该是批量 (下跌格数) * [open, close]固定一个价格 买入，不应该下跌幅度买入
    first = data.iloc[0]
    # 基础持仓数量
    base_count = 500
    while benchmark * (1 - grid) > first['open']:
        # 一手买入 或者 倍数买入
        # 买入
        # 基准价变更
        benchmark = float(Decimal(benchmark * (1 - grid)).quantize(Decimal('0.000')))
        # print(data.index[0], "建仓买入", benchmark)

        # 计算买入 count

        # 计算的操作
        base_count += count

    print(data.index[0], "建仓买入, 基准价:", benchmark, base_count)

    consume_money += Decimal(benchmark) * Decimal(base_count)
    # 添加记录
    while base_count > 0:
        opt.append(history.History(stock_code, 1, benchmark, count))
        base_count -= count

    opt_b.append([data.index[0], benchmark, 1])

    max_consume_money = consume_money

    for day_up_down in data.index:
        open = data.loc[day_up_down].values[0]
        high = data.loc[day_up_down].values[1]
        low = data.loc[day_up_down].values[2]
        # close = data.loc[day_up_down].values[3]

        # 盘前
        # 如果 opt 为空，没有任何操作， 基准价 > 开盘价，触发买入
        # if len(opt) == 0 and benchmark > open:
        if benchmark * (1 - grid) > open:
            # 一手买入 或者 倍数买入
            # 买入
            # 基准价变更
            benchmark = open
            print(day_up_down, "开盘买入", benchmark)

            # 计算的操作
            consume_money += Decimal(benchmark) * Decimal(count)
            if consume_money.compare(max_consume_money) > 0:
                max_consume_money = consume_money

            # 添加记录
            h = history.History(stock_code, 1, open, count)
            opt.append(h)
            opt_b.append([day_up_down, benchmark, 1])
        elif benchmark * (1 + grid) <= open:
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
                h = history.History(stock_code, -1, benchmark, count)
                opt.pop()
                opt_s.append([day_up_down, benchmark, -1])

        while benchmark * (1 - grid) >= low:
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
            h = history.History(stock_code, 1, benchmark, count)
            opt.append(h)
            opt_b.append([day_up_down, benchmark, 1])

        # open = high开盘价就是最高价的情况 到时候再触发 low 会多买，这是一个假收益
        # 不会那么巧吧开盘价跟最高价一样
        if len(opt) > 0 and open != high:
            while len(opt) > 0 and benchmark * (1 + grid) <= high:
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
                h = history.History(stock_code, -1, benchmark, count)
                opt.pop()
                opt_s.append([day_up_down, benchmark, -1])

    # print(high, low)
    df_b = pd.DataFrame(opt_b, columns=['timestamp', 'price', 'opt'])
    # data['timestamp'] = pd.to_datetime(data['timestamp'])
    df_b.set_index('timestamp', drop=True, inplace=True)
    df_b.sort_index()
    df_s = pd.DataFrame(opt_s, columns=['timestamp', 'price', 'opt'])
    # data['timestamp'] = pd.to_datetime(data['timestamp'])
    df_s.set_index('timestamp', drop=True, inplace=True)
    df_s.sort_index()

    plt.plot(df_b.index, df_b['price'], 'or')
    plt.plot(df_s.index, df_s['price'], 'og')
    # plt.scatter(x1, y1, marker='o', label="circle")
    # plt.scatter(x2, y2, marker='^', label="triangle")
    plt.show()
    print("总收益", Decimal(profit).quantize(Decimal('0.000')))
    print("总资金消耗", Decimal(max_consume_money).quantize(Decimal("0.000")))


if __name__ == '__main__':
    print_hi('PyCharm')
