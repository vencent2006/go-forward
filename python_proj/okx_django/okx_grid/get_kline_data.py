from okx_sdk.market import get_klines
from datetime import datetime
from pytz import timezone


class BarData:
    def __init__(self, instId, datetime, interval, open_price, high_price, low_price, close_price, volume,
                 gateway_name):
        self.instId = instId
        self.datetime = datetime
        self.interval = interval
        self.open_price = open_price
        self.high_price = high_price
        self.low_price = low_price
        self.close_price = close_price
        self.volume = volume
        self.gateway_name = gateway_name

    def __str__(self):
        return str(self.__class__) + ":" + str(self.__dict__)

    def __repr__(self):
        return str(self.__class__) + ":" + str(self.__dict__)


if __name__ == '__main__':
    CHINA_TZ = timezone("Asia/Shanghai")
    instId = "BTC-USDT"
    res = get_klines(instId)
    if res["code"] != "0":
        print(res)
    bar_data = []
    for rd in res["data"]:
        dt = datetime.fromtimestamp(int(rd[0]) / 1000)
        # [ts,o,h,l,c,vol,volCcy,confirm]
        bar = BarData(
            instId=instId,
            datetime=CHINA_TZ.localize(dt),
            interval="m",
            open_price=float(rd[1]),
            high_price=float(rd[2]),
            low_price=float(rd[3]),
            close_price=float(rd[4]),
            volume=float(rd[5]),
            gateway_name="okex"
        )
        print(rd)
        bar_data.append(bar)
    print(bar_data[0])
