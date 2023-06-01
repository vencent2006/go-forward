import unittest

from grid_demo_3.db_wrapper import *


class MyTestCase(unittest.TestCase):

    def test_get_trade_history(self):
        instId = "BTC-USDT"
        side = side_buy
        res = get_trade_history(instId, side)
        print(res)

    def test_insert_trade(self):
        instId = "BTC-USDT"
        side = side_buy
        price = 1.0
        size = 100
        client_order_id = "client_oid_123"
        channel_order_id = 456

        res = insert_trade(instId=instId, side=side, price=price, size=size, client_order_id=client_order_id,
                           channel_order_id=channel_order_id)
        print(res)


if __name__ == '__main__':
    unittest.main()
