import okx.PublicData as PublicData
import json
from constant import *

publicDataAPI = PublicData.PublicAPI(flag=flag)


def get_instruments():
    """
    Retrieve a list of instruments with open contracts.
    :return:
    """
    result = publicDataAPI.get_instruments(
        instType="SPOT"
    )
    print_pretty_json(result)
    return result


if __name__ == '__main__':
    get_instruments()
