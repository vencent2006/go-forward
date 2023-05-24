# live trading
flag = "1"  # live trading: 0, demo trading: 1

# return data
CODE_SUCCESS = "0"

# instrument
INST_TYPE_SPOT = "SPOT"

# trade
TRADE_SIDE_BUY = "buy"
TRADE_SIDE_SELL = "sell"
TRADE_MODE_CASH = "cash"

# order type
ORDER_TYPE_LIMIT = "limit"
ORDER_TYPE_MARKET = "market"

# order status
ORDER_STATUS_CANCELED = "canceled"  # 撤单成功
ORDER_STATUS_LIVE = "live"  # 等待成交
ORDER_STATUS_PART_FILLED = "partially_filled"  # 部分成交
ORDER_STATUS_FILLED = "filled"  # 完全成交

# print_json = False  # 关闭日志
print_json = True  # 开启日志
