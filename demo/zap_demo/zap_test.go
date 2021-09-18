/**
 * @Author: vincent
 * @Description:
 * @File:  zap_test
 * @Version: 1.0.0
 * @Date: 2021/9/2 16:10
 */

package zap_demo

import (
	"net/http"
	"testing"

	"go.uber.org/zap"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

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
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		_ = resp.Body.Close()
	}
}

func TestZap(t *testing.T) {
	InitLogger()
	defer logger.Sync()
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}
func InitLoggerSugar() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}
func simpleHttpGetSugar(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		_ = resp.Body.Close()
	}
}

func TestZapSugar(t *testing.T) {
	InitLoggerSugar()
	defer sugarLogger.Sync()
	simpleHttpGetSugar("www.baidu.com")
	simpleHttpGetSugar("http://www.baidu.com")
}
