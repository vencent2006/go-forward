import time
from decimal import Decimal

import matplotlib.pyplot as plt
import matplotlib.ticker as mticker
import numpy as np
import pandas as pd

import sdk_wrapper
from history import History


def print_list(name: str, li: []):
    """
    打印list
    :param name:
    :param li:
    :return:
    """
    print(name)
    i = 0
    for p in li:  # 百分比区间
        print('%d %.2f' % (i, p))
        i += 1


def cal_price_levels(highest: float, lowest: float, perc_levels: []) -> []:
    """
    价格网格列表
    :param highest: 最高价
    :param lowest: 最低价
    :param perc_levels: 比例网格列表
    :return:
    """
    mid = (highest + lowest) / 2
    # 百分比区间
    return [mid * x for x in perc_levels]


def cal_perc_levels(grid: float, grid_num: int) -> []:
    """
    比例网格列表
    :param grid: 网格大小
    :param grid_num: 网格数量
    :return: list
    """
    half_grid_num = grid_num // 2  # 网格数目的1/2
    # 百分比区间
    return [x for x in np.arange(
        1 + grid * half_grid_num, 1 - grid * half_grid_num - grid / 2, -grid)]


class GridStrategy:

    def __init__(self, instId: str, highest: float, lowest: float, grid: float, grid_num: int,
                 budget: float):

        # 基础信息
        self.instId = instId  # 交易对
        self.highest = highest  # 最高价
        self.lowest = lowest  # 最低价
        self.grid = grid  # 网格中每一个格的比例
        self.grid_num = grid_num  # 网格数量
        self.budget = budget  # 预算
        self.balance = budget  # 余额，初始化时余额等于预算
        self.consume = 0.0  # 消费
        self.max_consume = 0.0  # 最大消费
        self.profit = 0.0  # 收益
        self.profit_ratio = 0.0  # 收益率
        self.perc_levels = cal_perc_levels(grid, grid_num)  # 比例网格列表
        self.price_levels = cal_price_levels(highest, lowest, self.perc_levels)  # 价格网格列表

        # 操作信息
        self.tickers = []  # 交易行情的list
        self.last_price_index = None  # 记录上一次穿越的网格
        self.target = 0.0  # 目标仓位
        self.opt_b_stack = []  # 操作栈, 主要是用来配对buy和sell
        self.opt_b = []  # operation buy list
        self.opt_s = []  # operation sell list

        # dataframe
        self.df_s = None  # sell的dataframe
        self.df_b = None  # buy的dataframe
        self.df_ticker = None  # 行情的dataframe

        # print param
        print('**** param ****')
        print('instId', instId)
        print('highest', highest)
        print('lowest', lowest)
        print('grid', grid)
        print('grid_num', grid_num)
        print('highest:%.2f, lowest:%.2f, mid:%.2f' % (highest, lowest, (highest + lowest) / 2))
        print_list('perc_levels', self.perc_levels)  # 百分比区间
        print_list('price_levels', self.price_levels)  # 价格区间

    # def compute_profit(self):
    #     self.profit = self.budget - self.consume
    #     self.profit_ratio = self.profit / self.budget

    def add_ticker(self, ticker: {}):
        """
        记录ticker
        :param ticker:
        :return:
        """
        close = ticker['close']
        ts = ticker['ts']
        self.tickers.append([ts, close])

    def add_consume(self, delta: float):
        """
        计算消费数值, 并同步计算最大消费值
        :param delta: 变化量
        :return:
        """
        # 累加consume
        self.consume += delta
        # 更新max_consume
        if self.consume > self.max_consume:
            self.max_consume = self.consume

    def open_a_position(self, ticker: {}):
        """
        建仓
        :param ticker, 从sdk获取的行情信息
        :return:
        """
        close = ticker['close']
        ts = ticker['ts']
        if self.last_price_index is not None:
            # 已经有了价格索引了，不是建仓的必要条件，直接返回
            print("不满足建仓条件 | last_price_index = ", self.last_price_index, ", not None")
            return

        print('尝试建仓 | processing %s | close %.2f' % (ts, close))

        # 1.寻找开仓的位置
        price_levels = self.price_levels
        target = 0.0
        last_price_index_tmp = 0
        opt_b_stack_tmp = []
        for i in range(len(price_levels)):
            # price_levels, 是从大到小排列的
            if close > price_levels[i]:
                opt_b_stack_tmp.append(i)
                last_price_index_tmp = i
                target = i / (len(price_levels) - 1)  # 计算目标建仓的比例
                print('\tstart with i = ', i, ', target = ', self.target)
                break

        # 2.判断是否找到开仓的位置
        if target is not None and target > 0:
            # 找到了开仓的位置，尝试发起try buy
            money = self.budget * target
            size = float(money / close)
            client_order_id = sdk_wrapper.gen_client_order_id()
            res = sdk_wrapper.try_buy(instId=instId, price=close, size=size, client_order_id=client_order_id)
            if not res:
                # 建仓失败
                print("建仓失败 | try_buy failed | close %.2f | client_order_id %s" % (close, client_order_id))
                return
            else:
                # 建仓成功
                exchange_order_id = res  # 交易所的订单id
                self.last_price_index = last_price_index_tmp  # 建仓成功，把last_price_index设置下
                self.target = target  # 记录下目标仓位信息
                print(
                    '\t建仓成功: close=%.2f, buy, percent=%.2f, target=%.2f, last_price_index=%d, exchange_order_id = %s'
                    % (close, target, target, self.last_price_index, exchange_order_id))
                # 添加记录
                h = History(instId, 1, close, size, money)
                self.opt_b.append([ts, close, 1])
                self.add_consume(money)  # 消费值要累加下
                for i in range(len(opt_b_stack_tmp)):
                    # 这个stack的比例可能不是一个grid，就是会是多个grid
                    self.opt_b_stack.append(h)  # buy stack 要记录下
        else:
            # 没有建到仓
            print('不满足建仓条件 | 达不到建仓的位置')

    def process_one_ticker(self, ticker: {}):
        """
        处理一个建仓后的ticker
        :return:
        """
        # 根据ticker判断要不要sell|buy
        close = ticker['close']
        ts = ticker['ts']

        # 1. 判定是普通ticker处理过程
        if self.last_price_index is None:
            # 不是建仓
            print("不满足建仓条件 | last_price_index = ", self.last_price_index, ", is None")
            return
        price_levels = self.price_levels
        print('普通过程 | processing %s | close %.2f | last_price_index=%s, price=%.2f, target=%.2f' % (
            ts, close, self.last_price_index, price_levels[self.last_price_index], self.target))

        # 2. 实时行情
        signal = False
        while True:

            # 1. 初始化 upper 和 lower
            upper = None
            lower = None
            if self.last_price_index > 0:
                upper = price_levels[self.last_price_index - 1]
            if self.last_price_index < len(price_levels) - 1:
                lower = price_levels[self.last_price_index + 1]

            # 2 判断:还不是最轻仓，继续涨，就再卖一档
            if upper != None and close > upper:
                if len(self.opt_b_stack) > 0:  # 必须得先buy才能卖

                    # 游标变化
                    last_price_index_tmp = self.last_price_index - 1
                    target = last_price_index_tmp / (len(price_levels) - 1)
                    signal = True
                    # sell的逻辑
                    money = self.budget * target
                    size = float(money / close)
                    client_order_id = sdk_wrapper.gen_client_order_id()
                    res = sdk_wrapper.try_sell(instId=instId, price=close, size=size, client_order_id=client_order_id)
                    if res == False:
                        # sell 异常，输出日志
                        print(
                            '\ttry_sell failed | upper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, '
                            'target=%.2f' % (
                                upper, close, price_levels[last_price_index_tmp], grid, target))
                    else:
                        # sell 成功
                        self.last_price_index = last_price_index_tmp  # 更新last_price_index
                        exchange_order_id = res  # 交易所的订单id
                        # todo 如为了避免滑点, 分别sell
                        print(
                            '\tupper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, target=%.2f, exchange_order_id=%s' % (
                                upper, close, price_levels[self.last_price_index], grid, target, exchange_order_id))
                        # 计算profit
                        count = self.opt_b_stack[len(self.opt_b_stack) - 1].count
                        temp = float(
                            ((Decimal(close) - Decimal(self.opt_b_stack[len(self.opt_b_stack) - 1].price)) * Decimal(
                                count)).quantize(
                                Decimal('0.00')))
                        self.profit += temp
                        money = float(Decimal(close) * Decimal(count))
                        self.opt_b_stack.pop()  # buy stack 要弹栈
                        self.add_consume(-money)
                        self.opt_s.append([ts, price_levels[self.last_price_index], -1])
                # 结束本次循环
                continue

            # 3. 判断:还不是最重仓，继续跌，再买一档
            if lower != None and close < lower:
                # 游标变化
                last_price_index_tmp = self.last_price_index + 1
                signal = True
                # buy的逻辑
                # try buy
                target = last_price_index_tmp / (len(price_levels) - 1)
                money = self.budget * target
                size = float(money / close)
                client_order_id = sdk_wrapper.gen_client_order_id()
                res = sdk_wrapper.try_buy(instId=instId, price=close, size=size, client_order_id=client_order_id)
                if res == False:
                    # 建仓失败
                    print(
                        '\ttry_buy failed | upper=%.2f, close=%.2f, price=%.2f, sell, percent=%.3f, '
                        'target=%.2f' % (
                            upper, close, price_levels[last_price_index_tmp], grid, target))
                else:
                    # buy成功
                    self.last_price_index = last_price_index_tmp
                    exchange_order_id = res  # 交易所的订单id
                    print('\tclose=%.2f, buy, percent=%.2f, target=%.2f, last_price_index=%d, exchange_order_id=%s' % (
                        close, target, target, self.last_price_index, exchange_order_id))
                    # 添加记录
                    h = History(instId, 1, close, size, money)
                    self.opt_b_stack.append(h)
                    self.opt_b.append([ts, close, 1])
                    self.add_consume(money)

                continue

            # 没到要交易的区间，退出循环
            break

    def run(self):
        """
        算法主循环
        :return:
        """
        # main loop
        ticker_cnt = 0  # ticker 计数
        max_ticker_cnt = 100
        sleep_seconds = 10
        while True:
            # 判定结束条件
            ticker_cnt += 1
            if ticker_cnt > max_ticker_cnt:
                print('reach %d times(max=%d), so stop' % (ticker_cnt, max_ticker_cnt))
                break

            # sleep
            time.sleep(sleep_seconds)

            # 1. 实时获取 close
            ticker = sdk_wrapper.get_latest_close(instId)
            if not ticker:
                print('failed | sdk_wrapper.get_latest_close(instId:%s)' % instId)
                continue
            print(ticker)

            # 2. 记录ticker
            self.add_ticker(ticker=ticker)

            # 3. 根据是否开仓了进行处理
            if self.last_price_index is None:
                # 3.1 开始建仓
                self.open_a_position(ticker=ticker)
            else:
                # 3.2 已经建仓了，处理普通的ticker
                self.process_one_ticker(ticker=ticker)
        return

    def draw_pic(self):
        """
        画图
        :return:
        """

        # ticker df
        self.df_ticker = pd.DataFrame(self.tickers, columns=['timestamp', 'close'])
        self.df_ticker.set_index('timestamp', drop=True, inplace=True)
        self.df_ticker.sort_index()

        # buy df
        self.df_b = pd.DataFrame(self.opt_b, columns=['timestamp', 'price', 'opt'])
        self.df_b.set_index('timestamp', drop=True, inplace=True)
        self.df_b.sort_index()

        # sell df
        self.df_s = pd.DataFrame(self.opt_s, columns=['timestamp', 'price', 'opt'])
        self.df_s.set_index('timestamp', drop=True, inplace=True)
        self.df_s.sort_index()

        plt.plot(self.df_ticker.index, self.df_ticker['close'])  # 按照close画图
        plt.xticks(rotation=40, size=7)
        # 设置标签
        plt.xlabel('day')
        plt.ylabel('close')  # 画的是收盘价
        plt.title(self.instId)  # title
        plt.gca().xaxis.set_major_locator(mticker.MultipleLocator(len(self.df_ticker) // 2))
        # plt.show()

        # 6. 作图
        plt.plot(self.df_b.index, self.df_b['price'], 'or', label='buy')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
        plt.plot(self.df_s.index, self.df_s['price'], 'og', label='sell')  # 这里'or'代表中的'o'代表画圈，'r'代表颜色为红色，后面的依次类推
        # plt.scatter(x1, y1, marker='o', label="circle")
        # plt.scatter(x2, y2, marker='^', label="triangle")
        plt.legend()  # 加上label
        plt.tight_layout()
        plt.ticklabel_format(axis="y", style='plain')  # plain, 关闭坐标轴科学计数法; sci，打开坐标轴科学计数法
        plt.show()

    def print_tickers(self):
        print("***** df_ticker ******")
        print(self.df_ticker)

    def print_op_buy(self):
        print("***** df_buy ******")
        print(self.df_b)

    def print_op_sell(self):
        print("***** df_sell ******")
        print(self.df_s)

    def print_profit(self):
        print("***** profit ******")
        print('profit: %.2f' % self.profit)
        print('profit_ratio: %.2f' % self.profit_ratio)

    def print_asset(self):
        print("***** asset ******")
        print('consume: %.2f' % self.consume)
        print('max_consume: %.2f' % self.max_consume)
        print('balance: %.2f' % self.balance)
        print('budget: %.2f' % self.budget)
        print('profit: %.2f' % self.profit)
        print('profit_ratio: %.2f' % self.profit_ratio)

    def stat(self):
        self.print_tickers()
        self.print_op_buy()
        self.print_op_sell()
        self.print_asset()


if __name__ == '__main__':
    # 1. var
    instId = 'BTC-USDT'
    lowest = 25000
    highest = 30000
    grid = 0.005
    grid_num = 10
    budget = 100000.0  # 余额

    # 2. run strategy
    strategy = GridStrategy(instId=instId, highest=highest, lowest=lowest, grid=0.005, grid_num=10, budget=budget)
    strategy.run()
    strategy.stat()
    strategy.draw_pic()
