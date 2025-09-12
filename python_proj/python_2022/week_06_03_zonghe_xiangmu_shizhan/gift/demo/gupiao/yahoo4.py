import os

import matplotlib.pyplot as plt
import pandas as pd

proxy = "http://127.0.0.1:4780"
os.environ["http_proxy"] = proxy
os.environ["https_proxy"] = proxy

# 获取茅台（600519.SS）的股票数据，日期范围从 2020-01-01 到 2021-01-01
# stock_data = yf.download('600519.SS', start='2020-01-01', end='2021-01-01', auto_adjust=False)
# stock_data.to_csv('600519.SS.csv')
stock_data = pd.read_csv('files/600519.SS.csv', skiprows=[1, 2])
print(stock_data)
print(type(stock_data))
print(stock_data.columns)
stock_data.rename(columns={'Price': 'Date'}, inplace=True)
stock_data['Date'] = pd.to_datetime(stock_data['Date'], format='%Y-%m-%d')
print(stock_data)
# 删除"Volume"和"Adj Close"列
# stock_data_cleaned = stock_data.drop(columns=['Adj Close', 'Volume'])

# 绘制茅台收盘价曲线
plt.figure(figsize=(10, 6))
plt.plot(stock_data['Close'], label='Close Price')
plt.title('Maotai Stock Price (2020)', fontsize=14)
plt.xlabel('Date', fontsize=12)
plt.ylabel('Close Price (CNY)', fontsize=12)
plt.legend()
plt.grid(True)
plt.show()
