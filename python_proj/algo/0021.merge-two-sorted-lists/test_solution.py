import unittest
from solution import *
from typing import List, Optional


def buildListNode(l: List[int]) -> Optional[ListNode]:
    lastNode = None
    for i in range(-1, -len(l) - 1, -1):
        print(l[i])
        node = ListNode(l[i], lastNode)
        lastNode = node

    if l is None:
        return None
    else:
        return lastNode


def listNode2List(node: ListNode) -> list:
    l = []
    while node is not None:
        l.append(node.val)
        node = node.next
    return l


class MyTestCase(unittest.TestCase):
    def test_case_1(self):
        list1 = buildListNode([1, 2, 4])
        list2 = buildListNode([1, 3, 4])
        res = Solution().mergeTwoLists(list1, list2)
        wantRes = [1, 1, 2, 3, 4, 4]
        self.assertEqual(listNode2List(res), wantRes)  # add assertion here

    def test_case_2(self):
        list1 = buildListNode([])
        list2 = buildListNode([])
        res = Solution().mergeTwoLists(list1, list2)
        wantRes = []
        self.assertEqual(listNode2List(res), wantRes)  # add assertion here

    def test_case_3(self):
        list1 = buildListNode([])
        list2 = buildListNode([0])
        res = Solution().mergeTwoLists(list1, list2)
        wantRes = [0]
        self.assertEqual(listNode2List(res), wantRes)  # add assertion here

    def test_buildListNode(self):
        l = [1, 2, 3, 4]
        res = buildListNode(l)
        print(res)

    def test_listNode2List(self):
        node1 = ListNode(4, None)
        node2 = ListNode(3, node1)
        node3 = ListNode(2, node2)
        head = ListNode(1, node3)
        l = listNode2List(head)
        print(l)


if __name__ == '__main__':
    unittest.main()
