# coding:utf-8
def countdown(n):
    print('count down from ', n)
    while n >= 0:
        print('n: ', n)
        newvalue = (yield n)
        if newvalue is not None:
            pass

# if __name__ == '__main__':
#     double = double_money()
#     double.send(None)
#     print('first')
#     print(double.send(10))
#     print('second')
#     print(double.send(20))
#     print('third')
#     print(double.send(30))
