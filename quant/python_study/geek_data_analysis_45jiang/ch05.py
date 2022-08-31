import pandas as pd
from pandasql import sqldf, load_meat, load_births

STAR_NUM = 50


def print_star():
    print('*' * STAR_NUM)


print_star()
x1 = pd.Series([1, 2, 3, 4])
x2 = pd.Series(data=[1, 2, 3, 4], index=['a', 'b', 'c', 'd'])
print(x1)
print(x2)

print_star()
d = {'a': 'b', 'c': 'd', 'e': 'f'}
print(pd.Series(d))

print_star()
data = {'Chinese': [66, 95, 93, 90, 80], 'English': [65, 85, 92, 88, 90], 'Math': [30, 98, 96, 77, 90]}
df1 = pd.DataFrame(data)
df2 = pd.DataFrame(data,
                   index=['ZhangFei', 'GuanYu', 'ZhaoYun', 'HuangZhong', 'DianWei'],
                   columns=['Math', 'Chinese', 'English'])
print(df1)
print(df2)


def plus(df, n, m):
    df['new1'] = (df[u'Chinese'] + df[u'English']) * m
    df['new2'] = (df[u'Chinese'] + df[u'English']) * n
    return df


print(df1)
df1 = df1.apply(plus, axis=1, args=(2, 3,))
print(df1)

print_star()
df1 = pd.DataFrame({'name': ['ZhangFei', 'GuanYu', 'a', 'b', 'c'], 'data1': range(5)})
pysqldf = lambda sql: sqldf(sql, globals())
sql = "select * from df1 where name ='ZhangFei'"
print(pysqldf(sql))
sql = "select * from df1 where name ='GuanYu'"
print(pysqldf(sql))

# 练习题
# 对于下表的数据，请使用 Pandas 中的 DataFrame 进行创建，并对数据进行清洗。同时新增一列“总和”计算每个人的三科成绩之和。
# ['张飞', '关羽', '赵云', '黄忠', '典韦', '典韦']
data = {'语文': [66, 95, 95, 90, 80, 80], '英语': [65, 85, 92, 88, 90, 90], '数学': [None, 98, 96, 77, 90, 90]}
df = pd.DataFrame(data, index=['张飞', '关羽', '赵云', '黄忠', '典韦', '典韦'], columns=['语文', '英语', '数学'])
print(df.dtypes)
# df['数学'] = df['数学'].astype(int)
df['数学'].fillna(df['数学'].mean(), inplace=True)
print(df)
df = df.drop_duplicates()
print(df)
# df['总和'] = df['语文']+df['英语']+df['数学']
df['总和'] = df.sum(axis=1)  # 这个更加简便
print(df)
