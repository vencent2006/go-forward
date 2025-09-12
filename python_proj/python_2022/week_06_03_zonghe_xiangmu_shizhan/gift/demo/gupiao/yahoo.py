import os

import yfinance as yf

# 获取特斯拉（TSLA）数据
tsla = yf.Ticker("TSLA")

# 获取历史行情
# hist = tsla.history(period="1d")  # 获取过去1年的数据
# print(hist.head())
proxy = "http://127.0.0.1:4780"
os.environ["http_proxy"] = proxy
os.environ["https_proxy"] = proxy

# 快速下载多只股票
# data = yf.download("AAPL MSFT GOOGL", start="2023-01-01", end="2023-12-31")
data = yf.download("AAPL", start="2023-12-01", end="2023-12-31")
print(data)
