"""
@desc: 用来创建交易策略，生成交易信号
"""
import datetime
import matplotlib.pyplot as plt

import data.stock as st
import numpy as np


def compose_signal(df):
    """
    整合信号
    :param df: 行情数据，带signal
    :return: dataFrame
    """
    # 整合信号
    df['buy_signal'] = np.where((df['buy_signal'] == 1) & (df['buy_signal'].shift(1) == 1), 0, df['buy_signal'])
    df['sell_signal'] = np.where((df['sell_signal'] == -1) & (df['sell_signal'].shift(1) == -1), 0, df['sell_signal'])
    df['signal'] = df['buy_signal'] + df['sell_signal']
    return df


def calculate_prof_pct(df):
    """
    计算单次收益率: 开仓、平仓（开仓的全部股数）
    :param df: dataFrame，带signal的行情数据
    :return: dataFrame，带收益率profit_pct
    """
    # 筛选信号不为0的，并且计算涨跌幅
    df['profit_pct'] = df.loc[df['signal'] != 0, 'close'].pct_change()
    df = df[df['signal'] == -1]  # 只看平仓的
    return df


def calculate_cum_prof(df):
    """
    计算累加收益率
    :param df: dataFrame
    :return: dataFrame
    """
    # 累计收益
    df['cum_profit'] = (1 + df['profit_pct']).cumprod() - 1
    return df


def calculate_max_drawdown(df):
    """
    计算最大回撤比
    :param df: dataFrame
    :return: dataFrame
    """

    # 选取时间周期（时间窗口）
    window = 252  # 一年的工作日是252天
    # 计算时间周期中的最大净值
    df['roll_max'] = df['close'].rolling(window=window, min_periods=1).max()
    # 计算当天的回撤比 = (谷值 - 峰值) / 峰值 = 谷值/峰值 - 1
    df['daily_dd'] = df['close'] / df['roll_max'] - 1
    # 选取时间周期内最大的回撤比，即最大回撤;min因为是负值
    df['max_dd'] = df['daily_dd'].rolling(window=window, min_periods=1).min()

    return df


def week_period_strategy(code, frequency, start_date, end_date):
    """
    以"周"为周期的策略，周四买入，周一卖出
    :param code: 股票标的
    :param frequency: 频率
    :param start_date: 开始日期
    :param end_date: 结束日期
    :return: dataFrame
    """
    df = st.get_single_price(code, frequency, start_date, end_date)

    # 新建周期字段
    df['weekday'] = df.index.weekday
    # 周四买入, 1表示买入
    df['buy_signal'] = np.where((df['weekday'] == 3), 1, 0)
    # 周一卖出, -1表示卖出
    df['sell_signal'] = np.where((df['weekday'] == 0), -1, 0)

    # 整合信号
    df = compose_signal(df)

    # 计算单次收益率: 开仓、平仓（开仓的全部股数）
    df = calculate_prof_pct(df)

    # 计算累计收益率
    df = calculate_cum_prof(df)

    return df


if __name__ == '__main__':
    # data = week_period_strategy('000001.XSHE', 'daily', None, datetime.datetime.today())
    # # signal: 1=>buy, -1=>sell
    # # print(data[['close', 'weekday', 'buy_signal', 'sell_signal', 'signal']])
    # print(data[['close', 'signal', 'profit_pct', 'cum_profit']])
    # data[['profit_pct', 'cum_profit']].plot()
    # plt.show()

    # 平安银行的最大回撤
    data = st.get_single_price('000001.XSHE', 'daily', '2006-01-01', '2021-01-01')
    data = calculate_max_drawdown(data)
    print(data[['close', 'roll_max', 'daily_dd', 'max_dd']])
    data[['daily_dd', 'max_dd']].plot()
    plt.show()
