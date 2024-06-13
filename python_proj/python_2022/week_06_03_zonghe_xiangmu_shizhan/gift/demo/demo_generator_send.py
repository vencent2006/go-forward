# coding:utf-8
def countdown(n):
    print('count down from ', n)
    while n >= 0:
        print('n: ', n)
        newvalue = (yield n)
        if newvalue is not None:
            n = newvalue
        else:
            n -= 1


if __name__ == '__main__':
    c = countdown(5)
    for x in c:
        print('x =', x)
        if x == 5:
            print('++++++')
            c.send(3)
            print('------')
