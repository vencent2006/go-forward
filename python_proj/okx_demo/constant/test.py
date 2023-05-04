import json

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

# order
ORDER_TYPE_LIMIT = "limit"
ORDER_TYPE_MARKET = "market"


def print_pretty_json(jsonStr):
    print(json.dumps(jsonStr, indent=2))
