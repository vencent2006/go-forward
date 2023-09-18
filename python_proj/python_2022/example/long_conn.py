import time

import requests


def requests_long(uri):
    session = requests.Session()
    for i in range(50):
        response = session.get(uri)
        print(response.status_code)
        time.sleep(300)


if __name__ == '__main__':
    uri = 'http://127.0.0.1:9002'
    requests_long(uri)
