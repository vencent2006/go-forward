# 导包
from constant.key import db_conf
from utils.mysql_util import Database
from utils.time_util import get_current_second

# 实例化时就直接传参数
db = Database(**db_conf)

# table
tb_order_history = 'trade_order'
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
    cond = f"market = '{instId}' and side = '{side}'"
    print(cond)
    return db.select_all(table_name=tb_order_history, factor_str=cond)


def insert_trade(instId: str, side: str, price: float, size: float, client_order_id: str, channel_order_id: str):
    """
    记录订单交易历史
    :param instId: 交易对
    :param side: buy or sell
    :param price: 价格
    :param size: 数量
    :param client_order_id: 业务订单id
    :param channel_order_id: 交易所订单id
    :return:
    """
    if side != side_buy and side != side_sell:
        raise NameError(f"invalid side({side})")
    ts = get_current_second()
    """
CREATE TABLE `trade_order` (
  `trade_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `side` varchar(10) NOT NULL COMMENT 'buy/sell',
  `market` varchar(50) DEFAULT NULL COMMENT '交易对',
  `status` int(11) NOT NULL COMMENT '状态',
  `channel_name` varchar(10) DEFAULT NULL COMMENT '交易所名称',
  `amount` decimal(40,20) DEFAULT NULL COMMENT '数量',
  `price` decimal(40,20) DEFAULT NULL COMMENT '价格',
  `uniq_id` varchar(100) DEFAULT NULL COMMENT '业务订单id',
  `channel_filled_amount` decimal(40,20) NOT NULL DEFAULT '0.00000000000000000000',
  `channel_avg_filled_price` decimal(40,20) NOT NULL DEFAULT '0.00000000000000000000',
  `channel_order_id` bigint(20) DEFAULT NULL,
  `channel_order_create_time` datetime DEFAULT NULL,
  `status_note` varchar(250) DEFAULT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`trade_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
  """
    market = instId
    status = 0
    channel_name = 'okx'
    amount = size
    uniq_id = client_order_id
    # channel_filled_amount = 0.0
    # channel_avg_filled_price = 0.0
    # channel_order_create_time = 0
    # status_note = ''

    columns = "side, market, status, channel_name, amount, price, uniq_id, channel_order_id"
    vals = (side, market, status, channel_name, amount, price, uniq_id, channel_order_id)

    db.insert(table_name=tb_order_history, columns=columns, value=vals)


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
