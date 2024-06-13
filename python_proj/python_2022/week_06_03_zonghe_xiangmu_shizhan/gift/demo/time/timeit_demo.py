# coding:utf-8

from timeit import Timer

a = Timer('t=a;a=b;b=t', 'a=1;b=2').timeit()
print(a)

a = Timer('a,b=b,a', 'a=1;b=2').timeit()
print(a)
