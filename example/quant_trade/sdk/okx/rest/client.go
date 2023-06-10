package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// ClientRest is the rest api client
type ClientRest struct {
	//Account     *Account
	//SubAccount  *SubAccount
	Trade *Trade
	//Funding     *Funding
	Market *Market
	//PublicData  *PublicData
	//TradeData   *TradeData
	apiKey      string
	secretKey   []byte
	passphrase  string
	isSimulated bool   // 是否是模拟交易
	baseURL     string // 域名地址
	client      *http.Client
}

// NewClient returns a pointer to a fresh ClientRest
func NewClient(apiKey, secretKey, passphrase string, baseURL string, isSimulated bool) *ClientRest {
	c := &ClientRest{
		apiKey:      apiKey,
		secretKey:   []byte(secretKey),
		passphrase:  passphrase,
		baseURL:     baseURL,
		isSimulated: isSimulated,
		client:      http.DefaultClient,
	}
	//c.Account = NewAccount(c)
	//c.SubAccount = NewSubAccount(c)
	//c.Trade = NewTrade(c)
	//c.Funding = NewFunding(c)
	c.Market = NewMarket(c)
	//c.PublicData = NewPublicData(c)
	//c.TradeData = NewTradeData(c)
	return c
}

// Do the http request to the server
func (c *ClientRest) Do(method, path string, private bool, params ...map[string]string) (*http.Response, error) {
	u := fmt.Sprintf("%s%s", c.baseURL, path)
	var (
		r    *http.Request
		err  error
		j    []byte
		body string
	)
	if method == http.MethodGet {
		r, err = http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return nil, err
		}

		if len(params) > 0 {
			q := r.URL.Query()
			for k, v := range params[0] {
				q.Add(k, strings.ReplaceAll(v, "\"", ""))
			}
			r.URL.RawQuery = q.Encode()
			if len(params[0]) > 0 {
				path += "?" + r.URL.RawQuery
			}
		}
	} else {
		j, err = json.Marshal(params[0])
		if err != nil {
			return nil, err
		}
		body = string(j)
		if body == "{}" {
			body = ""
		}
		r, err = http.NewRequest(method, u, bytes.NewBuffer(j))
		if err != nil {
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
	}
	if err != nil {
		return nil, err
	}
	if private {
		timestamp, sign := c.sign(method, path, body)
		r.Header.Add("OK-ACCESS-KEY", c.apiKey)
		r.Header.Add("OK-ACCESS-PASSPHRASE", c.passphrase)
		r.Header.Add("OK-ACCESS-SIGN", sign)
		r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	}
	if c.isSimulated {
		r.Header.Add("x-simulated-trading", "1")
	}
	res, err := c.client.Do(r)
	if err == nil && res.StatusCode != http.StatusOK {
		// status 非200
		return nil, fmt.Errorf("http status(%d) | %s", res.StatusCode, res.Status)
	}
	return res, err
}

// Status
// Get event status of system upgrade
//
// https://www.okex.com/docs-v5/en/#rest-api-status
//func (c *ClientRest) Status(req requests.Status) (responses responses.Status, err error) {
//	p := "/api/v5/system/status"
//	m := okex.S2M(req)
//	res, err := c.Do(http.MethodGet, p, false, m)
//	if err != nil {
//		return
//	}
//	defer res.Body.Close()
//	d := json.NewDecoder(res.Body)
//	err = d.Decode(&responses)
//	return
//}

func (c *ClientRest) sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	t := time.Now().UTC().Format(format)
	ts := fmt.Sprint(t)
	s := ts + method + path + body
	p := []byte(s)
	h := hmac.New(sha256.New, c.secretKey)
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
