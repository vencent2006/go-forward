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
    df = df[df['signal'] != 0]
    df['profit_pct'] = (df['close'] - df['close'].shift(1)) / df['close'].shift(1)
    df = df[df['signal'] == -1]  # 只看平仓的
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

    return df


if __name__ == '__main__':
    data = week_period_strategy('000001.XSHE', 'daily', None, datetime.datetime.today())
    # signal: 1=>buy, -1=>sell
    # print(data[['close', 'weekday', 'buy_signal', 'sell_signal', 'signal']])
    print(data[['close', 'signal', 'profit_pct']])
    print(data.describe())
    data['profit_pct'].plot()
    plt.show()

