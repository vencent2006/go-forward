class History:
    """买入操作历史类"""
    # 操作 1 买入 -1 卖出
    stock_code = ""
    count = 0
    price = 0.0
    opt_type = 0

    def __init__(self, stock_code, opt_type, price, count):
        self.stock_code = stock_code
        self.opt_type = opt_type
        self.price = price
        self.count = count
