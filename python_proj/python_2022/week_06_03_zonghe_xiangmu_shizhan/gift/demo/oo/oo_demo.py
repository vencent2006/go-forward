# coding:utf-8

class Parent:
    def myMethod(self):
        print('调用父类方法')


class Child(Parent):
    def myMethod(self):
        print('调用子类方法')


c = Child()
c.myMethod()  # 调用子类方法
super(Child, c).myMethod()  # 调用父类方法
