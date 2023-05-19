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


def get_history_klines(instId, bar="1m", start='', end=''):
    """
    获取交易产品历史K线数据
    https://www.okx.com/docs-v5/zh/#rest-api-market-data-get-candlesticks-history
    :param instId: 交易对
    :param start: 开始时间
    :param end: 结束时间
    :param bar: 单位
    :return:返回值数组顺序分别为是：[ts,o,h,l,c,vol,volCcy,confirm]
    {
      "code": "0",
      "msg": "",
      "data": [
        [
          "1684126860000",#开始时间，Unix时间戳的毫秒数格式，如 1597026383085
          "27305.4",#开盘价格 open
          "27312.2",#最高价，high
          "27283.7",#最低价，low
          "27283.7",#收盘价, close
          "5178.2846985", # 交易量，以张为单位
          "141385301.776316512",#交易量，以币为单位
          "141385301.776316512",# 交易量，以计价货币为单位
          "1" #K线状态 0 代表 K 线未完结，1 代表 K 线已完结
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
    result = marketDataAPI.get_history_candlesticks(instId=instId, bar=bar, before=start, after=end)
    print_pretty_json(result)
    return result


if __name__ == '__main__':
    get_ticker("BTC-USDT")
