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


def calculate_sharpe(df):
    """
    计算夏普比率，返回的是年化的夏普比率
    :param df: 行情数据 dataFrame( stock )
    :return: float 夏普比例，年夏普比率
    """
    # 夏普比率：投资者额外承受的每一单位风险所获得的额外收益
    # 公式: sharpe = (回报率的均值 - 无风险利率)/回报率的标准差
    # 因子项
    daily_return = df['close'].pct_change()  # 回报率
    # 1.回报率的均值 = 回报率.mean = 日涨跌幅.mean
    avg_return = daily_return.mean()
    # 2.无风险利率 = 这里就直接取0，国债3%/252就约等于0了
    # 3.回报率的标准层 = 日涨跌幅.std
    std_return = daily_return.std()
    # 计算夏普: sharpe = 回报率的均值 / 回报率的标准差
    sharpe = avg_return / std_return
    # 年化夏普比率
    sharpe_year = sharpe * np.sqrt(252)
    return sharpe, sharpe_year
