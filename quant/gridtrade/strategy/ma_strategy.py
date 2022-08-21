import data.stock as st
import numpy as np
import strategy.base as stb
import pandas as pd


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
    df = stb.compose_signal(df)

    # print
    # print(df[['close', 'short_ma', 'long_ma', 'buy_signal', 'sell_signal', 'signal']])

    # 删除多余的column
    df.drop(labels=['buy_signal', 'sell_signal'], axis=1)

    return df


if __name__ == '__main__':
    df = st.get_single_price('000001.XSHE', 'daily', '2020-01-01', '2021-01-01')
    df = ma_strategy(df)
    df = df[df['signal'] != 0]
    print('开仓次数:', int(len(df)/2))
    print(df[['close', 'short_ma', 'long_ma', 'signal']])
