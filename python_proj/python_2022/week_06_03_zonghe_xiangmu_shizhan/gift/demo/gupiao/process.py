import pandas as pd
import plotly.express as px

# 读取时跳过第二行（股票代码行）
d = pd.read_csv('files/AAPL.csv', skiprows=[1, 2])

# Price列更名为Date
d = d.rename(columns={'Price': 'date'})
print(d)
df = pd.DataFrame()
df["datetime"] = pd.to_datetime(d["date"])  # 转换为datetime类型
df["Close"] = d["Close"]
df.rename(columns={'Close': 'close'}, inplace=True)
print(df)

fig = px.line(df, x="datetime", y="close")  # 显式指定x和y轴
fig.show()
