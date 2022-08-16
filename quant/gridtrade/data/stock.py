import pandas as pd

from pass_joinquant import JQ_ACCOUNT
from pass_joinquant import JQ_PASSWORD
from jqdatasdk import *

auth(JQ_ACCOUNT, JQ_PASSWORD)

'''设置行列显示不忽略'''
pd.set_option('display.max_rows', 10000)  # 显示的行数最大值
pd.set_option('display.max_columns', 100) # 显示的列数最大值


'''交易所的code
# 上海证券交易所 .XSHG '600519.XSHG' 贵州茅台
# 深圳证券交易所 .XSHE '000001.XSHE' 平安银行
'''

# 获取所有A股的行情数据
# stocks = list(get_all_securities(types=['stock']))
# print(stocks)

# 获取股票行情数据
# for stock_code in stocks:
#     print("正在获取股票行情数据，股票代码：", stock_code)
#     # fq:复权(fuquan)，默认是前复权
#     # 聚宽的数据，和东方财富比较match，不知道为啥和同花顺对不上
#     df = get_price(stock_code, count=3, end_date='2022-06-08',
#                      frequency='daily', fields=['open', 'high', 'close', 'low'])
#     print(df)
#     time.sleep(3)

'''resample 函数的使用'''


# # 转换周期：日K转换为周K
# # 获取日K
# stock_code = '000001.XSHE'
# df = get_price(stock_code, start_date='2020-01-01', end_date='2020-12-31', frequency='daily', panel=False)
# # df.index.weekday是周几,0是周一
# df['weekday'] = df.index.weekday
# print(df)
#
# # 获取周K（当周的）：开盘价（当周第一天），收盘价（当周最后一天），最高价（当周），最低价（当周）
# df_week = pd.DataFrame()
# sample_period = 'M'
# df_week['open'] = df['open'].resample(sample_period).first()
# df_week['close'] = df['close'].resample(sample_period).last()
# df_week['high'] = df['high'].resample(sample_period).max()
# df_week['low'] = df['low'].resample(sample_period).min()
#
# # 汇总统计：统计一下月成交量、成交额（sum）
# df_week['volume(sum)'] = df['volume'].resample(sample_period).sum()
# df_week['money(sum)'] = df['money'].resample(sample_period).sum()
#
# # print df_week
#
# print(df_week)
# print('stock_code =', stock_code)
# print('sample_period =', sample_period)



'''获取股票财务指标'''
df = get_fundamentals(query(indicator), statDate='2020') # 获取财务指标数据
#df.to_csv('dir_path')

# 基于盈利指标选股: eps, operating_profit, roe, inc_net_profit_year_on_year
# eps	每股收益EPS(元)	每股收益(摊薄)＝净利润/期末股本；分子从单季利润表取值，分母取季度末报告期股本值；净利润指归属于母公司股东的净利润(元)
# operating_profit	经营活动净收益(元)	营业总收入-营业总成本
# roe	净资产收益率ROE(%)	归属于母公司股东的净利润*2/（期初归属于母公司股东的净资产+期末归属于母公司股东的净资产）
# inc_total_revenue_year_on_year	营业总收入同比增长率(%)	营业总收入同比增长率是企业在一定期间内取得的营业总收入与其上年同期营业总收入的增长的百分比，以反映企业在此期间内营业总收入的增长或下降等情况
df = df[(df['eps'] > 0) & (df['operating_profit'] > 2212172617) & (df['roe'] > 11) & (df['inc_net_profit_year_on_year'] > 10)]
print(df.head())

