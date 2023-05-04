import unittest
from market import *


def my_sum(a, b):
    return a + b


class MyTestCase(unittest.TestCase):
    def test_demo1(self):
        res = my_sum(1, 2)
        self.assertEqual(res, 3)  # add assertion here

    def test_get_ticker(self):
        instId = "BTC-USDT"
        res = get_ticker("BTC-USDT")
        self.assertEqual(res["data"][0]["instType"], "SPOT")
        self.assertEqual(res["data"][0]["instId"], "BTC-USDT")


if __name__ == '__main__':
    unittest.main()
