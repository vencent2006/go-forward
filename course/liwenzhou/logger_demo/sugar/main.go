package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer logger.Sync()

	simpleHttpGet("www.sogou.com")
	simpleHttpGet("http://www.sogou.com")
}
