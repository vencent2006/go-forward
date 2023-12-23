class MyNumbers:
    def __iter__(self):
        self.a = 1
        return self

    """
    说明第一次调用next，应该是设置的初始值，即1；而不是上来就是2
    """

    def __next__(self):
        x = self.a
        self.a += 1
        return x


myClass = MyNumbers()
myIter = iter(myClass)

print(next(myIter))  # 1
print(next(myIter))  # 2
print(next(myIter))  # 3
print(next(myIter))  # 4
print(next(myIter))  # 5
