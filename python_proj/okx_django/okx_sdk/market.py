import okx.MarketData as MarketData
from constant import *
from utils import *

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
    print_pretty_json(result)
    return result


def get_tickers():
    """
    获取现货报价
    :return: 列表
    """
    result = marketDataAPI.get_tickers(instType=INST_TYPE_SPOT)
    print_pretty_json(result)
    return result


def get_klines(instId):
    """
    获取交易产品K线数据
    https://www.okx.com/docs-v5/zh/#rest-api-market-data-get-candlesticks
    :param instId: 交易对
    :return:
{
  "code": "0",
  "msg": "",
  "data": [
    [
      "1684123440000",
      "27205.7",
      "27212.2",
      "27205.7",
      "27212.2",
      "1870.55037703",
      "50893212.573010258",
      "50893212.573010258",
      "0"
    ],
    [
      "1684123380000",
      "27212.3",
      "27221.9",
      "27204",
      "27205.9",
      "5736.92466366",
      "156115779.520956775",
      "156115779.520956775",
      "1"
    ]
  ]
}

    """
    result = marketDataAPI.get_candlesticks(instId=instId)
    print_pretty_json(result)
    return result


def get_history_klines(instId):
    """
    获取交易产品历史K线数据
    https://www.okx.com/docs-v5/zh/#rest-api-market-data-get-candlesticks-history
    :param instId: 交易对
    :return:
    {
      "code": "0",
      "msg": "",
      "data": [
        [
          "1684126860000",
          "27305.4",
          "27312.2",
          "27283.7",
          "27283.7",
          "5178.2846985",
          "141385301.776316512",
          "141385301.776316512",
          "1"
        ],
        [
          "1684126800000",
          "27286.8",
          "27363.4",
          "27286.8",
          "27307.6",
          "15435.68443124",
          "421676366.313617402",
          "421676366.313617402",
          "1"
        ],
        [
          "1684126740000",
          "27271.4",
          "27291.9",
          "27271.4",
          "27287.9",
          "5129.96041458",
          "139933132.23992391",
          "139933132.23992391",
          "1"
        ]
      ]
    }

    """
    result = marketDataAPI.get_history_candlesticks(instId=instId)
    print_pretty_json(result)
    return result


if __name__ == '__main__':
    get_ticker("BTC-USDT")
