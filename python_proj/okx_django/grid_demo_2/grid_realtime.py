from decimal import Decimal
import history
import time

"""
这个没完成，准备直接尝试基于固定价格线的算法
"""


def checkBalance(balance: Decimal, consume: Decimal, fee=0) -> bool:
    """
    检查余额是否足够
    :param balance: 余额
    :param consume: 计划消费的金额
    :param fee: 手续费
    :return: true: 余额足够; false: 余额不足
    """
    return balance >= (consume + fee)


def grid_realtime():
    """
    实时网格交易
    :return:
    """
    # 1. 变量设定
    instId = 'BTC-USDT'
    max = float(40000)  # 上限
    min = float(20000)  # 下限
    balance = Decimal(1000000)  # 余额
    benchmark = float(Decimal((max + min) / 2).quantize(Decimal('0.000')))  # 基准价
    print('max=%.2f, min=%.2f, benchmark=%.2f, balance=%.2f' % (max, min, benchmark, balance))

    grid = 0.05  # 网格大小比例
    count = 1  # 一手买多少
    max_consume_money = Decimal(0)  # 花费的钱的最大值，也就是你要准备的最大资金
    consume_money = Decimal(0)  # 花费的钱, buy会增加，sell会减少
    profit = 0  # 利润
    opt = []  # 操作记录，[日期, 价格, -1/1] 用于画点
    opt_b = []  # operation buy
    opt_s = []  # operation sell

    # 2. 建仓
    close = float(123)
    if benchmark >= close:
        # 一手买入 或者 倍数买入
        # 买入
        # 基准价变更
        benchmark = close
        # 计算的操作
        consume_money += Decimal(benchmark) * Decimal(count)
        if not checkBalance(balance, consume_money):
            # 余额不足，无法建仓
            print('余额不足，无法建仓: balance=%.2f, consume_money=%.2f, benchmark=%.2f' % (
                balance, consume_money, benchmark))
            return
        if consume_money.compare(max_consume_money) > 0:  # 更新最大花费值
            max_consume_money = consume_money
        # 添加记录
        print(close, "买入", benchmark)
        h = history.History(instId, 1, close, count)
        opt.append(h)
        opt_b.append([close, benchmark, 1])

    # 3. 循环获取价格
    while True:
        # sleep 下
        time.sleep(1)

        # 2. 获取ticker todo 补充下
        close = float(111)

        # 3. 止盈 or 止损
        # todo 平仓啥的在什么地方加呢

        # 4. 计算buy or sell
        if benchmark * (1 - grid) > close:
            # 一手买入 或者 倍数买入
            # 买入
            # 基准价变更
            benchmark = close

            # 计算的操作
            consume_money += Decimal(benchmark) * Decimal(count)
            if consume_money.compare(max_consume_money) > 0:  # 更新最大花费值
                max_consume_money = consume_money

            # 添加记录
            print(close, "买入", benchmark)
            h = history.History(instId, 1, close, count)
            opt.append(h)
            opt_b.append([day_up_down, benchmark, 1])
        elif benchmark * (1 + grid) <= close:  # 比较开盘价
            if len(opt) > 0:
                # 卖出
                # 基准价变更
                benchmark = close

                # 计算的操作
                # 利润
                temp = float(
                    ((Decimal(benchmark) - Decimal(opt[len(opt) - 1].price)) * count).quantize(Decimal('0.00')))
                profit += temp
                consume_money -= Decimal(benchmark) * Decimal(count)

                print(close, "卖出", benchmark, "之前买入价", opt[len(opt) - 1].price, "收益", temp)
                # 修改记录
                h = history.History(instId, -1, benchmark, count)
                opt.pop()  # todo 为啥要pop呢，是要保证有买才能卖吗？
                opt_s.append([close, benchmark, -1])
            else:
                break

        pass


if __name__ == '__main__':
    grid_realtime()
