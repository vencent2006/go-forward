import pandas as pd

from data.pass_joinquant import JQ_ACCOUNT
from data.pass_joinquant import JQ_PASSWORD
from data.var_def import DATA_DIR
from jqdatasdk import *
import os
import datetime

auth(JQ_ACCOUNT, JQ_PASSWORD)

'''设置行列显示不忽略'''
pd.set_option('display.max_rows', 10000)  # 显示的行数最大值
pd.set_option('display.max_columns', 100)  # 显示的列数最大值

'''交易所的code
# 上海证券交易所 .XSHG '600519.XSHG' 贵州茅台
# 深圳证券交易所 .XSHE '000001.XSHE' 平安银行
'''


def init_db():
    """
    初始化股票数据库
    :return:
    """
    # 1. 获取所有股票代码
    stocks = get_stock_list()
    # 2. 存储到csv文件中
    for code in stocks:
        update_daily_price(stock_code=code, stock_type='price')


def get_stock_list():
    """
    获取所有A股股票列表
    上海证券交易所.XSHG
    深圳证券交易所.XSHE
    :return: stock_list
    """

    # 记住这里是list(index)
    stock_list = list(get_all_securities(['stock']).index)
    return stock_list


def get_single_price(code, frequency, start_date=None, end_date=None):
    """
    获取单个股票行情数据
    :param code: 股票代码
    :param frequency: 频率，eg:daily(每日）
    :param start_date: 开始日期
    :param end_date: 结束日期
    :return: dataFrame
    """
    # fq:复权(fuquan)，默认是前复权(pre)
    # 聚宽的数据，和东方财富比较match，不知道为啥和同花顺对不上

    # 如果start_date is None, 默认从上市日期开始
    if start_date is None:
        start_date = get_security_info(code).start_date
    # 如果end_date is None， 取最新日期（就是当日）
    if end_date is None:
        end_date = datetime.datetime.today()
    # 获取行情数据
    df = get_price(security=code, start_date=start_date, end_date=end_date, frequency=frequency, panel=False)
    return df


def export_data(df, filename, data_type, mode=None):
    """
    导出股票相关数据
    :param df: 股票行情数据
    :param filename: 文件名
    :param data_type: 股票数据类型，可以是price finance
    :param mode: a代表追加，none代表默认w写入
    :return: 无
    """

    file_root = DATA_DIR + data_type + '/' + filename + '.csv'
    df.index.names = ['date']
    if mode == 'a':
        df.to_csv(file_root, mode=mode, header=False)
        # 删除重复值
        df = pd.read_csv(file_root)  # 读取数据
        df = df.drop_duplicates(subset=['date'])  # 以日期列为准, 注意需要重新赋值
        df.to_csv(file_root, index=False)  # 重新写入 index=False表示不要再添加index列了
    else:
        df.to_csv(file_root)
    print('已成功存储至', file_root)


def get_csv_data(code, start_date, end_date, type='price'):
    """
    获取本地数据，且顺便完成数据更新工作
    :param code: str，股票代码
    :param start_date: 起始日期（包含）
    :param end_date: 结束日期（包含）
    :param type: 默认"price"
    :return: dataFrame
    """
    # 使用update直接更新
    update_daily_price(code, type)

    # 读取数据库对应的股票csv文件
    file_root = DATA_DIR + type + '/' + code + '.csv'
    df = pd.read_csv(file_root, index_col='date')

    # 根据日期参数筛选数据
    # 注意：要加括号，不然会以为是运算而不是条件判断
    return df[(df.index >= start_date) & (df.index <= end_date)]


def transfer_price_freq(df, freq):
    """
    将数据转换为指定周期：开盘价（周期第一天），收盘价（周期最后一天），最高价（本周期），最低价（本周期）
    :param df: 原始k线数据
    :param freq: 时间频率周期（目标周期）
    :return: 转化后的周期数据
    """
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
    """
    获取单个股票财务指标
    :param code: 股票code
    :param date: 统计日期
    :param statDate: 财报日期
    :return: dataFrame
    """

    # 基于盈利指标选股: eps, operating_profit, roe, inc_net_profit_year_on_year eps	每股收益EPS(元)	每股收益(
    # 摊薄)＝净利润/期末股本；分子从单季利润表取值，分母取季度末报告期股本值；净利润指归属于母公司股东的净利润(元) operating_profit	经营活动净收益(元)	营业总收入-营业总成本 roe
    # 净资产收益率ROE(%)	归属于母公司股东的净利润*2/（期初归属于母公司股东的净资产+期末归属于母公司股东的净资产） inc_total_revenue_year_on_year	营业总收入同比增长率(%)
    # 营业总收入同比增长率是企业在一定期间内取得的营业总收入与其上年同期营业总收入的增长的百分比，以反映企业在此期间内营业总收入的增长或下降等情况

    # 获取财务指标数据
    df = get_fundamentals(query(indicator).filter(indicator.code == code), date=date, statDate=statDate)
    return df


def get_single_valuation(code, date, statDate):
    """
    获取单个股票估值指标
    :param code: 股票code
    :param date: 统计日期
    :param statDate: 财报日期
    :return: dataFrame
    """

    # pe_ratio	市盈率(PE, TTM)	每股市价为每股收益的倍数，反映投资人对每元净利润所愿支付的价格，用来估计股票的投资报酬和风险	市盈率（PE，TTM）=（股票在指定交易日期的收盘价 *
    # 截止当日公司总股本）/归属于母公司股东的净利润TTM。

    # 获取估值数据
    df = get_fundamentals(query(valuation).filter(valuation.code == code), date=date, statDate=statDate)
    return df


def calculate_change_pct(df):
    """
    涨跌幅 = (当期收盘价-前期收盘价）/前期收盘价
    :param df: dataFrame，带有收盘价
    :return: dataFrame, 带有涨跌幅
    """
    df['close_pct'] = (df['close'] - df['close'].shift(1)) / df['close'].shift(1)
    return df


def update_daily_price(stock_code, stock_type):
    """
    更新每日股票行情
    :param stock_code: 股票代码
    :param stock_type: price（股票行情），finance（财务数据）
    :return:
    """
    # 3.1是否存在文件：不存在（重新获取），存在（执行3.2）
    file_root = DATA_DIR + stock_type + '/' + stock_code + '.csv'
    if os.path.exists(file_root):  # 如果存在对应文件
        # 3.2 获取增量数据（code，start_date=对应股票csv中最新日期，end_date=今天）
        start_date = pd.read_csv(file_root, usecols=['date'])['date'].iloc[-1]  # 取最后一个, 首先要定位到列上面
        df = get_single_price(stock_code, 'daily', start_date, datetime.datetime.today())
        # 3.3 追加到已有文件中
        export_data(df, stock_code, 'price', 'a')
    else:
        # 重新获取改股票行情数据
        df = get_single_price(stock_code, 'daily')
        export_data(df, stock_code, 'price')

    print('股票数据已经更新成功:', stock_code)

# if __name__ == '__main__':
#     code = '000001.XSHG'
#     start_date = '2021-02-01'
#     end_date = '2021-03-1'
#     frequency = 'daily'
#     df = get_price(security=code, start_date=start_date, end_date=end_date, frequency=frequency, panel=False)
#     print(df)
#     df2 = get_single_price(code=code,frequency=frequency,start_date=start_date,end_date=end_date)
#     print(df2)
