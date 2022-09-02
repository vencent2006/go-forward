import urllib3
import urllib
from urllib import request

# print(dir(urllib))
# print('#'*20)
# print(dir(urllib3))

# urllib request
resp = request.urlopen("http://image.baidu.com/")
print('#' * 20, 'type(urllib request)')
print(type(resp))
# print(resp.read().decode())

# urllib3: pool manager
http = urllib3.PoolManager()
resp_dat = http.request('GET', "http://image.baidu.com/")
print('#' * 20, 'type(pool manager)')
print(type(resp_dat))
print(dir(resp_dat))
