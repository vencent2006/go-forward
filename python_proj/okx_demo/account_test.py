import unittest
from account import *


class MyTestCase(unittest.TestCase):

    def test_get_all_account_balance(self):
        res = get_all_account_balance()
        # 返回的是一个字符串的"0"
        self.assertEqual(res["code"], "0")

    def test_get_one_account_balance(self):
        ccy = "BTC"
        res = get_one_account_balance(ccy)
        # 返回的是一个字符串的"0"
        self.assertEqual(res["code"], "0")


if __name__ == '__main__':
    unittest.main()
