import unittest
from trade import *
from utils import *


class MyTestCase(unittest.TestCase):
    def test_place_buy_limit_order(self):
        instId = "BTC-USDT"
        price = "19000"
        size = "0.01"
        clientOrderId = "myorder" + get_current_second(to_str=True)
        res = place_buy_limit_order(instId, price, size, clientOrderId)
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_place_buy_market_order(self):
        instId = "BTC-USDT"
        size = "0.01"
        clientOrderId = "myorder" + get_current_second(to_str=True)
        res = place_buy_market_order(instId, size, clientOrderId)
        self.assertEqual(res["code"], CODE_SUCCESS)


if __name__ == '__main__':
    unittest.main()
