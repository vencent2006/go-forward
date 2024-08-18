import random
import sys

guess = sys.argv[1]
print('你的输入', guess)
b = str(random.randint(1, 3))
print('机器的数', b)
if b == guess:
    print('right')
else:
    print('wrong')
