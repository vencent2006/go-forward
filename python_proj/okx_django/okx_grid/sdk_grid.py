from okx_sdk import *

# API参数
base_currency = 'BTC'
quote_currency = 'USDT'
instId = base_currency + '-' + quote_currency
grid_size = 10  # 网格数量
grid_interval = 1000  # 网格间隔（单位：USDT）
buy_amount = 0.001  # 买入数量（单位：BTC）
initial_balance = 10000  # 初始本金


# 创建订单
def create_order(side, price, size):
    if side == 'buy':
        return buy_limit_order(instId, price, size)
    else:
        return sell_limit_order(instId, price, size)


# 获取最新行情价格
def get_latest_price():
    ticker = get_ticker(instId)
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
        # todo 如何处理滑点: buy 不到; 或者 sell 不出去
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
