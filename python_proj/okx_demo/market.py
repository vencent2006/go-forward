import okx.MarketData as MarketData
import json
from constant import *

# 主要参考这个 https://www.okx.com/learn/spot-trading-with-jupyter-notebook

marketDataAPI = MarketData.MarketAPI(flag=flag)


def get_ticker(instId):
    """
    获取指定交易对的报价行情
    :param instId: 交易对
    :return:
    {
      "code": "0",
      "msg": "",
      "data": [
        {
          "instType": "SPOT",
          "instId": "BTC-USDT",
          "last": "28587.4",
          "lastSz": "3",
          "askPx": "28587.5",
          "askSz": "102.61139811",
          "bidPx": "28587.4",
          "bidSz": "104.75868491",
          "open24h": "29250.6",
          "high24h": "29963.1",
          "low24h": "28283",
          "volCcy24h": "174348873161.777148756",
          "vol24h": "5973302.77732658",
          "ts": "1682933882503",
          "sodUtc0": "29240.1",
          "sodUtc8": "29661.9"
        }
      ]
    }
    """
    result = marketDataAPI.get_ticker(instId=instId)
    print(json.dumps(result, indent=2))
    return result


def get_tickers():
    """
    获取现货报价
    :return: 列表
    """
    result = marketDataAPI.get_tickers(instType="SPOT")
    print(json.dumps(result, indent=2))
    return result


if __name__ == '__main__':
    get_ticker("BTC-USDT")
