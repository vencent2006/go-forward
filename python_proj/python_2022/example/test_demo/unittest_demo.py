import unittest

from doc_test import *


class TestStatisticalFunctions(unittest.TestCase):
    def test_average(self):
        print("enter test_average")
        self.assertEqual(average([20, 30, 70]), 40.0)
        self.assertEqual(round(average([1, 5, 7]), 1), 4.3)
        self.assertRaises(ZeroDivisionError, average, [])
        self.assertRaises(TypeError, average, 20, 30, 70)


if __name__ == '__main__':
    unittest.main()  # Calling from the command line invokes all tests
