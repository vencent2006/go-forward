import time

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from decimal import Decimal

from xq import get_data
import history

stock_code = "SH512000"
# 从雪球获取数据
data = get_data(stock_code, "day")

data = np.array(data)

data = np.delete(data[:, 0:6], 1, axis=1)
# 时间戳转日期字符串
for row in data:
    row[0] = time.strftime("%Y-%m-%d", time.localtime(row[0] / 1000))

print(data)

data = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close'])
data.set_index('timestamp', drop=True, inplace=True)
data.sort_index()

print(data)

# 最大值中的最大值
max = data['high'].max()
min = data['low'].min()
# 基准价
benchmark = float(Decimal((max + min) / 2 * 1.2).quantize(Decimal('0.000')))
print(max, min, benchmark)

grid = 0.05
count = 100
max_consume_money = Decimal(0)
consume_money = Decimal(0)  # 总资金消耗?
opt = []
# [日期, 价格, -1/1] 用于画点
# 历史
opt_b = []  # opt buy
opt_s = []  # opt sell
# 总利润
profit = 0

# 建仓
first = data.iloc[0]
while benchmark * (1 - grid) >= first['open']:
    # 一手买入 或者 倍数买入
    # 买入
    # 基准价变更
    benchmark = float(Decimal(benchmark * (1 - grid)).quantize(Decimal('0.000')))
    print(data.index[0], "建仓买入", benchmark)

    # 计算的操作
    consume_money += Decimal(benchmark) * Decimal(count)

    # 添加记录
    h = history.History(stock_code, 1, benchmark, count)
    opt.append(h)
    opt_b.append([data.index[0], benchmark, 1])

# 主逻辑
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

print("总收益", profit)
print("总资金消耗", consume_money)

df_b = pd.DataFrame(opt_b, columns=['timestamp', 'price', 'opt'])
df_b.set_index('timestamp', drop=True, inplace=True)
df_b.sort_index()

df_s = pd.DataFrame(opt_s, columns=['timestamp', 'price', 'opt'])
df_s.set_index('timestamp', drop=True, inplace=True)
df_s.sort_index()

# 设置标签
plt.xlabel('day')
plt.ylabel('close')
plt.title(stock_code)
plt.plot(df_b.index, df_b['price'], 'or')
plt.plot(df_s.index, df_s['price'], 'og')

plt.show()

print("总收益", Decimal(profit).quantize(Decimal('0.000')))
print("总资金消耗", Decimal(max_consume_money).quantize(Decimal("0.000")))
