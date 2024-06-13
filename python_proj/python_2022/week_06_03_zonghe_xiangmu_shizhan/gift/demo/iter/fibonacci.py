# coding:utf-8

def fibonacci(n):
    a, b, counter = 0, 1, 1
    while True:
        if counter > n:
            return
        yield a
        a, b = b, a + b
        counter += 1


f = fibonacci(10)
for x in f:
    print(x)
