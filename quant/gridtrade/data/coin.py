from binance.client import Client
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

# 从同包下的其他文件引入变量
from pass_binance import BINANCE_API_KEY
from pass_binance import BINANCE_SECRET_KEY


# 取历史k线数据
# 从binance中获取数据
# 只取前5列,date, open, high, low, close
def get_data_frame(symbol, interval, startTime, endTime):
    # 输入公钥私钥初始化接口
    client = Client(BINANCE_API_KEY, BINANCE_SECRET_KEY)
    bars = client.get_historical_klines(symbol, interval, startTime, endTime)
    # 只取前5列,date, open, high, low, close
    for line in bars:
        del line[5:]
    df = pd.DataFrame(bars, columns=['date', 'open', 'high', 'low', 'close'])
    return df


# 导出数据
def export_data(symbol, interval, startTime, endTime):
    # 获取数据数据
    df = get_data_frame(symbol, interval, startTime, endTime)
    print(df)

    # fileName
    fileName = symbol + '-' + startTime + '-' + endTime + '.csv'
    print(fileName)

    # save to csv
    df.to_csv(fileName)


# 加载数据
def load_data(file_root):
    print("loading " + file_root)
    data = pd.read_csv(file_root, index_col='date')
    data = pd.DataFrame(data, columns=['open', 'high', 'low', 'close'])
    return data


# 更新数据
def update_data(symbol):
    print(symbol)


if __name__ == '__main__':
    # 准备好数据获取需要的参数
    symbol = 'BTCUSDT'  # 比特币（美元计价）
    interval = Client.KLINE_INTERVAL_5MINUTE  # 5min间隔
    startTime = '2022-06-11'  # todo这个应该是格林尼治时间吧
    endTime = '2022-06-12'

    # bars = get_data_frame(symbol, interval, startTime, endTime)
    # print(bars)

    # save to csv
    # export_data(symbol, interval, startTime, endTime)

    # read from csv
    fileName = 'blockcoin/BTCUSDT-2022-06-11-2022-06-12.csv'
    res = load_data(fileName)
    print(res)
