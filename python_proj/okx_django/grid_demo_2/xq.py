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


def get_data(stock_code, period):
    cookies = get_cookies(stock_code)
    print(stock_code, period)

    begin = int(time.time() * 1000)
    time_type = "before"
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
        return data["data"]["item"]
    else:
        print(resp)
