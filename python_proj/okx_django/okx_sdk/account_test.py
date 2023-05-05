import unittest
from okx_sdk.account import *


class MyTestCase(unittest.TestCase):

    def test_get_all_account_balance(self):
        res = get_all_account_balance()
        # 返回的是一个字符串的"0"
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_one_account_balance(self):
        ccy = "BTC"
        res = get_one_account_balance(ccy)
        # 返回的是一个字符串的"0"
        self.assertEqual(res["code"], CODE_SUCCESS)


if __name__ == '__main__':
    unittest.main()
