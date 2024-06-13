# coding:utf-8
import cmath
import random
import unittest


class ExerciseTestCase(unittest.TestCase):
    def test_add(self):
        a = input('please input a:')
        b = input('please input b:')
        print(type(a))
        print(type(a))
        c = float(a) + float(b)
        print('数字{0}和{1} 相加的结果是{2}'.format(a, b, c))

    def test_sqart(self):
        num = float(input('请输入一个数字： '))
        num_sqrt = num ** 0.5
        print(' %0.3f 的平方根为 %0.3f' % (num, num_sqrt))

        num = int(input("请输入一个数字: "))
        num_sqrt = cmath.sqrt(num)
        print('{0} 的平方根为 {1:0.3f}+{2:0.3f}j'.format(num, num_sqrt.real, num_sqrt.imag))

    def test_random(self):
        # random.random() 返回一个介于 0.0 和 1.0 之间的随机小数
        print(random.random())
        # random.randint(a, b) 用于返回一个介于 a 和 b 之间的整数（包括 a 和 b）
        for i in range(10):
            print(random.randint(1, 4), end=' ')
        print()
        # random.choice(sequence) 用于从序列中随机选择一个元素
        list1 = list(range(1, 6))
        for i in range(10):
            print(random.choice(list1), end=' ')
        print()
        # random.shuffle(sequence) 用于将序列中的元素进行随机排序
        print(list1)
        random.shuffle(list1)
        print(list1)
        random.shuffle(list1)
        print(list1)