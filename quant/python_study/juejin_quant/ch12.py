import numpy as np
import matplotlib.pyplot as plt


def random_walk(nsteps=1000):
    draws = np.random.randint(0, 2, size=nsteps)
    print(f'random walk direction is {draws}')  # random walk direction is [1 0 1 ... 0 1 0]
    steps = np.where(draws > 0, 1, -1)  # 将0转换为-1
    walk = steps.cumsum()  # 累加方式记录轨迹
    return walk


# 多样本随机漫步轨迹-plot
def simplot_random_walk():
    _ = [plt.plot(np.arange(2000), random_walk(nsteps=2000), c='b', alpha=0.05) for _ in np.arange(0, 1000)]
    plt.xlabel('游走步数')
    plt.ylabel('分布轨迹')
    plt.title(u"模拟随机漫步")
    plt.rcParams['font.sans-serif'] = ['Heiti TC']  # 显示中文
    plt.show()


simplot_random_walk()
