tinydict = {'name': 'runoob', 'likes': 123, 'url': 'www.runoob.com'}
print(type(tinydict))
print(tinydict)

# python 的 dict 不是 map
# map是个 内置函数 而已

numbers = [x + 1 for x in range(5)]
print(numbers)
res = map(lambda x: x ** 2, numbers)
print(res)
for key in res:
    print(key)
squared = list(res)
print(squared)

from functools import reduce

numbers = [1, 2, 3, 4, 5]

# 使用 reduce() 和 lambda 函数计算乘积
product = reduce(lambda x, y: x * y, numbers)

print(product)  # 输出：120
