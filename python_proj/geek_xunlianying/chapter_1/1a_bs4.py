import requests
from bs4 import BeautifulSoup as bs

user_agent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Mobile Safari/537.36"

header = {'user-agent':user_agent}

myurl = 'https://movie.douban.com/top250'

response = requests.get(myurl, headers=header)

bs_info = bs(response.text, 'html.parser')

for tags in bs_info.find_all('div', attrs={'class':'hd'}):
    for atag in tags.find_all('a', ):
        print(atag.get('href')) # 获取链接
        print(atag.find('span', ).text) # 获取名字
