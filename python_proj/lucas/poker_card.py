cardText = [str(number) for number in range(2, 11)] + ['J', 'Q', 'K', 'A']
cardSuit = ['♠', '♥', '♣', '♦']
cardDict = {s + t: cardText.index(t) for t in cardText for s in cardSuit}
cardDict.update({'Joker-': 12, 'Joker+': 13})
print(cardDict)

oneSetCard = [s + t for t in cardText for s in cardSuit] + ['Joker-', 'Joker+']
print(oneSetCard)

import random

for c in range(len(oneSetCard)):
    oneSetCard.append(oneSetCard.pop(random.randint(0, len(oneSetCard) - 1)))

myCard = oneSetCard[int(input('牌已洗好，输入您选第几张(0-53):'))]
robotCard = oneSetCard[random.randint(0, len(oneSetCard) - 1)]
print('您抽中%s,计算机抽中%s' % (myCard, robotCard))
if cardDict[myCard] > cardDict[robotCard]:
    print('恭喜，您赢了')
elif cardDict[myCard] == cardDict[robotCard]:
    print('平局')
else:
    print('遗憾，您输了')
