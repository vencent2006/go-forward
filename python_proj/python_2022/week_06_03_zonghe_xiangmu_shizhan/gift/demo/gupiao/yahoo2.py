import os

import yfinance as yf

# end_time = datetime.now()
# start_time = end_time - timedelta(days=10)

start_time = '2025-08-01'
end_time = '2025-08-10'

proxy = "http://127.0.0.1:4780"
os.environ["http_proxy"] = proxy
os.environ["https_proxy"] = proxy

tk = "AAPL"  # AAPL TSLA
# data = yf.download(tickers=tk, start=start_time, end=end_time, interval="1h")
data = yf.download(tickers=tk, start=start_time, end=end_time)
data.to_csv(tk + ".csv", index=True)
print(data)
