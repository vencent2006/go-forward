import unittest
from okx_sdk.market import *


class MyTestCase(unittest.TestCase):

    def test_get_ticker(self):
        instId = "BTC-USDT"
        res = get_ticker(instId)
        self.assertEqual(res["code"], CODE_SUCCESS)
        self.assertEqual(res["data"][0]["instType"], INST_TYPE_SPOT)
        self.assertEqual(res["data"][0]["instId"], instId)

    def test_get_tickers(self):
        res = get_tickers()
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_klines(self):
        instId = "BTC-USDT"
        res = get_klines(instId)
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_history_klines(self):
        instId = "BTC-USDT"
        res = get_history_klines(instId)
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_history_klines_start_end(self):
        instId = "BTC-USDT"
        start = "1669824000000"  # 2022-12-01 00:00:00
        end = "1672502400000"  # 2023-01-01 00:00:00
        res = get_history_klines(instId, bar="1Dutc", start=start, end=end)
        self.assertEqual(res["code"], CODE_SUCCESS)

    def test_get_lastest_close(self):
        instId = "BTC-USDT"
        res = get_ticker(instId)
        self.assertEqual(res["code"], CODE_SUCCESS)
        self.assertEqual(res["data"][0]["instType"], INST_TYPE_SPOT)
        self.assertEqual(res["data"][0]["instId"], instId)
        """
        {
          "code": "0",
          "msg": "",
          "data": [
            {
              "instType": "SPOT",
              "instId": "BTC-USDT",
              "last": "27145.9",
              "lastSz": "158.13880974",
              "askPx": "27147.5",
              "askSz": "118.87250831",
              "bidPx": "27147.4",
              "bidSz": "90.57158482",
              "open24h": "27011.5",
              "high24h": "27495.3",
              "low24h": "26980.7",
              "volCcy24h": "80250621069.41386445",
              "vol24h": "2942226.83285096",
              "ts": "1684893330510",
              "sodUtc0": "27225.3",
              "sodUtc8": "27324.5"
            }
        ]
        """
        timestamp = float(float(res['data'][0]['ts']) / 1000)
        time_array = time.localtime(timestamp)
        ts = time.strftime("%Y-%m-%d %H:%M:%S", time_array)
        close = float(res['data'][0]['last'])
        print({'timestamp': ts, 'close': close})


if __name__ == '__main__':
    unittest.main()
