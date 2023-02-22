from pandas.tseries.holiday import USFederalHolidayCalendar
from pandas.tseries.offsets import CustomBusinessDay

US_BUSINESS_DAY = CustomBusinessDay(calendar=USFederalHolidayCalendar())
from pandas.tseries.holiday import USFederalHolidayCalendar
from pandas.tseries.offsets import CustomBusinessDay

US_BUSINESS_DAY = CustomBusinessDay(calendar=USFederalHolidayCalendar())
import pandas as pd
from backtesting import Strategy
from backtesting import Backtest
import pandas_ta as ta
import yfinance as yf
import plotly.graph_objects as go
from plotly.subplots import make_subplots


def CHOP(df, chop_len, atr_len):
    #  Calculate Choppiness
    chop_series = ta.chop(high=df['High'], low=df['Low'], close=df['Close'], length=chop_len, atr_length=atr_len)
    return chop_series


def plot_chart(i, symbol, df, current_price, buy_grid, sell_grid, buy_stop_loss, sell_stop_loss):
    light_palette = {}
    light_palette["bg_color"] = "#ffffff"
    light_palette["plot_bg_color"] = "#ffffff"
    light_palette["grid_color"] = "#e6e6e6"
    light_palette["text_color"] = "#2e2e2e"
    light_palette["dark_candle"] = "black"
    light_palette["light_candle"] = "steelblue"
    light_palette["volume_color"] = "#c74e96"
    light_palette["border_color"] = "#2e2e2e"
    light_palette["color_1"] = "#5c285b"
    light_palette["color_2"] = "#802c62"
    light_palette["color_3"] = "#a33262"
    light_palette["color_4"] = "#c43d5c"
    light_palette["color_5"] = "#de4f51"
    light_palette["color_6"] = "#f26841"
    light_palette["color_7"] = "#fd862b"
    light_palette["color_8"] = "#ffa600"
    light_palette["color_9"] = "#3366d6"
    palette = light_palette
    #  Array of colors for support/resistance lines
    buy_grid_colors = ["#e28743", "#e28743", "#e28743", "#e28743", "#e28743"]
    sell_grid_colors = ["#2596be", "#2596be", "#2596be", "#2596be", "#2596be"]
    #  Create sub plots
    fig = make_subplots(rows=1, cols=1, subplot_titles=[f"{i} {symbol} Chart", ], \
                        specs=[[{"secondary_y": False}]], \
                        vertical_spacing=0.04, shared_xaxes=True)
    #  Plot close price
    fig.add_trace(
        go.Scatter(x=df.index, y=df['Close'], line=dict(color="blue", width=1), name=f"Close"),
        row=1, col=1)
    #  Current price
    fig.add_hline(y=current_price, line_width=0.6, line_dash="solid", line_color="blue", row=1, col=1)
    #  Add buy and sell grids
    i = 0
    for level in buy_grid:
        line_color = buy_grid_colors[i] if i < len(buy_grid_colors) else buy_grid_colors[0]
        fig.add_hline(y=level, line_width=0.6, line_dash="dash", line_color=line_color, row=1, col=1)
        i += 1
    #  stop loss
    fig.add_hline(y=buy_stop_loss, line_width=0.6, line_dash="solid", line_color="red", row=1, col=1)
    i = 0
    for level in sell_grid:
        line_color = sell_grid_colors[i] if i < len(sell_grid_colors) else sell_grid_colors[0]
        fig.add_hline(y=level, line_width=1, line_dash="dash", line_color=line_color, row=1, col=1)
        i += 1
    #  stop loss
    fig.add_hline(y=sell_stop_loss, line_width=1, line_dash="solid", line_color="red", row=1, col=1)
    fig.update_layout(
        title={'text': '', 'x': 0.5},
        font=dict(family="Verdana", size=12, color=palette["text_color"]),
        autosize=True,
        width=1280, height=720,
        xaxis={"rangeslider": {"visible": False}},
        plot_bgcolor=palette["plot_bg_color"],
        paper_bgcolor=palette["bg_color"])
    fig.update_yaxes(visible=False, secondary_y=True)
    #  Change grid color
    fig.update_xaxes(showline=True, linewidth=1, linecolor=palette["grid_color"], gridcolor=palette["grid_color"])
    fig.update_yaxes(showline=True, linewidth=1, linecolor=palette["grid_color"], gridcolor=palette["grid_color"])
    file_name = f"{i}_{symbol}_grid_trading_1.png"
    fig.write_image(file_name, format="png")
    return fig


