import random

volunteer = 500
print('volunteer = ', volunteer)
while True:
    if volunteer <= 500 - 10:
        print('您已经选够10个人了')
        exit()
    test = random.randint(1, volunteer)
    if input('还剩下' + str(volunteer) + '人, 这是随机挑选的第' + str(test) + '号, 您满意吗(y/n)?') == 'y':
        volunteer = volunteer - 1
