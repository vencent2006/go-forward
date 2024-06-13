# coding:utf-8
import unittest


class LambdaTestCase(unittest.TestCase):
    def test_map(self):
        numbers = [x for x in range(1, 6)]
        print(numbers)


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


if __name__ == '__main__':
    unittest.main()
