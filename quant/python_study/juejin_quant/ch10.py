import matplotlib.pyplot as plt
import numpy as np

# 生成数据
x = np.linspace(0.5, 7.5, 1000)
y = np.sin(x)

# 函数式绘图
# 创建图形并设置大小
plt.figure(figsize=(12, 8))

# sin(x)图
plt.plot(x, y, '--g', lw=2, label='sin(x)')

# 调整坐标轴范围
plt.xlim(0, 10)
plt.ylim(-1.5, 1.5)
# 设置坐标轴标签
plt.xticks(np.arange(0, 10, 2), ['2015-07-02', '2015-08-02', '2015-09-02', '2015-10-02', '2015-11-02'], \
           rotation=45, fontsize=15)
plt.yticks(np.arange(-1, 1.5, 1), [u'最小值', u'零值', u'最大值'], fontsize=15)

# 设置轴标签
plt.xlabel('X axis', fontsize=15)
plt.ylabel('Y axis', fontsize=15)
# 设置网格线
plt.grid(True, ls=':', color='r', alpha=0.5)
# 设置标题
plt.title(u"Functional Programming", fontsize=25)
# 添加图例
plt.legend(loc='upper right', fontsize=15)

# 添加sin(x)的最高点注释
plt.annotate('max sell',
             xy=(np.pi / 2, 1),  # 箭头指向点的坐标
             xytext=(np.pi / 2, 1.3),  # 注释文本左端的坐标
             weight='regular',  # 注释文本的字体粗细风格，bold是粗体，regular是正常粗细
             color='g',  # 注释文本的颜色
             fontsize=15,  # 注释文本的字体大小
             arrowprops={
                 'arrowstyle': '->',  # 箭头类型
                 'connectionstyle': 'arc3',  # 连接类型
                 'color': 'g'  # 箭头颜色
             })

# 添加sin(x)的最低点注释
plt.annotate('min buy',
             xy=(np.pi * 3 / 2, -1),
             xytext=(np.pi * 3 / 2, -1.3),
             weight='regular',
             color='r',
             fontsize=15,
             arrowprops={
                 'arrowstyle': '->',
                 'connectionstyle': 'arc3',
                 'color': 'r'
             })
# 显示图形
# plt.rcParams['font.sans-serif'] = ['Heiti TC'] # 都可以显示中文
plt.rcParams['font.sans-serif'] = ['Arial Unicode MS']
plt.show()
