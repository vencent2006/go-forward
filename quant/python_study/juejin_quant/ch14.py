import pandas_datareader.data as web
import datetime
import pandas as pd
import matplotlib.pyplot as plt

# 获取上证指数的2017.1.1到至今的交易数据
code = '000001.SS'
source = 'yahoo'
# source = 'iex'
start_date = datetime.datetime(2017, 1, 1)
end_date = datetime.date.today()

df = web.DataReader(code, source, start_date, end_date)
df = pd.DataFrame(df)
print('------ df.head:\n {}'.format(df.head()))
print('------ df.tail:\n {}'.format(df.tail()))
print('------ df.shape:\n {}'.format(df.shape))
print('------ df.describe:\n {}'.format(df.describe()))
print('------ df.info:')
print(df.info())

df.Close.plot(c='b')
plt.legend(['close'], loc='best')
plt.show()
