# coding:utf-8
class CounterList(list):
    def __init__(self, *args):
        super().__init__(*args)
        self.counter = 0

    def __getitem__(self, index):
        self.counter += 1
        return super().__getitem__(index)


if __name__ == '__main__':
    cl = CounterList(range(10))
    print(cl)
    print(cl.counter)
    cl.reverse()
    print(cl.counter)
    del cl[3:6]
    print(cl)
    print(cl.counter)
    print(cl[2] + cl[4])
    print(cl.counter)
