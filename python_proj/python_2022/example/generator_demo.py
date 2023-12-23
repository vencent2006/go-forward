def countdown(n):
    while n > 0:
        yield n
        n -= 1


generator = countdown(5)

print(next(generator))  # 5
print(next(generator))  # 4
print(next(generator))  # 3

for value in generator:
    print(value)  # 2, 1
