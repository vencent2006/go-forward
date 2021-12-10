package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const fileName = "access_json"

/*
{
    "@timestamp": "2021-12-09T15:19:48+08:00",
    "remote_addr": "10.41.1.46",
    "remote_user": "-",
    "body_bytes_sent": "121",
    "request_time": "0.029",
    "status": "200",
    "host": "verified.biz.weibo.com",
    "request": "GET /bluev/interface/external/matrix/accesswhitelist?source=4037146678&uid=1687465131&domain=1005051687465131 HTTP/1.1",
    "request_method": "GET",
    "uri": "/bluev/verify/bluev/applications/bluev/index.php/interface/external/matrix/accesswhitelist",
    "http_referrer": "https://weibo.com/1687465131/profile?topnav=1&wvr=6",
    "http_x_forwarded_for": "49.7.38.64",
    "http_user_agent": "Weibo.com Swift framework HttpRequest class"
}
*/
type AccessLine struct {
	Timestamp     string `json:"@timestamp"`
	RemoteAddr    string `json:"remote_addr"`
	RemoteUser    string `json:"remote_user"`
	BodyBytesSent string `json:"body_bytes_sent"`
	RequestTime   string `json:"request_time"`
	Status        string `json:"status"`
	Host          string `json:"host"`
	Request       string `json:"request"`
	RequestMethod string `json:"request_method"`
	Uri           string `json:"uri"`
	HttpReferrer  string `json:"http_referrer"`
	HttpXForward  string `json:"http_x_forward"`
	HttpUserAgent string `json:"http_user_agent"`
}

type AccessRequest struct {
	Timestamp     int64
	BodyBytesSent int
	RequestTime   float64
	Status        int
	RawLine       AccessLine
}

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		var req AccessRequest
		err := json.Unmarshal(scanner.Bytes(), &req.RawLine)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%+v\n", req.AccessLine)
		//loc, _ := time.LoadLocation("Asia/Shanghai")
		tTmp, err := time.Parse(time.RFC3339, req.RawLine.Timestamp)
		if err != nil {
			panic(err)
		}
		fmt.Printf("tTmp is %+v\n", tTmp)
		req.Timestamp = tTmp.Unix()
		req.BodyBytesSent, err = strconv.Atoi(req.RawLine.BodyBytesSent)
		req.RequestTime, err = strconv.ParseFloat(req.RawLine.RequestTime, 32)
		req.Status, err = strconv.Atoi(req.RawLine.Status)

		fmt.Printf("%+v\n", req)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
