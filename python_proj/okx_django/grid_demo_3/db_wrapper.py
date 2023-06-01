# 导包
from utils.mysql_util import Database
from utils.time_util import get_current_second

# 设置连接数据库的参数
config = {
    "host": "127.0.0.1",
    "port": 3307,
    "database": "quant",
    "charset": "utf8",
    "user": "root",
    "passwd": "root"
}

# 实例化时就直接传参数
db = Database(**config)

# table
tb_order_history = 'order_history'
tb_ticker_history = 'ticker_history'

side_buy = 'buy'
side_sell = 'sell'


def get_trade_history(instId: str, side: str) -> []:
    """
    获取所有交易记录
    :param instId: 交易对
    :param side: buy or sell
    :return:
    """
    cond = f"instId = {instId} and side = {side}"
    return db.select_all(table_name=tb_order_history, factor_str=cond)


def insert_trade_history(instId: str, side: str, price: float, size: float, money: float, client_order_id: str,
                         exchange_order_id: str):
    """
    记录订单交易历史
    :param instId: 交易对
    :param side: buy or sell
    :param price: 价格
    :param size: 数量
    :param money: 金钱
    :param client_order_id: 业务订单id
    :param exchange_order_id: 交易所订单id
    :return:
    """
    if side != side_buy and side != side_sell:
        raise NameError(f"invalid side({side})")
    ts = get_current_second()
    vals = (instId, side, price, size, money, client_order_id, exchange_order_id, ts)
    db.insert(table_name=tb_order_history, value=vals)


def update_order(instId: str, client_order_id: str, status: str):
    """
    更新订单信息
    :param instId: 交易对
    :param client_order_id: 业务订单id
    :param status: 订单状态
    :return:
    """
    cond = f"instId = {instId} and client_order_id = {client_order_id}"
    vals = {'status': status, 'update_time': get_current_second()}
    db.update(tb_ticker_history, vals, cond)


def get_ticker_history(instId: str) -> []:
    """
    获取行情信息
    :param instId: 交易对
    :return:
    """
    return db.select_all(table_name=tb_ticker_history)


def insert_ticker(instId: str, close: str):
    ts = get_current_second()
    vals = (instId, close, ts)
    db.insert(table_name=tb_ticker_history, value=vals)
