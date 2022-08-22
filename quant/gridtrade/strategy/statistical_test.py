import data.stock as st
import strategy.ma_strategy as ma
import matplotlib.pyplot as plt
from scipy import stats


def ttest(data_return):
    """
    对策略收益进行t检验
    :param data_return: dateFrame, 单次收益率
    :return: float，t值和p值
    """

    # 调用假设检验ttest函数：scipy, 获取t、p, 并且忽略NAN
    t, p = stats.ttest_1samp(data_return, 0, nan_policy='omit')

    # 判断是否与理论均值有显著性差异: α=0.05
    p_value = p / 2  # 获取单边p值

    # 打印
    print("t-value:", t)
    print("p-value:", p_value)
    print("是否拒绝H0(True:拒绝，False:无法拒绝): 收益均值 = 0:", (p < 0.05))

    return t, p_value


if __name__ == '__main__':
    # 股票列表
    stocks = ['000001.XSHE', '000858.XSHE', '002594.XSHE']
    start_date = '2016-12-01'
    end_date = '2021-01-01'

    # 循环获取数据
    for code in stocks:
        print("******* 股票代码:", code)
        data = st.get_single_price(code, 'daily', start_date, end_date)
        data = ma.ma_strategy(data)

        # 策略的单次收益率
        returns = data['profit_pct']
        # print(returns)

        # 绘制一下分布图用于观察
        plt.hist(returns, bins=30)
        plt.show()

        # 对多个股票进行计算、测试
        ttest(returns)
