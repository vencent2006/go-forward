import unittest
from okx_grid.raw_grid import *


class MyTestCase(unittest.TestCase):
    def test_get_latest_price(self):
        res = get_latest_price()
        print(res)
        # self.assertEqual(True, False)  # add assertion here


if __name__ == '__main__':
    unittest.main()
