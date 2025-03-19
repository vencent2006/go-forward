import numpy as np
import pandas as pd
import streamlit as st

# 创建示例数据

df = pd.DataFrame(np.random.randn(10, 2), columns=['A', 'B'])

# 添加滑动条

slider_val = st.slider('选择一个阈值', -2.0, 2.0, 0.0)

# 根据滑动条值过滤数据

filtered_df = df[df['A'] > slider_val]

# 显示数据表格

st.write('筛选后的数据：')

st.write(filtered_df)
