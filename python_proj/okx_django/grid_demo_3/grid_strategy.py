import numpy as np
from decimal import Decimal


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


def get_close_realtime():
    return 0


def grid_strategy(instId, lowest, highest):
    """
    网格策略
    :return:
    """
    # 1. 变量
    mid = (highest + lowest) / 2
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

    while True:
        close = 12345
        if last_price_index == None:
            # 2. 开仓
            for i in range(len(price_levels)):
                if close > price_levels[i]:
                    last_price_index = i
                    target = i / (len(price_levels) - 1)
            if target > 0:
                print('建仓: close=%.2f, buy, percent=%.2f, target=%.2f' % (close, target, target))
                # 添加记录
                money = balance * target
                h = History(instId, 1, close, float(money / close), money)
                opt_b_stack.append(h)
                continue
        else:
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
                    if len(opt_b_stack) > 0:
                        # 必须得先buy才能卖
                        last_price_index = last_price_index - 1
                        signal = True
                        # sell的逻辑
                        # todo 如为了避免滑点，可以分别sell
                        print('close=%.2f, sell, percent=%.2f, target=%.2f' % (close, grid, target))
                        # 计算profit
                        temp = float(
                            ((Decimal(close) - Decimal(opt_b_stack[len(opt_b_stack) - 1].price)) * opt_b_stack[
                                len(opt_b_stack) - 1].count).quantize(Decimal('0.00')))
                        profit += temp
                    continue
                # 还不是最重仓，继续跌，再买一档
                if lower != None and close < lower:
                    last_price_index = last_price_index + 1
                    signal = True
                    # buy的逻辑
                    # todo 如为了避免滑点，可以分别buy
                    print('close=%.2f, buy, percent=%.2f, target=%.2f' % (close, grid, target))
                    # 添加记录
                    money = balance * target
                    h = History(instId, 1, close, float(money / close), money)
                    opt_b_stack.append(h)
                    continue
                break


if __name__ == '__main__':
    pass
