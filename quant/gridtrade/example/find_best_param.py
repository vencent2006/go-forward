"""
@desc: 寻找最优参数(以MA双均线策略为例)
"""
import pandas as pd

import strategy.ma_strategy as ma
import data.stock as st

# 参数1: 股票池
stocks = ['000001.XSHE']
code = '000001.XSHE'
start_date = '2016-01-01'
end_date = '2021-01-01'
data = st.get_csv_data(code, start_date, end_date)
# data = st.get_single_price(code, 'daily', start_date, end_date)

# 参数2：周期参数
params = [5, 10, 20, 60, 120, 250]

# 存放参数与收益
res = []

# 匹配并计算不同的周期参数对：5-10，5-20 ...... 120-250
for short in params:
    for long in params:
        if long > short:
            # data_res是中间结果
            data_res = ma.ma_strategy(df=data, short_window=short, long_window=long)
            # 获取周期参数，及其对应累计收益率
            cum_profit = data_res['cum_profit'].iloc[-1]  # 获取乐基收益率最终数
            res.append([short, long, cum_profit])  # 将参数放入结果列表

# 将结果列表转换为df，并找到最优参数
res = pd.DataFrame(res, columns=['short_win', 'long_win', 'cum_profit'])
# 排序
# todo: 跟老师的数据对不上呢
res = res.sort_values(by='cum_profit', ascending=False)  # 按收益倒序排列

print(res)
