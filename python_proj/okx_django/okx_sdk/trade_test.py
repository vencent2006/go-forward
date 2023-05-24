import unittest

from trade import *
from utils import *


class MyTestCase(unittest.TestCase):
    def test_buy_limit_order(self):
        instId = "BTC-USDT"
        price = "19000"
        price = "26797"
        size = "0.01"
        size = "0.00001"
        clientOrderId = "myorder" + get_current_second(to_str=True)
        res = buy_limit_order(instId, price, size, clientOrderId)
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_sell_limit_order(self):
        instId = "BTC-USDT"
        price = "19000"
        size = "0.01"
        clientOrderId = "myorder" + get_current_second(to_str=True)
        res = sell_limit_order(instId, price, size, clientOrderId)
        self.assertEqual(res["code"], CODE_SUCCESS)

    # def test_place_buy_market_order(self):
    #     instId = "BTC-USDT"
    #     size = "0.01"
    #     clientOrderId = "myorder" + get_current_second(to_str=True)
    #     res = buy_market_order(instId, size, clientOrderId)
    #     self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_order_by_exchange_order(self):
        instId = "BTC-USDT"
        order_id = str(581538103723147264)
        res = get_order_by_exchange_order(instId=instId, exchange_order_id=order_id)
        self.assertEqual(res["code"], CODE_SUCCESS)
        """
        {
  "code": "0",
  "data": [
    {
      "accFillSz": "0.00001",
      "algoClOrdId": "",
      "algoId": "",
      "avgPx": "26797",
      "cTime": "1684921488383",
      "cancelSource": "",
      "cancelSourceReason": "",
      "category": "normal",
      "ccy": "",
      "clOrdId": "myorder1684921487",
      "fee": "-0.000000008",
      "feeCcy": "BTC",
      "fillPx": "26797",
      "fillSz": "0.00001",
      "fillTime": "1684921492971",
      "instId": "BTC-USDT",
      "instType": "SPOT",
      "lever": "",
      "ordId": "581538103723147264",
      "ordType": "limit",
      "pnl": "0",
      "posSide": "net",
      "px": "26797",
      "quickMgnType": "",
      "rebate": "0",
      "rebateCcy": "USDT",
      "reduceOnly": "false",
      "side": "buy",
      "slOrdPx": "",
      "slTriggerPx": "",
      "slTriggerPxType": "",
      "source": "",
      "state": "filled",
      "sz": "0.00001",
      "tag": "",
      "tdMode": "cash",
      "tgtCcy": "",
      "tpOrdPx": "",
      "tpTriggerPx": "",
      "tpTriggerPxType": "",
      "tradeId": "78462798",
      "uTime": "1684921492973"
    }
  ],
  "msg": ""
}
        """

    def test_get_order_by_client_order(self):
        # order_id =  581538103723147264
        # "clOrdId": "myorder1684921487"
        instId = "BTC-USDT"
        client_order_id = "myorder1684921487"
        res = get_order_by_client_order(instId=instId, client_order_id=client_order_id)
        self.assertEqual(res["code"], CODE_SUCCESS)


if __name__ == '__main__':
    unittest.main()
