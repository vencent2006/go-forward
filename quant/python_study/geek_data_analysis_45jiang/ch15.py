# coding:utf-8
import tushare as ts
import datetime
import pandas as pd

df_sh = ts.get_hist_data('sh', start='2017-01-01', end=datetime.datetime.now().strftime('%Y-%m-%d'))
print(df_sh.info())  # 查看交易数据概览信息
print(df_sh.axes)  # 查看行和列的轴标签
# print(df_sh)

# todo  这个转换重要
df_sh.index = pd.to_datetime(df_sh.index)
df_sh.sort_index(inplace=True)  # 本地排序
print(df_sh.axes)  # 查看行和列的轴标签

# get_k_data 返回2008年数据
df_sh = ts.get_k_data('sh', start='2008-01-01', end=datetime.datetime.now().strftime('%Y-%m-%d'))
print(df_sh.head())
df_sh.index = pd.to_datetime(df_sh.date)
print(df_sh.head())
df_sh.drop(axis=1, columns='date', inplace=True)
print(df_sh.head())
