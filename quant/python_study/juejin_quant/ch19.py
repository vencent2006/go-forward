import pandas as pd
import pandas_datareader.data as web
import datetime

# 获取上证指数交易数据 pandas-datareader模块data.DataReader()方法
df = web.DataReader("000001.SS", "yahoo", datetime.datetime(2017, 1, 1), datetime.datetime(2018, 2, 1))
df = pd.DataFrame(df)
print('#'*20)
print(df.head(10))  # 查看前几行
print('*'*20)