class GridStrategy(Strategy):
    chop_len = 14
    atr_len = 1
    num_grid_lines = 5  # number of grid lines for buy/sell
    grid_interval = 10 / 10000  # 10 pips, 50 pips, or 100 pips or whatever
    take_profit_interval = 20 / 10000  # pips
    stop_loss_interval = 10 / 10000  # pips
    buy_grid_prices = []
    sell_grid_prices = []
    executed_buy_grid_prices = []
    executed_sell_grid_prices = []
    last_purchase_price = 0
    long_hold = 0
    short_hold = 0
    buy_stop_loss_price = 0
    sell_stop_loss_price = 0
    grid_in_progress = False
    grid_start_index = 0  # time index when grid starts
    grid_max_interval = 2000  # max time steps to run the grid
    i = 0

    def init(self):
        super().init()
        #  Calculate indicators
        self.chop = self.I(CHOP, self.data.df, self.chop_len, self.atr_len)

    def reset_grid(self):
        self.grid_in_progress = False
        self.buy_grid_prices = []
        self.sell_grid_prices = []
        self.grid_start_index = 0
        self.buy_stop_loss_price = 0
        self.sell_stop_loss_price = 0

    def next(self):
        super().init()
        self.i += 1
        #  Check ranging or trending markets
        is_ranging = False
        if self.chop[-1] > 50 and self.chop[-2] <= 50:
            is_ranging = True
        #  Set up new grid for ranging -> against the trend
        current_price = self.data.Close[-1]
        if not self.grid_in_progress and is_ranging:
            self.reset_grid()
            self.grid_in_progress = True
            self.grid_start_index = self.i
            #  Stop loss
            buy_stop_loss = current_price - (self.num_grid_lines * self.grid_interval) - self.stop_loss_interval
            sell_stop_loss = current_price + (self.num_grid_lines * self.grid_interval) + self.stop_loss_interval
            #  Set buy/sell grid prices
            for i in range(1, self.num_grid_lines + 1):
                #  Calculate buy grid price
                grid_buy_price = current_price - (i * self.grid_interval)
                buy_take_profit = grid_buy_price + self.take_profit_interval
                self.buy_grid_prices.append(grid_buy_price)
                #  Create buy order
                self.buy(size=0.1, limit=grid_buy_price, sl=buy_stop_loss, tp=buy_take_profit)
                #  Calculate sell grid price
                grid_sell_price = current_price + (i * self.grid_interval)
                sell_take_profit = grid_sell_price - self.take_profit_interval
                self.sell_grid_prices.append(grid_sell_price)
                #  Create sell order
                self.sell(size=0.1, limit=grid_sell_price, sl=sell_stop_loss, tp=sell_take_profit)
            #  Optional - Plot the grid
            # plot_chart(self.i, symbol, df, current_price, self.buy_grid_prices, self.sell_grid_prices, buy_stop_loss, sell_stop_loss)


def run_backtest(df):
    # If exclusive orders (each new order auto-closes previous orders/position),
    # cancel all non-contingent orders and close all open trades beforehand
    bt = Backtest(df, GridStrategy, cash=10000, commission=.00075, trade_on_close=True, exclusive_orders=False,
                  hedging=False)
    stats = bt.run()
    print(stats)
    bt.plot()


# MAIN
if __name__ == '__main__':
    symbol = "EURUSD=X"
    #  Download data
    # intervals: 1m,2m,5m,15m,30m,60m,90m,1h,1d,5d,1wk,1mo,3mo
    interval = "1m"
    #  periods:  1d,5d,1mo,3mo,6mo,1y,2y,5y,10y,ytd,max
    data = yf.download(tickers=symbol, period="5d", interval=interval)
    df = pd.DataFrame(data)
    df.dropna(inplace=True)
    df.reset_index(inplace=True)
    #  Run backtest
    run_backtest(df)
