# coding:utf-8
def double_money():
    print('输入你的金钱，我来负责乘以2，机会仅限三次')
    # 需要先执行一次 send(None) 到这里
    x = yield
    x = yield x * 2
    x = yield x * 2
    x = yield x * 2


if __name__ == '__main__':
    double = double_money()
    double.send(None)
    print('first')
    print(double.send(10))
    print('second')
    print(double.send(20))
    print('third')
    print(double.send(30))
