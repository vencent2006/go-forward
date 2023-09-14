# coding:utf-8


import time
import calendar

ticks = time.time()
print(type(ticks))
print("current time is", ticks)

localtime = time.localtime(ticks)
print(localtime)

cal = calendar.month(2016, 1)
print("以下输出2016年1月份的日历:")
print(cal)
