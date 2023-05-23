import numpy as np
from decimal import Decimal
import pandas as pd
import constant
import os.path
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker


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


def backtest_grid_trading(df: pd.DataFrame, instId: str) -> dict:
    """
    网格策略
    :return:
    """
    # 1. 变量
    max = df['close'].max()
    min = df['close'].min()
    mid = (max + min) / 2
    # 百分比区间
    grid = 0.005
    perc_levels = [x for x in np.arange(
        1 + grid * 5, 1 - grid * 5 - grid / 2, -grid)]
    # 价格区间
    price_levels = [mid * x for x in perc_levels]
    # 记录上一次穿越的网格
    last_price_index = None
    # 目标仓位
    target = 0
    # 操作
    opt_b_stack = []  # 操作栈, 主要是用来配对buy和sell
    opt_b = []  # operation buy
    opt_s = []  # operation sell
    # 收益率
    profit = 0
    # 余额
    balance = 100000.0
    # 最大花费金额
    consume_money = 0
    max_consume_money = 0

    print('price_level', price_levels)
    i = 0
    for p in price_levels:
        print('%d %.1f' % (i, p))
        i += 1

    for day_up_down in df.index:  # 获取index，再使用loc进行遍历

        close = df.loc[day_up_down].values[3]

        if last_price_index == None:
            print('processing %s | close %.2f | last_price_index=%s, target=%.2f' % (
                day_up_down, close, last_price_index, target))
            # 2. 开仓
            for i in range(len(price_levels)):
                if close > price_levels[i]:
                    last_price_index = i
                    target = i / (len(price_levels) - 1)
                    print('\tstart with i = ', i, ', target = ', target)
                    break
            if target is not None and target > 0:
                print('\t建仓: close=%.2f, buy, percent=%.2f, target=%.2f, last_price_index=%d' % (
                    close, target, target, last_price_index))
                # 添加记录
                money = balance * target
                h = History(instId, 1, close, float(money / close), money)
                opt_b_stack.append(h)
                opt_b.append([day_up_down, close, 1])
                balance -= money
                continue
        else:
            print('processing %s | close %.2f | last_price_index=%s, price=%.2f, target=%.2f' % (
                day_up_down, close, last_price_index, price_levels[last_price_index], target))
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
                        opt_s.append([day_up_down, price_levels[last_price_index], -1])

                    continue
                # 还不是最重仓，继续跌，再买一档
                if lower != None and close < lower:
                    # 游标变化
                    last_price_index = last_price_index + 1
                    signal = True
                    # buy的逻辑
                    # todo 如为了避免滑点，可以分别buy
                    target = last_price_index / (len(price_levels) - 1)
                    print(
                        '\tclose=%.2f, buy, percent=%.3f, target=%.2f' % (price_levels[last_price_index], grid, target))
                    # 添加记录
                    money = balance * target
                    h = History(instId, 1, close, float(money / close), money)
                    opt_b_stack.append(h)
                    opt_b.append([day_up_down, price_levels[last_price_index], 1])
                    balance -= money

                    continue
                break

    # buy df
    df_b = pd.DataFrame(opt_b, columns=['timestamp', 'price', 'opt'])
    df_b.set_index('timestamp', drop=True, inplace=True)
    df_b.sort_index()
    # sell df
    df_s = pd.DataFrame(opt_s, columns=['timestamp', 'price', 'opt'])
    df_s.set_index('timestamp', drop=True, inplace=True)
    df_s.sort_index()

    # return {'df_b': df_b, 'df_s': df_s, 'profit': profit, 'max_consume_money': max_consume_money}
    return {'df_b': df_b, 'df_s': df_s, 'profit': profit}


def draw_pic(df: pd.DataFrame, dic: dict, title: str):
    plt.plot(df.index, df['close'])  # 按照close画图
    plt.xticks(rotation=40, size=7)
    # 设置标签
    plt.xlabel('day')
    plt.ylabel('close')  # 画的是收盘价
    plt.title(title)
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


if __name__ == '__main__':
    file_name = "BTC-USDT.csv"
    df = pd.read_csv(file_name, index_col=0)
    print(df)

    # res
    instId = os.path.basename(file_name)
    dic = backtest_grid_trading(df, instId)
    df_b = dic['df_b']
    df_s = dic['df_s']
    profit = dic['profit']

    # print
    print("***** df_b ******")
    print(df_b)
    print("***** df_s ******")
    print(df_s)
    print("***** profit ******")
    print('profit: %.2f' % profit)

    # 3. plot
    draw_pic(df, dic, instId)
