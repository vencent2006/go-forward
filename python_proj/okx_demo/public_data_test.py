import unittest
from public_data import *


class MyTestCase(unittest.TestCase):

    def test_get_instruments(self):
        res = get_instruments()
        # 返回的是一个字符串的"0"
        self.assertEqual(res["code"], "0")


if __name__ == '__main__':
    unittest.main()
