import akshare as ak

# 获取沪深京A股历史行情数据
stock_zh_a_hist_df = ak.stock_zh_a_hist(symbol="000001", period="daily", start_date="20230101", end_date="20231231")
print(stock_zh_a_hist_df.head())
