import urllib3
import re
import pandas as pd
import numpy as np

# 访问行业板块数据
http = urllib3.PoolManager();

pages = 4
conts = []
for p in range(1, pages + 1):
    url = "http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?cb=jQuery1124012582582823807198_1554554782636&type=CT&token=4f1862fc3b5e77c150a2b985b12db0fd&sty=FPGBKI&js=({data:[(x)],recordsFiltered:(tot)})&cmd=C._BKHY&st=(ChangePercent)&sr=-1&p=%d" % p
    url += "&ps=20&_=1554554783027"
    try:
        resp_dat = http.request('GET', url)
        pattern = re.compile(r'BK(.*?)"')
        bk_list = re.findall(pattern, resp_dat.data.decode())
        for bk in bk_list:
            conts.append(bk)
        print(resp_dat.data.decode())
    except Exception as e:
        print(resp_dat.status)
        print(e)

print(conts)

df = pd.DataFrame(np.zeros((len(conts), 7)),
                  columns=[u'板块名称', u'BK涨跌幅', u'总市值', u'换手率', u'涨跌家数', u'领涨股票', u'SK涨跌幅'])
i = 1
for num, bk_dat in enumerate(conts):
    print(bk_dat)

    bk_dat = bk_dat.split(',')
    df.loc[df.index[num], u'板块名称'] = bk_dat[1]
    df.loc[df.index[num], u'BK涨跌幅'] = bk_dat[2]
    df.loc[df.index[num], u'总市值'] = bk_dat[3]
    df.loc[df.index[num], u'换手率'] = bk_dat[4]
    df.loc[df.index[num], u'涨跌家数'] = bk_dat[5]
    df.loc[df.index[num], u'领涨股票'] = bk_dat[8]
    df.loc[df.index[num], u'SK涨跌幅'] = bk_dat[10]
    print(df)
    i = i+1
    if i == 3:
        exit()
print(df)
