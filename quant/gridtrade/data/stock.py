import pandas as pd

from pass_joinquant import JQ_ACCOUNT
from pass_joinquant import JQ_PASSWORD
from jqdatasdk import *
import datetime

auth(JQ_ACCOUNT, JQ_PASSWORD)

'''设置行列显示不忽略'''
pd.set_option('display.max_rows', 10000)  # 显示的行数最大值
pd.set_option('display.max_columns', 100) # 显示的列数最大值


'''交易所的code
# 上海证券交易所 .XSHG '600519.XSHG' 贵州茅台
# 深圳证券交易所 .XSHE '000001.XSHE' 平安银行
'''



def get_stock_list():
    '''
    获取所有A股股票列表
    :return: stock_list
    '''
    stock_list = list(get_all_securities(types=['stock']))
    return stock_list

def get_single_price(code, frequency, start_date, end_date):
    '''
    获取单个股票行情数据
    :param code: 股票代码
    :param frequency: 频率，eg:daily(每日）
    :param start_date: 开始日期
    :param end_date: 结束日期
    :return: dataFrame
    '''
    # fq:复权(fuquan)，默认是前复权(pre)
    # 聚宽的数据，和东方财富比较match，不知道为啥和同花顺对不上
    df = get_price(code, start_date=start_date, end_date=end_date, frequency=frequency, panel=False)
    return df

def export_data(df, filename, type):
    '''
    导出股票相关数据
    :param df: 股票行情数据
    :param filename: 文件名
    :param type: 股票数据类型，可以是price finance
    :return: 无
    '''
    # todo: file_root要正确填写
    file_root = 'absolute_path' + type + '/' + filename + '.csv'
    df.to_csv(file_root)
    print('已成功存储至', file_root)


def transfer_price_freq(df, freq):
    '''
    将数据转换为指定周期：开盘价（周期第一天），收盘价（周期最后一天），最高价（本周期），最低价（本周期）
    :param df: 原始k线数据
    :param freq: 时间频率周期（目标周期）
    :return: 转化后的周期数据
    '''
    df_trans = pd.DataFrame()
    df_trans['open'] = df['open'].resample(freq).first()
    df_trans['close'] = df['close'].resample(freq).last()
    df_trans['high'] = df['high'].resample(freq).max()
    df_trans['low'] = df['low'].resample(freq).min()

    # 汇总统计：统计一下周期成交量、成交额（sum）, 暂时不需要返回
    # df_trans['volume(sum)'] = df['volume'].resample(freq).sum()
    # df_trans['money(sum)'] = df['money'].resample(freq).sum()

    return df_trans

def get_single_finance(code, date, statDate):
    '''
    获取单个股票财务指标
    :param code: 股票code
    :param date: 统计日期
    :param statDate: 财报日期
    :return: dataFrame
    '''

    # 基于盈利指标选股: eps, operating_profit, roe, inc_net_profit_year_on_year
    # eps	每股收益EPS(元)	每股收益(摊薄)＝净利润/期末股本；分子从单季利润表取值，分母取季度末报告期股本值；净利润指归属于母公司股东的净利润(元)
    # operating_profit	经营活动净收益(元)	营业总收入-营业总成本
    # roe	净资产收益率ROE(%)	归属于母公司股东的净利润*2/（期初归属于母公司股东的净资产+期末归属于母公司股东的净资产）
    # inc_total_revenue_year_on_year	营业总收入同比增长率(%)	营业总收入同比增长率是企业在一定期间内取得的营业总收入与其上年同期营业总收入的增长的百分比，以反映企业在此期间内营业总收入的增长或下降等情况
    df = get_fundamentals(query(indicator).filter(indicator.code == code), date=date, statDate=statDate)
    return df

def get_single_valuation(code, date, statDate):
    '''
    获取单个股票估值指标
    :param code: 股票code
    :param date: 统计日期
    :param statDate: 财报日期
    :return: dataFrame
    '''

    # pe_ratio	市盈率(PE, TTM)	每股市价为每股收益的倍数，反映投资人对每元净利润所愿支付的价格，用来估计股票的投资报酬和风险	市盈率（PE，TTM）=（股票在指定交易日期的收盘价 * 截止当日公司总股本）/归属于母公司股东的净利润TTM。
    df = get_fundamentals(query(valuation).filter(valuation.code == code), date=date, statDate=statDate)
    return df
