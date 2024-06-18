# coding:utf-8
import calendar
import cmath
import random
import unittest
from functools import reduce


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

    def test_sushu(self):
        """
        prime number，素数
        :return:
        """
        lower = int(input('input lower:'))
        upper = int(input('input upper:'))
        print('\nlower:{}, upper:{}'.format(lower, upper))
        for n in range(lower, upper + 1):
            if n > 1:
                for i in range(2, n):
                    if n % i == 0:
                        break
                else:
                    print(n)

    def test_jiecheng(self):
        n = int(input('input a integer: '))
        if n < 0:
            raise ValueError('minus not has factorial')
        elif n == 0:
            res = 1
        else:
            numbers = list(range(1, n + 1))
            res = reduce(lambda x, y: x * y, numbers)

        print(f'\n{n} factorial = {res}\n')

    def test_9_multiple_9(self):
        for i in range(1, 10):
            for j in range(1, i + 1):
                print(f'{j}x{i}={i * j}', end=' ')
            print()

    def test_char_transform(self):
        """Python ASCII码与字符相互转换"""
        c = input('input a char: ')
        print()
        a = int(input('input a ASCII: '))
        print()
        print(c + ' ASCII = ', ord(c))
        print(a, ' mapping char = ', chr(a))

    def test_calendar(self):
        """Python 生成日历"""
        yy = int(input('input year: '))
        mm = int(input('input month: '))
        print(calendar.month(yy, mm))

    def test_recur_fibo(self):
        def recur_fibo(n):
            if n <= 1:
                return n
            else:
                return recur_fibo(n - 1) + recur_fibo(n - 2)

        nterms = int(input('input terms count: '))
        if nterms <= 0:
            return 'input positive integer'
        else:
            print('fibonacci sequence')
            for i in range(nterms):
                print(recur_fibo(i), end=' ')
