# You'll need to install PyJWT via pip 'pip install PyJWT' or your project packages file

import time

import jwt

METABASE_SITE_URL = "http://localhost:3000"
METABASE_SECRET_KEY = "d3fd27a68f9422bd9494ab483d3021e304ee4af8e67b6a0248465605a35858ea"

payload = {
    "resource": {"dashboard": 2},
    "params": {

    },
    "exp": round(time.time()) + (60 * 10)  # 10 minute expiration
}
token = jwt.encode(payload, METABASE_SECRET_KEY, algorithm="HS256")

iframeUrl = METABASE_SITE_URL + "/embed/dashboard/" + token + "#bordered=true&titled=true"

print(iframeUrl)
