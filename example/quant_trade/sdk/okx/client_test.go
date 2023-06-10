package okx

import (
	"example/quant_trade/sdk/okx/rest"
	requests "example/quant_trade/sdk/okx/rest/request/market"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_sdk(t *testing.T) {
	r := rest.NewClient(apiKey, secretKey, passphrase, restURL, isSimulated)
	var req requests.GetTicker
	req.InstId = "BTC-USDT"
	ticker, err := r.Market.GetTicker(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("ticker = %+v", ticker)
	for _, ti := range ticker.Tickers {
		t.Logf("%+v", ti)
	}
}

func Test_http(t *testing.T) {
	resp, err := http.Get("https://www.okx.com/api/v5/market/ticker?instId=BTC-USDT")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}
