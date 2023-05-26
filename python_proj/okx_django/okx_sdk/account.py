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


if __name__ == '__main__':
    get_all_account_balance()
