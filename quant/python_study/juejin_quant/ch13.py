import numpy as np
import random
import matplotlib.pyplot as plt

INIT_MONEY = 1000


# 创建简易的市场模型
def simple_market(win_rate, play_cnt=1000, stock_num=9, position=0.01, commission=0.01, lever=False):
    my_money = np.zeros(play_cnt)
    my_money[0] = INIT_MONEY  # 初始资金
    once_chip = my_money[0] * position  # 初始仓位
    lose_count = 1
    binomial = np.random.binomial(stock_num, win_rate, play_cnt)  # 伯努利分布
    for i in range(1, play_cnt):
        if binomial[i] > stock_num // 2:
            my_money[i] = my_money[i - 1] + once_chip if lever == False else my_money[i - 1] + once_chip * lose_count
            lose_count = 1
        else:
            my_money[i] = my_money[i - 1] - once_chip if lever == False else my_money[i - 1] - once_chip * lose_count
            lose_count += 1
        my_money[i] -= commission
        if my_money[i] <= 0:
            break
    return my_money


# 创建简易的市场模型应用仓位管理
def posit_manage(play_cnt=1000, stock_num=9, commission=0.01):
    my_money = np.zeros(play_cnt)
    my_money[0] = INIT_MONEY

    for i in range(1, play_cnt):
        win_rate = np.random.random(size=1)  # 生成[0,1)之间的浮点数
        binomial = np.random.binomial(stock_num, win_rate, 1)
        once_chip = my_money[0] * win_rate * 0.1

        if binomial > stock_num // 2:
            my_money[i] = my_money[i - 1] + once_chip
        else:
            my_money[i] = my_money[i - 1] - once_chip
        my_money[i] -= commission
        if my_money[i] <= 0:
            break
    return my_money


def strategy1(trader):
    # 概率50% 无手续费 参加1000次
    cnt = 1000

    # 第一行第一列图形
    ax1 = plt.subplot(4, 2, 1)
    # 第一行第二列图形
    ax2 = plt.subplot(4, 2, 2)

    plt.sca(ax1)
    _ = [plt.plot(np.arange(cnt), simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0), alpha=0.6) for _ in
         np.arange(0, trader)]
    plt.title(u"概率50% 无手续费 参加1000次 线图", fontsize=8)
    plt.sca(ax2)
    _ = plt.hist([simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0)[-1] for _ in np.arange(0, trader)], bins=30)
    plt.title(u"初始资金{}下,{}人的收益分布".format(INIT_MONEY, trader), fontsize=8)


def strategy2(trader):
    # 概率50% 手续费0.01 参加500000次
    cnt = 500000

    # 第一行第一列图形
    ax3 = plt.subplot(4, 2, 3)
    # 第一行第二列图形
    ax4 = plt.subplot(4, 2, 4)

    plt.sca(ax3)
    _ = [plt.plot(np.arange(cnt), simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0.01), alpha=0.6) for _ in
         np.arange(0, trader)]
    plt.title(u"概率50% 手续费0.01 参加500000次 线图", fontsize=8)
    plt.sca(ax4)
    _ = plt.hist([simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0.01)[-1]
                  for _ in np.arange(0, trader)], bins=30)
    plt.title(u"初始资金{}下,{}人的收益分布".format(INIT_MONEY, trader), fontsize=8)


def strategy3(trader):
    # 概率50% 无手续费 参加1000次 始能加注
    cnt = 1000

    # 第一行第一列图形
    ax3 = plt.subplot(4, 2, 5)
    # 第一行第二列图形
    ax4 = plt.subplot(4, 2, 6)

    plt.sca(ax3)
    _ = [plt.plot(np.arange(cnt), simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0, lever=True), alpha=0.6)
         for _ in np.arange(0, trader)]

    plt.title(u"概率50% 无手续费 参加1000次 始能加注 线图", fontsize=8)
    plt.sca(ax4)
    _ = plt.hist([simple_market(0.5, play_cnt=cnt, stock_num=9, commission=0, lever=True)[-1]
                  for _ in np.arange(0, trader)], bins=30)
    plt.title(u"初始资金{}下,{}人的收益分布".format(INIT_MONEY, trader), fontsize=8)


def strategy4(trader):
    # 仓位管理 手续费0.01 参加1000次
    # 第一行第一列图形
    ax3 = plt.subplot(4, 2, 7)
    # 第一行第二列图形
    ax4 = plt.subplot(4, 2, 8)

    cnt = 1000

    plt.sca(ax3)
    _ = [plt.plot(np.arange(cnt), posit_manage(play_cnt=cnt, stock_num=9, commission=0.01), alpha=0.6) for _ in
         np.arange(0, trader)]
    plt.title(u"带仓位管理、手续费的线图", fontsize=8)

    plt.sca(ax4)
    _ = plt.hist([posit_manage(play_cnt=cnt, stock_num=9, commission=0)[-1] for _ in np.arange(0, trader)], bins=30)
    plt.title(u"初始资金{}下,{}人的收益分布".format(INIT_MONEY, trader), fontsize=8)


if __name__ == '__main__':
    trader_cnt = 50
    plt.figure(1)
    plt.tight_layout()
    # 调整子图间的距离，主要是行高
    plt.subplots_adjust(left=None, bottom=None, right=None, top=None, wspace=0.5, hspace=1.0)
    strategy1(trader=trader_cnt)
    strategy2(trader=trader_cnt)
    strategy3(trader=trader_cnt)
    strategy4(trader=trader_cnt)
    plt.rcParams['font.sans-serif'] = ['Heiti TC']  # 显示中文
    plt.show()
