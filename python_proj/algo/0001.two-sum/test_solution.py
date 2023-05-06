import unittest
from solution import *


class MyTestCase(unittest.TestCase):

    def test_normal1(self):
        nums = [2, 7, 11, 15]
        target = 9
        want_res = [0, 1]
        res = Solution().twoSum(nums, target)
        self.assertEqual(res, want_res, "test failed ******")

    def test_normal2(self):
        nums = [3, 2, 4]
        target = 6
        want_res = [1, 2]
        res = Solution().twoSum(nums, target)
        self.assertEqual(res, want_res, "test failed ******")

    def test_normal3(self):
        nums = [3, 3]
        target = 6
        want_res = [0, 1]
        res = Solution().twoSum(nums, target)
        self.assertEqual(res, want_res, "test failed ******")


"""
    def test_foreach(self):
        nums_list = [[2, 7, 11, 15], [3, 2, 4], [3, 3]]
        target_list = [9, 6, 6]
        want_list = [[0, 1], [1, 2], [0, 1]]
        i = 0
        while i < len(nums_list):
            nums = nums_list[i]
            target = target_list[i]
            want_res = want_list[i]
            res = Solution().twoSum(nums, target)
            self.assertEqual(res, want_res, "test failed ******")
            i += 1
"""

if __name__ == '__main__':
    unittest.main()
