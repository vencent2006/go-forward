"""
@desc: 用来创建交易策略，生成交易信号
"""

import data.stock as st
import numpy as np


def week_period_strategy(code, frequency, start_date, end_date):
    df = st.get_single_price(code, frequency, start_date, end_date)
    # 新建周期字段
    df['weekday'] = df.index.weekday
    # 周四买入, 1表示买入
    df['buy_signal'] = np.where((df['weekday'] == 3), 1, 0)
    # 周一卖出, -1表示卖出
    df['sell_signal'] = np.where((df['weekday'] == 0), -1, 0)
    # 模拟重复买入，周五买入
    # df['buy_signal'] = np.where((df['weekday'] == 3) | (df['weekday'] == 4), 1, 0)
    # 模拟重复卖出，周二卖出
    # df['sell_signal'] = np.where((df['weekday'] == 0) | (df['weekday'] == 1), -1, 0)
    # 整合信号
    df['buy_signal'] = np.where((df['buy_signal'] == 1) & (df['buy_signal'].shift(1) == 1), 0, df['buy_signal'])
    df['sell_signal'] = np.where((df['sell_signal'] == -1) & (df['sell_signal'].shift(1) == -1), 0, df['sell_signal'])
    df['signal'] = df['buy_signal'] + df['sell_signal']

    return df


if __name__ == '__main__':
    df = week_period_strategy('000001.XSHE', 'daily', '2020-01-01', '2020-03-01')
    # signal: 1=>buy, -1=>sell
    print(df[['close', 'weekday', 'buy_signal', 'sell_signal', 'signal']])
