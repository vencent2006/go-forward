import pandas as pd
import plotly.express as px

ts_df = pd.read_csv("TSLA.csv", skiprows=[1, 2]).rename(columns={'Price': 'Date'})
ap_df = pd.read_csv("AAPL.csv", skiprows=[1, 2]).rename(columns={'Price': 'Date'})
ts_df["Date"] = pd.to_datetime(ts_df["Date"], format='%Y-%m-%d')
ap_df["Date"] = pd.to_datetime(ap_df["Date"], format='%Y-%m-%d')
print(ts_df)

df = pd.DataFrame()
df["ts"] = ts_df["Close"]
df["ap"] = ap_df["Close"]
print(type(df))
print(df)

fig = px.line(df)
fig.show()
