import data.stock as st
import strategy.base as stb
import pandas as pd
import matplotlib.pyplot as plt

# 获取3只股票的数据：比亚迪，明德时代，隆基
codes = ['002594.XSHE', '300750.XSHE', '601012.XSHG']

# 容器：存放夏普
sharpe_list = []
for code in codes:
    data = st.get_single_price(code, 'daily', '2018-10-01', '2021-01-01')
    print(data.head())

    # 计算每只股票的夏普比率
    daily_sharpe, annual_sharpe = stb.calculate_sharpe(data)
    sharpe_list.append([code, annual_sharpe])  # 存放 [[c1,s1],[c2,s2]...]
    print(sharpe_list)

# 可视化3只股票比较
df_sharpe = pd.DataFrame(sharpe_list, columns=['code', 'sharpe']).set_index('code')
print(df_sharpe)

# 绘制bar图
df_sharpe.plot.bar(title='Compare Annual Sharpe Ratio')
plt.xticks(rotation=20)
plt.show()
