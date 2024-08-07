import unittest
from market import *


class MyTestCase(unittest.TestCase):

    def test_get_ticker(self):
        instId = "BTC-USDT"
        res = get_ticker(instId)
        self.assertEqual(res["code"], CODE_SUCCESS)
        self.assertEqual(res["data"][0]["instType"], INST_TYPE_SPOT)
        self.assertEqual(res["data"][0]["instId"], instId)


if __name__ == '__main__':
    unittest.main()
