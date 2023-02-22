from urllib.request import urlopen
import json

myURL = urlopen("https://services.tokenview.io/vipapi/monitor/webhookhistory/eth?page=1&apikey=DHJ6YHJkO8bavEWTNxAJ")
json_str = myURL.read()
json_object = json.loads(json_str)
print(repr(json_object))
json_formatted_str = json.dumps(json_object, indent=2)
print(json_formatted_str)
