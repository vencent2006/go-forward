import unittest

import numpy as np


class MyTestCase(unittest.TestCase):
    list1 = [1, 2, 3, 4, 5]
    list2 = [2, 3, 4, 5, 6]

    def test_diff_1(self):
        """使用循环比较"""
        difference = []
        for item in self.list1:
            if item not in self.list2:
                difference.append(item)
        print(difference)

    def test_diff_2(self):
        """使用列表推导式"""
        difference = [item for item in self.list1 if item not in self.list2]
        print(difference)

    def test_diff_3(self):
        """使用集合操作"""
        # self.assertEqual(True, False)  # add assertion here
        set1 = set(self.list1)
        set2 = set(self.list2)
        # list1有，并且，list2没有
        diff = list(set1 - set2)
        print(diff)

    def test_diff_4(self):
        """使用列表函数"""
        difference = list(filter(lambda item: item not in self.list2, self.list1))
        print(difference)

    def test_diff_5(self):
        """使用列表函数"""
        difference = np.setdiff1d()


if __name__ == '__main__':
    unittest.main()
