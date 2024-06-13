# coding:utf-8
import unittest

from MyNumber import MyNumbers


class DemoTestCase(unittest.TestCase):
    def test_iter(self):
        list = [1, 2, 3, 4]
        it = iter(list)

        while True:
            try:
                print(next(it))
            except StopIteration:
                print('iter finished')
                return


class MyNumberTestCase(unittest.TestCase):
    def test_iter(self):
        myNumber = MyNumbers()
        myiter = iter(myNumber)
        for x in myiter:
            print(x)


if __name__ == '__main__':
    unittest.main()
