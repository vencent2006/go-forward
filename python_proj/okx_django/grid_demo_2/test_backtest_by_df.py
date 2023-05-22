import os.path
import unittest

import pandas as pd

import constant
from backtest_by_df import backtest_grid_trading


class Test(unittest.TestCase):
    def test_backtest_grid_trading(self):
        file_name = constant.test_data_csv_file
        df = pd.read_csv(file_name, index_col=0)
        print(df)

        # res
        dic = backtest_grid_trading(df, os.path.basename(file_name))
        df_b = dic['df_b']
        df_s = dic['df_s']
        profit = dic['profit']
        max_consume_money = dic['max_consume_money']

        # print
        print("***** df_b ******")
        print(df_b)
        print("***** df_s ******")
        print(df_s)
        print("***** profit ******")
        print(profit)
        print("***** max_consume_money ******")
        print(max_consume_money)

        self.assertEqual(profit, 1647.5)
        self.assertEqual(max_consume_money, 15882)


if __name__ == '__main__':
    unittest.main()
