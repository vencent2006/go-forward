import threading

import requests

sum_200 = 0
sum_429 = 0


def make_request():
    global sum_200, sum_429
    url = "http://127.0.0.1:19000/aj/reward/dongtai_get_bangdan_list?vuid=7951028452"
    try:
        response = requests.get(url)
        if response.status_code == 200:
            sum_200 += 1
        elif response.status_code == 429:
            sum_429 += 1
        print(f"Status Code: {response.status_code}, Response: {response.text}")
    except Exception as e:
        print(f"Request failed: {e}")


# 创建n个线程模拟并发
n = 200
threads = []
for i in range(n):
    t = threading.Thread(target=make_request)
    threads.append(t)
    t.start()

# 等待所有线程完成
for t in threads:
    t.join()
print(f"sum_200: {sum_200}, sum_429: {sum_429}")
