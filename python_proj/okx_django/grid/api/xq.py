# 雪球
import json
import time

import requests


def get_cookies(stock_code):
    # https://xueqiu.com/S/%s
    # 需要先请求页面获取cookie
    url_str = "https://xueqiu.com/S/%s" % stock_code
    headers = {
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
    }
    resp = requests.request("GET", url_str, headers=headers)
    if resp.status_code == 200:
        return resp.cookies
    else:
        print(resp, resp.text)


# "1m"
# "5m"
# "15m"
# "60m"
# "120m"
# "day"
# "week"
# "quarter"
# "year"
# 日期 开盘价 收盘价 最高价 最低价
def get_data(stock_code, period):
    cookies = get_cookies(stock_code)
    print(stock_code, period)
    # symbol=00700&begin=1622598594118&period=day&type=before&count=-284&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance
    #                    1622513284350
    begin = int(time.time() * 1000)
    time_type = "before"
    # time_type = "normal"
    count = 2000000000
    url_string = "https://stock.xueqiu.com/v5/stock/chart/kline.json?symbol=%s&begin=%d&period=%s&type=%s&count=%d" \
                 "&indicator=%s" % (
                     stock_code, begin, period, time_type, -count, "kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance")
    print(url_string)

    headers = {
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
    }

    resp = requests.request("GET", url_string, headers=headers, cookies=cookies)
    resp.encoding = 'utf-8'
    if resp.status_code == 200:
        data = json.loads(resp.text)
        # 06 = {str}
        # 涨跌额 'chg'
        # 07 = {str}
        # 涨跌幅 'percent'
        # 08 = {str}
        # 'turnoverrate'
        # 09 = {str}
        # 'amount'
        # 10 = {str}
        # 'volume_post'
        # 11 = {str}
        # 'amount_post'
        # 12 = {str}
        # 'pe'
        # 13 = {str}
        # 'pb'
        # 14 = {str}
        # 'ps'
        # 15 = {str}
        # 'pcf'
        # 16 = {str}
        # 'market_capital'
        # 17 = {str}
        # 'balance'
        # 18 = {str}
        # 'hold_volume_cn'
        # 19 = {str}
        # 'hold_ratio_cn'
        # 20 = {str}
        # 'net_volume_cn'
        # 21 = {str}
        # 'hold_volume_hk'
        # 22 = {str}
        # 'hold_ratio_hk'
        # 23 = {str}
        # 'net_volume_hk'
        return data["data"]["item"]
    else:
        print(resp)
