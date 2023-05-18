from okx_sdk import market
from decimal import Decimal
import time
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker

# 1. 读取sdk数据，并格式化数据
instId = "BTC-USDT"
bar = "1Dutc"
res = market.get_history_klines(instId, bar=bar)
# print(res)
data = np.array(res["data"])
data = data[:, 0:5]
print(data)
for row in data:
    row[0] = time.strftime("%Y-%m-%d", time.localtime(int(row[0]) / 1000))

df = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close'])
df.set_index('timestamp', drop=True, inplace=True)
df = df.sort_index(ascending=True)  # 根据index的值进行排序，这里是要升序
df[['open', 'high', 'low', 'close']] = df[['open', 'high', 'low', 'close']].apply(pd.to_numeric)
print(df)

# 2. plot
plt.plot(df.index, df['close'])  # 按照close画图
plt.xticks(rotation=20, size=7)
# 设置标签
plt.xlabel('day')
plt.ylabel('close')  # 画的是收盘价
plt.title(instId)
plt.gca().xaxis.set_major_locator(ticker.MultipleLocator(18))
plt.show()

# 3. 变量设定
max = df['high'].max()
min = df['low'].min()
benchmark = float(Decimal((max + min) / 2 * 1.2).quantize(Decimal('0.000')))  # 基准价
print(max, min, benchmark)
# print(df.loc[df['high'].idxmax()]) # 查看是哪一行是最大
# open      24118.9
# high     510000.0
# low       23550.4
# close     76000.0
# Name: 2023-03-14, dtype: float64