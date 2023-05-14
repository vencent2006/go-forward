import requests
import time
from constant import *

# API参数
base_currency = 'BTC'
quote_currency = 'USDT'
grid_size = 10  # 网格数量
grid_interval = 1000  # 网格间隔（单位：USDT）
buy_amount = 0.001  # 买入数量（单位：BTC）
initial_balance = 10000  # 初始本金

# API请求地址
base_url = 'https://www.okex.com'


# 创建订单
def create_order(side, price, size):
    timestamp = str(int(time.time() * 1000))
    method = 'POST'
    path = '/api/v5/trade/order'
    message = timestamp + method + path
    sign = sign_message(message)

    headers = {
        'OK-ACCESS-KEY': api_key,
        'OK-ACCESS-SIGN': sign,
        'OK-ACCESS-TIMESTAMP': timestamp,
        'OK-ACCESS-PASSPHRASE': passphrase,
        'Content-Type': 'application/json',
        'x-simulated-trading': flag
    }

    data = {
        'instId': base_currency + '-' + quote_currency,
        'tdMode': 'cash',
        'side': side,
        'ordType': 'limit',
        'px': str(price),
        'sz': str(size)
    }

    response = requests.post(base_url + path, headers=headers, json=data)
    return response.json()


# 签名请求消息
def sign_message(message):
    import hmac
    import base64
    import hashlib

    key = base64.b64decode(secret_key)
    message = message.encode('utf-8')
    sign = hmac.new(key, message, hashlib.sha256).digest()
    sign = base64.b64encode(sign).decode('utf-8')
    return sign


# 获取最新行情价格
def get_latest_price():
    path = '/api/v5/market/ticker'
    params = {
        'instId': base_currency + '-' + quote_currency
    }
    response = requests.get(base_url + path, params=params)
    response_json = response.json()
    ticker = response_json['data'][0]
    return float(ticker['last'])


# 计算收益率
def calculate_return(balance):
    return (balance - initial_balance) / initial_balance


# 网格交易主函数
def grid_trading():
    current_price = get_latest_price()
    # “//”是一个算术运算符，表示整数除法，它可以返回商的整数部分（向下取整）
    buy_price = current_price - (grid_size // 2) * grid_interval
    sell_price = current_price + (grid_size // 2) * grid_interval
    balance = initial_balance

    for i in range(grid_size):
        create_order('buy', buy_price, buy_amount)
        create_order('sell', sell_price, buy_amount)
        buy_price += grid_interval
        sell_price += grid_interval

        # 休眠2秒，避免频繁请求API
        time.sleep(2)

        current_price = get_latest_price()
        print('Current Price: ', current_price)
        print('Buy Price: ', buy_price)
        print('Sell Price: ', sell_price)
        print('---')

        # 更新本金
        balance -= buy_price * buy_amount

    # 计算收益率
    final_balance = balance + current_price * (grid_size // 2) * buy_amount
    return_rate = calculate_return(final_balance)

    print('Final Balance: ', final_balance)
    print('Return Rate: ', return_rate)


if __name__ == '__main__':
    # 运行网格交易策略
    grid_trading()
