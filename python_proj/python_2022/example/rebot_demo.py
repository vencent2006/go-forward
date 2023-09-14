# coding:utf-8

import urllib.robotparser

rp = urllib.robotparser.RobotFileParser()
rp.set_url("https://www.weibo.com/robots.txt")
rp.read()
rrate = rp.request_rate("*")
print(rrate)
