import logging

import okx.Trade as Trade
from constant import *
from utils import *

tradeAPI = Trade.TradeAPI(api_key, secret_key, passphrase, False, flag)


def buy_limit_order(instId: str, price, size, clientOrderId=""):
    """
    下单:买入限价单
    https://www.okx.com/docs-v5/zh/#rest-api-trade-place-order
    :param instId: 交易对
    :param price: 价格
    :param size: 数量
    :param clientOrderId: 自定义订单id，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
    :return:
    """

    result = tradeAPI.place_order(
        instId=instId,
        tdMode=TRADE_MODE_CASH,
        side=TRADE_SIDE_BUY,
        ordType=ORDER_TYPE_LIMIT,
        px=price,
        sz=size,
        clOrdId=clientOrderId
    )
    print_pretty_json(result)

    if result["code"] == CODE_SUCCESS:
        print("Successful order request，order_id = ", result["data"][0]["ordId"])
    else:
        print("Unsuccessful order request，error_code = ", result["data"][0]["sCode"], ", Error_message = ",
              result["data"][0]["sMsg"])

    return result


# def buy_market_order(instId, size, clientOrderId):
#     """
#     下单:买入市价单
#     :param instId: 交易对
#     :param size: 数量
#     :param clientOrderId: 自定义订单id，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
#     :return:
#     """
#     result = tradeAPI.place_order(
#         instId=instId,
#         tdMode=TRADE_MODE_CASH,
#         side=TRADE_SIDE_BUY,
#         ordType=ORDER_TYPE_MARKET,
#         sz=size,
#         clOrdId=clientOrderId
#     )
#
#     print_pretty_json(result)
#
#     if result["code"] == CODE_SUCCESS:
#         print("Successful order request，order_id = ", result["data"][0]["ordId"])
#     else:
#         print("Unsuccessful order request，error_code = ", result["data"][0]["sCode"], ", Error_message = ",
#               result["data"][0]["sMsg"])
#
#     return result


def sell_limit_order(instId: str, price, size, clientOrderId=""):
    """
    下单:买入限价单
    https://www.okx.com/docs-v5/zh/#rest-api-trade-place-order
    :param instId: 交易对
    :param price: 价格
    :param size: 数量
    :param clientOrderId: 自定义订单id，字母（区分大小写）与数字的组合，可以是纯字母、纯数字且长度要在1-32位之间
    :return:
    """

    result = tradeAPI.place_order(
        instId=instId,
        tdMode=TRADE_MODE_CASH,
        side=TRADE_SIDE_SELL,
        ordType=ORDER_TYPE_LIMIT,
        px=price,
        sz=size,
        clOrdId=clientOrderId
    )
    print_pretty_json(result)

    if result["code"] == CODE_SUCCESS:
        print("Successful order request，order_id = ", result["data"][0]["ordId"])
    else:
        print("Unsuccessful order request，error_code = ", result["data"][0]["sCode"], ", Error_message = ",
              result["data"][0]["sMsg"])

    return result
