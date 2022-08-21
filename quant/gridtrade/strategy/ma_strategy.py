import matplotlib.pyplot as plt

import data.stock as st
import numpy as np
import strategy.base as strat
import pandas as pd
import matplotlib.pyplot as plt


def ma_strategy(df, short_window=5, long_window=20):
    """
    双均线策略
    :param df: dataFrame, 投资标的行情数据（必须包含收盘价）
    :param short_window: 短期n日移动平均线，默认5
    :param long_window: 长期n日移动平均线，默认20
    :return:
    """
    df = pd.DataFrame(df)

    # 计算技术指标：ma短期，ma长期
    df['short_ma'] = df['close'].rolling(window=short_window).mean()
    df['long_ma'] = df['close'].rolling(window=long_window).mean()

    # 生成信号：金叉买入，死叉卖出
    df['buy_signal'] = np.where(df['short_ma'] > df['long_ma'], 1, 0)
    df['sell_signal'] = np.where(df['short_ma'] < df['long_ma'], -1, 0)

    # 过滤信号：st.compose_signal
    df = strat.compose_signal(df)

    # 计算单次收益
    df = strat.calculate_prof_pct(df)

    # 计算累计收益
    df = strat.calculate_cum_prof(df)

    # print
    # print(df[['close', 'short_ma', 'long_ma', 'buy_signal', 'sell_signal', 'signal']])

    # 删除多余的column
    df.drop(labels=['buy_signal', 'sell_signal'], axis=1)

    return df


if __name__ == '__main__':
    # 股票列表
    stocks = ['000001.XSHE', '000858.XSHE', '002594.XSHE']
    # 存放累计收益率
    cum_profits = pd.DataFrame()
    # 循环获取数据
    for code in stocks:
        data = st.get_single_price(code, 'daily', '2016-01-01', '2021-01-01')
        data = ma_strategy(data)
        cum_profits[code] = data['cum_profit'].reset_index(drop=True)  # 存储累计收益率

        # 折线图
        data['cum_profit'].plot(label=code)
        # # 筛选有信号点
        # data = data[data['signal'] != 0]
        print('开仓次数:', int(len(data)))
        # print(data[['close', 'signal', 'profit_pct', 'cum_profit', ]])

    # 预览
    print(cum_profits)

    # 可视化
    # cum_profits.plot()
    plt.title('Comparison of Ma Strategy')
    plt.show()
