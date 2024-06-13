# coding:utf-8
import unittest
from functools import reduce


class LambdaTestCase(unittest.TestCase):
    def test_map(self):
        # numbers = [x for x in range(1, 6)]
        numbers = list(range(1, 6))
        squared = list(map(lambda x: x ** 2, numbers))
        print(squared)

    def test_filter(self):
        numbers = list(range(1, 9))
        even_numbers = list(filter(lambda x: x % 2 == 0, numbers))
        print(even_numbers)

    def test_reduce(self):
        numbers = list(range(1, 6))
        product = reduce(lambda x, y: x * y, numbers)
        print(product)
