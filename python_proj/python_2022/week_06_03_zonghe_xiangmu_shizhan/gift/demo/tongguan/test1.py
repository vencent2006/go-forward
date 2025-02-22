import pandas as pd

# 账本 1 的数据
data1 = {
    '日期': ['2025-02-18', '2025-02-19', '2025-02-20'],
    '商品': ['苹果', '香蕉', '橙子'],
    '数量（斤）': [100, 50, 80],
    '类型': ['进货', '进货', '进货'],
    '单价（元/斤）': [2.5, 3.0, 2.0]
}

# 账本 2 的数据
data2 = {
    '日期': ['2025-02-18', '2025-02-19', '2025-02-20'],
    '商品': ['苹果', '香蕉', '橙子'],
    '数量（斤）': [100, 50, 80],
    '类型': ['销售', '销售', '销售'],
    '单价（元/斤）': [3.5, 3.2, 2.7]
}

# 将数据转换为 DataFrame
df1 = pd.DataFrame(data1)
df2 = pd.DataFrame(data2)
print(df1)
print(df2)
# 合并两个账本的数据
df_combined = pd.merge(df1, df2, on=['日期', '商品'], suffixes=('_账本1', '_账本2'))

# 计算销售总收入
df_combined['销售总收入（元）'] = df_combined['数量（斤）_账本2'] * df_combined['单价（元/斤）_账本2']

# 计算销售比例
df_combined['销售比例'] = df_combined['销售总收入（元）'] / df_combined['数量（斤）_账本2']

# 输出包含销售比例的 DataFrame
result_df = df_combined[['商品', '销售比例']]
print(result_df)
