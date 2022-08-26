# 1
name = input("What's your name?")
sum_int = 100 + 100
print('hello,%s' % name)
print('sum = %d' % sum_int)

# 2
while True:
    try:
        line = input()
        a = line.split()
        print(int(a[0]) + int(a[1]))
    except:
        break
