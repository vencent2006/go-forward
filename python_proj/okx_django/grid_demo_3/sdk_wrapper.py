import time
from datetime import datetime

import constant
import okx_sdk
import utils


def get_latest_close(instId: str):
    """
    返回最新的close结构{'timestamp':ts, 'close': close}
    :param instId:
    :return: fail, return False; success, {'timestamp':ts, 'close': close}
    """
    res = okx_sdk.get_ticker(instId)
    if res["code"] != '0':
        return False
    else:
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
        res = {'ts': ts, 'close': close, 'local_time': datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')[:-3]}
        print('get_ticker', instId, res)
        return res


def get_balance(instId: str) -> {}:
    """
    获取余额
    :param instId: 交易对
    :return: {'avail_balance': avail_balance, 'cash_balance': cash_balance, 'order_frozen': order_frozen}
    """
    ccy = 'USDT'
    res = okx_sdk.get_one_account_balance(ccy=ccy)
    if res["code"] != '0':
        return False
    else:
        detail = res['data'][0]['details'][0]
        avail_balance = detail['availBal']  # 可用余额
        cash_balance = detail['cashBal']  # 可用余额
        order_frozen = detail['ordFrozen']  # 订单冻结
        return {'avail_balance': avail_balance, 'cash_balance': cash_balance, 'order_frozen': order_frozen}


def gen_client_order_id():
    return utils.get_uuid()


def try_buy(instId: str, price, size, client_order_id):
    """
    尝试着去buy, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """
    return place_order('buy', instId=instId, price=price, size=size, client_order_id=client_order_id)


def try_sell(instId: str, price, size, client_order_id):
    """
    尝试着去sell, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """
    return place_order('sell', instId=instId, price=price, size=size, client_order_id=client_order_id)


def place_order(side: str, instId: str, price, size, client_order_id):
    """
    尝试着去buy, 是要拿到确切结果的
    :return: successful, exchange_order_id; failed, False
    """

    # 1. 下单，并获取下单结果
    if side == 'buy':
        res = okx_sdk.buy_limit_order(instId=instId, price=price, size=size, clientOrderId=client_order_id)
    elif side == 'sell':
        res = okx_sdk.sell_limit_order(instId=instId, price=price, size=size, clientOrderId=client_order_id)
    else:
        raise NameError("invalid side(%s)" % side)

    # 判定结果
    if res["code"] != '0':
        # sdk返回错误
        print('buy_limit_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return code', res['code'], 'return msg', res['msg'])
        return False
    elif res['data'][0]['sCode'] != '0':
        # sdk返回错误
        print('buy_limit_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return sCode', res['data'][0]['sCode'], 'return sMsg', res['data'][0]['sMsg'])
        return False

    # place buy limit order, successfully
    exchange_order_id = res['data'][0]['ordId']
    print('buy_limit_order successfully |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
          client_order_id, 'return exchange_order_id', exchange_order_id)

    # 2. 每隔5s，获取订单结果，最长等待10 * interval，
    time_sleep_cnt = 0
    interval = 5
    max_time_sleep_cnt = 10 * interval
    while time_sleep_cnt < max_time_sleep_cnt:
        print("****** 监听订单交易结果, time_sleep_cnt", time_sleep_cnt, "max_time_sleep_cnt", max_time_sleep_cnt)
        time.sleep(interval)
        time_sleep_cnt += interval
        res = okx_sdk.get_order_by_client_order(instId, client_order_id)
        if res["code"] != '0':
            # sdk 返回异常
            print('get_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
            continue
        else:
            data = res['data'][0]
            state = data['state']
            if state == constant.ORDER_STATUS_FILLED:
                # 已经成交，终态
                return exchange_order_id
            elif state == constant.ORDER_STATUS_CANCELED:
                # 已经取消，终态
                return False
            else:
                # 非终态，继续查询
                continue

    # 如果还没交易成功，那么就取消订单，并返回false
    res = okx_sdk.cancel_order_by_client_order(instId=instId, client_order_id=client_order_id)
    if res["code"] != '0':
        print('cancel_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
    elif res['data'][0]['sCode'] != '0':
        # sdk返回错误
        print('cancel_order_by_client_order failed |', 'instId', instId, 'price', price, 'size', size, 'clientOrderId',
              client_order_id, 'return sCode', res['data'][0]['sCode'], 'return sMsg', res['data'][0]['sMsg'])
        return False

    # 取消可能成功, 再查询下
    res = okx_sdk.get_order_by_client_order(instId=instId, client_order_id=client_order_id)
    if res["code"] != '0':
        # sdk 返回异常
        print('after cancel order | get_order_by_client_order failed | code: %s | msg: %s' % (res['code'], res['msg']))
    else:
        data = res['data'][0]
        state = data['state']
        if state == constant.ORDER_STATUS_FILLED:
            # 已经成交，终态
            return exchange_order_id
        elif state == constant.ORDER_STATUS_CANCELED:
            # 已经取消，终态
            return False
        else:
            # 非终态，继续查询
            print(
                'after cancel order | invalid state | get_order_by_client_order return state = %s(not final state)' % state)
            return False
