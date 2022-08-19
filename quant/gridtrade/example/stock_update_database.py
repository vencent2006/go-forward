import data.stock as st

# 初始化变量
code = '000002.XSHE'  # 平安银行
frequency = 'daily'
start_date = '2021-01-01'
end_date = '2021-02-01'
export_type = 'price'  # 股票

# # 调用一只股票的行情
# df = st.get_single_price(code=code, frequency=frequency, start_date=start_date, end_date=end_date)
#
# # 保存文件
#
# st.export_data(df=df, filename=code, data_type=export_type, mode=None)
#
# # print df
# df = st.get_csv_data(code, 'price')
# print(df)

# 实时更新数据：假设每天更新日K数据 > 存到csv文件里面 > df.to_csv(append)

# 1. 获取所有股票代码

# 2. 存储到csv文件中

# 3. 每日更新数据
# st.update_daily_price(code, 'price')

st.init_db()
