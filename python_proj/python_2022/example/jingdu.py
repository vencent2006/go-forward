# coding: utf-8

import time
import sys  # 引入 sys 模块
from timeit import Timer

Timer('t=a; a=b; b=t', 'a=1; b=2').timeit()
