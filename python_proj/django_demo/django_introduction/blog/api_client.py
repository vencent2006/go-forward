import json

import requests


def call_weibo_api(query, user_id):
    """
    调用微博 AIGC API 的方法。

    :param query: 用户查询的内容
    :param user_id: 用户 ID
    :return: API 响应的 JSON 数据，如果出错则返回 None
    """
    url = 'https://aigc.intra.weibo.com/api/creator/v1/workflows/run'
    headers = {
        'Authorization': 'Bearer app-snAyZaEYJs16BAjWYIZjABNk',
        'Content-Type': 'application/json'
    }
    data = {
        "inputs": {"query": query},
        "response_mode": "blocking",
        "user": user_id
    }

    try:
        response = requests.post(url, headers=headers, data=json.dumps(data))
        response.raise_for_status()  # 检查请求是否成功
        return response.json()  # 返回响应的 JSON 数据
    except requests.RequestException as e:
        print(f"请求出错: {e}")
        return None
