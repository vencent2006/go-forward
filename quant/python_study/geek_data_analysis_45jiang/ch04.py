import numpy as np

"""
创建数组
"""
a = np.array([1, 2, 3])
b = np.array([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
b[1, 1] = 10
print(a.shape)
print(b.shape)
print(a.dtype)
print(b)

"""
结构数组
"""
person_type = np.dtype({
    'names': ['name', 'age', 'chinese', 'math', 'english'],
    'formats': ['S32', 'i', 'i', 'i', 'f']
})

people = np.array([("ZhangFei", 32, 75, 100, 90),
                   ("GuanYu", 24, 85, 96, 88.5),
                   ("ZhaoYun", 28, 85, 92, 96.5),
                   ("HuangZhong", 29, 65, 85, 100)], dtype=person_type)
ages = people[:]['age']
chineses = people[:]['chinese']
maths = people[:]['math']
englishs = people[:]['english']
print(np.mean(ages))
print(np.mean(chineses))
print(np.mean(maths))
print(np.mean(englishs))

# 连续数组的创建，算数运算
x1 = np.arange(1, 11, 2)
x2 = np.linspace(1, 9, 5)
print('x1 =', x1)
print('x2 =', x2)
print(np.add(x1, x2))
print(np.subtract(x1, x2))
print(np.multiply(x1, x2))
print(np.divide(x1, x2))
print(np.power(x1, x2))
print(np.remainder(x1, x2))

# 统计函数
a = np.array([[1, 2, 3], [4, 5, 6], [7, 8, 9]])
print(np.amin(a))
print(np.amin(a, 0))
print(np.amin(a, 1))
print(np.amax(a))
print(np.amax(a, 0))
print(np.amax(a, 1))
print(np.ptp(a))
print(np.ptp(a, 0))
print(np.ptp(a, 1))

# 排序
print('******* 排序 *******')
a = np.array([[4, 3, 2], [2, 4, 1]])
print(np.sort(a))
print(np.sort(a, axis=None))
print(np.sort(a, axis=0))
print(np.sort(a, axis=1))

# https://blog.csdn.net/studyvcmfc/article/details/107734576
a = np.random.randint(2, 40, size=(2, 3, 4))
print("=" * 40, 'a')
print(a)
print("=" * 40, 'np.amin(a, 0)')
print(np.amin(a, 0))
print("=" * 40, 'np.amin(a, 1)')
print(np.amin(a, 1))
print("=" * 40, 'np.amin(a, 2)')
print(np.amin(a, 2))
print("=" * 40, 'np.amin(a, (0, 2))')
print(np.amin(a, (0, 2)))

# 标准差和方差
a = np.array([1, 2, 3, 4])
print(np.std(a))  # 标准差
print(np.var(a))  # 方差

