import time

import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

from xq import get_data

stock_code = "SH512000"
# 从雪球获取数据
data = get_data(stock_code, "day")

data = np.array(data)

data = np.delete(data[:, 0:6], 1, axis=1)
# 时间戳转日期字符串
for row in data:
    row[0] = time.strftime("%Y-%m-%d", time.localtime(row[0] / 1000))

print(data)

data = pd.DataFrame(data, columns=['timestamp', 'open', 'high', 'low', 'close'])
data.set_index('timestamp', drop=True, inplace=True)
data.sort_index()

print(data)

plt.plot(data.index, data['close'])
# x轴文字倾斜程度
# plt.xticks(rotation=90, size=7)
plt.xticks(size=7)
# 设置标签
plt.xlabel('day')
plt.ylabel('close')
plt.title(stock_code)

from matplotlib.pyplot import MultipleLocator

# 把x轴的刻度间隔设置为40，并存在变量里
x_major_locator = MultipleLocator(40)
# ax为两条坐标轴的实例
ax = plt.gca()
# 把x轴的主刻度设置为40的倍数
ax.xaxis.set_major_locator(x_major_locator)
# 把x轴的刻度范围设置为0到1200
plt.xlim(0, 1200)
plt.show()
