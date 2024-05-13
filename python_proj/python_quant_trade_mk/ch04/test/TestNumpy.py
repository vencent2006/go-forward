import numpy as np
from unittest import TestCase

class TestNumpy(TestCase):
    """ 读取指定列
        numpy.loadtxt需要传入4个关键字参数：
        1.fname是文件名，数据类型为字符串str；
        2.delimiter是分隔符，数据类型为字符串str；
        3.usecols是读取的列数，数据类型为元组tuple, 其中元素个数有多少个，则选出多少列；
        4.unpack是是否解包，数据类型为布尔bool。
    #"""
    def testReadFile(self):
        file_name = "demo.csv"
        end_price, volumn = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(2, 6),
            unpack=True
        )
        print(end_price)
        print(volumn)

    # 计算最大值与最小值
    def testMaxAndMin(self):
        file_name = "./demo.csv"
        high_price, low_price = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(4, 5),
            unpack=True
        )
        print("max_price={}".format(high_price.max()))
        print("min_price={}".format(low_price.min()))

    # 计算极差
    # 计算股价近期最高价的最大值和最小值的差值 和 计算股价近期最低价的最大值和最小值的差值
    def testPtp(self):
        file_name = "./demo.csv"
        high_price, low_price = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(4, 5),
            unpack=True
        )
        print("max - min of high price:{}".format(np.ptp(high_price)))
        print("max - min of low price:{}".format(np.ptp(low_price)))

    # 计算成交量加权平均价格
    # 成交量加权平均价格，英文名VWAP(Volume-Weighted Average Price，成交量加权平均价格）是一个非常重要的经济学量，代表着金融资产的“平均”价格
    def testAVG(self):
        file_name = "./demo.csv"
        end_price, volumn = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(2, 6),
            unpack=True
        )
        print("avg_price={}".format(np.average(end_price)))
        print("VWAP={}".format(np.average(end_price, weights=volumn)))

    # 计算中位数
    # 收盘价的中位数
    def testMedian(self):
        file_name = "./demo.csv"
        end_price, volumn = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(2, 6),
            unpack=True
        )
        print("median={}".format(np.median(end_price)))

    # 计算方差
    # 收盘价的方差
    def testVar(self):
        file_name = "./demo.csv"
        end_price, volumn = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(2, 6),
            unpack=True
        )
        print("var={}".format(np.var(end_price)))
        print("var={}".format(end_price.var()))

    # 计算股票收益率、年波动率及月波动率
    # 波动率是对价格变动的一种度量，历史波动率可以根据历史价格数据计算得出。计算历史波动率时，需要用到对数收益率
    # 年波动率等于对数收益率的标准差除以其均值，再乘以交易日的平方根，通常交易日取250天
    # 月波动率等于对数收益率的标准差除以其均值，再乘以交易月的平方根。通常交易月取12月
    def testVolatility (self):
        file_name = "./demo.csv"
        end_price, volumn = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(2, 6),
            unpack=True
        )
        returns = np.diff(np.log(end_price))

        annual_volatility = returns.std() / returns.mean() * np.sqrt(250)

        monthly_volatility = returns.std() / returns.mean() * np.sqrt(12)

        print("annual_volatility={}".format(annual_volatility))
        print("monthly_volatility={}".format(monthly_volatility))


    def testOBV(self):
        file_name = "./demo.csv"
        closing_prices, volumes  = np.loadtxt(
            fname=file_name,
            delimiter=',',
            usecols=(3, 6),
            unpack=True
        )
        # 交易日的差分（后一天减去前一天）组成的差分数组，正的是赚了，负的是赔了
        diff_closing_price = np.diff(closing_prices)
        sign_closing_price = np.sign(diff_closing_price)
        print(diff_closing_price)
        print(sign_closing_price)
        # 利用条件筛选来得到盈亏指标：一参是差分数组，二参是条件数组，三参是各个条件对应的值数组
        sign_closing_price = np.piecewise(
            diff_closing_price,
            [diff_closing_price < 0,
             diff_closing_price == 0,
             diff_closing_price > 0],
            [-1, 0, 1]
        )
        # 得到盈亏量，带正负的
        obvs = volumes[1:] * sign_closing_price
        print(obvs)