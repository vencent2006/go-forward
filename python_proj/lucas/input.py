# a = input()
# print(a)
#
# print(input())
#
# print('Dear ' + input() + ':\n     Happy Birthday!(*^_^*)')

# print(eval('2*3'))
#
# a = 1
# b = 2
# c = 3
# print(a, b, c)
# a, b, c = c, a, a
# print(a, b, c)

# height = input("输入您身高:")
# print(float(height) > 1.4)

# if float(input('输入一个负数:')) < 0:
#     print('正确,门开了!')

import random

print(random.randint(1, 10))

# key = str(random.randint(1, 3))
# key = random.randint(1, 3)
# print(key)
# if input('猜一猜1~3:') == key:
#     print("you win!")
# else:
#     # print("you lose, i think it's " + key)
#     print("you lose, i think it's ")
print([x for x in range(10)])
team = ['张飞', '赵云', '张辽', '马谡', '韩信', '吕布', '关羽', '马超', '马忠', '黄忠']
print({name[0] for name in team})
