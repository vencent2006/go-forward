import okx_sdk
from django.http import JsonResponse


def JsonResponseOk(result):
    return JsonResponse(result, status=200)


# market
def get_ticker(request):
    instId = request.GET.get("instId")
    result = okx_sdk.get_ticker(instId)
    return JsonResponseOk(result)


def get_tickers(request):
    result = okx_sdk.get_tickers()
    return JsonResponseOk(result)


# account balance
def get_all_account_balance(request):
    result = okx_sdk.get_all_account_balance()
    return JsonResponseOk(result)


def get_one_account_balance(request):
    ccy = request.GET.get("ccy")
    result = okx_sdk.get_one_account_balance(ccy)
    return JsonResponseOk(result)


# public data
def get_instruments(request):
    result = okx_sdk.get_instruments()
    return JsonResponseOk(result)
