import numpy as np
import pandas as pd

data = {"grammar": ["Python", "C", "Java", "GO", np.nan, "SQL", "PHP", "Python"],
        "score": [1, 2, np.nan, 4, 5, 6, 7, 10]}

df = pd.DataFrame(data)
print(df)

results = df["grammar"].str.contains("Python")
results.fillna(value=False, inplace=True)
print(df[results])
