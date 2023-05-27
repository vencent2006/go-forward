import okx.Account as Account

from constant import *
from utils import *

accountAPI = Account.AccountAPI(api_key, secret_key, passphrase, False, flag)


def get_all_account_balance():
    """
    https://www.okx.com/docs-v5/zh/#rest-api-account-get-balance
    GET /api/v5/account/balance
    获取全部ccy的余额
    :return:
    """
    result = accountAPI.get_account_balance()
    print_pretty_json(result)
    return result


def get_one_account_balance(ccy):
    """
    https://www.okx.com/docs-v5/zh/#rest-api-account-get-balance
    获取指定ccy的余额
    eg: 获取账户中BTC、ETH两种资产余额 GET /api/v5/account/balance?ccy=BTC,ETH
    :return:
    """
    result = accountAPI.get_account_balance(ccy)
    print_pretty_json(result)
    return result
    """
    {
      "code": "0",
      "data": [
        {
          "adjEq": "",
          "details": [
            {
              "availBal": "8604.73203",
              "availEq": "",
              "cashBal": "8984.73203",
              "ccy": "USDT",
              "crossLiab": "",
              "disEq": "8986.331042094142",
              "eq": "8985.342654402157",
              "eqUsd": "8986.331042094142",
              "fixedBal": "0",
              "frozenBal": "380.6106244021568",
              "interest": "",
              "isoEq": "",
              "isoLiab": "",
              "isoUpl": "",
              "liab": "",
              "maxLoan": "",
              "mgnRatio": "",
              "notionalLever": "",
              "ordFrozen": "380",
              "spotInUseAmt": "",
              "stgyEq": "0.6106244021568",
              "twap": "0",
              "uTime": "1684921492973",
              "upl": "",
              "uplLiab": ""
            }
          ],
          "imr": "",
          "isoEq": "",
          "mgnRatio": "",
          "mmr": "",
          "notionalUsd": "",
          "ordFroz": "",
          "totalEq": "187794.6998519608",
          "uTime": "1684979659450"
        }
      ],
      "msg": ""
    }
    """


if __name__ == '__main__':
    get_all_account_balance()
