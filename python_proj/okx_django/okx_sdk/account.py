import okx.Account as Account
import json
from constant import *

accountAPI = Account.AccountAPI(api_key, secret_key, passphrase, False, flag)


def get_all_account_balance():
    """
    获取全部ccy的余额
    :return:
    """
    result = accountAPI.get_account_balance()
    print_pretty_json(result)
    return result


def get_one_account_balance(ccy):
    """
    获取指定ccy的余额
    :return:
    """
    result = accountAPI.get_account_balance(ccy)
    print_pretty_json(result)
    return result


if __name__ == '__main__':
    get_all_account_balance()
