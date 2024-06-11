# coding:utf-8

class MyClass:
    @staticmethod
    def smeth():
        print('this is a static method')

    @classmethod
    def cmeth(cls):
        print('this is a class method of', cls)


if __name__ == '__main__':
    MyClass.smeth()
    MyClass.cmeth()
