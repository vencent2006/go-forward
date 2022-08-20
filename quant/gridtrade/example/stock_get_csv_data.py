"""
@desc: 获取价格,并且计算涨跌幅
"""

import data.stock as st

# 1.获取平安银行的行情数据（日K）
code = '000002.XSHE'  # 平安银行
frequency = 'daily'
start_date = '2020-01-01'
end_date = '2020-02-01'
df = st.get_csv_data(code, start_date, end_date)
print(df)
exit()

# 2.计算日涨跌幅，验证有效性
df = st.calculate_change_pct(df)
print(df)  # 多了一列:close_pct

# 3.获取周K
df = st.transfer_price_freq(df, 'w')
print(df)

# 4.计算周涨跌幅，验证有效性
df = st.calculate_change_pct(df)
print(df)
