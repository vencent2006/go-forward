package controller

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreatePostHandler(t *testing.T) {
	//  1. 加载配置
	if err := settings.Init("../config.yaml"); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	//	2. 初始化日子
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	defer zap.L().Sync()
	zap.L().Info("config is ", zap.Any("config", settings.Conf))
	//	3. 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//	4. 初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	// snowflake
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)
	body := `{ 
		"community_id": 1,
		"title": "test",
		"content": "just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// 判断响应的内容是不是按预期返回了需要登录的错误
	// 方法1: 判断响应内容中是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "需要登录")
	// 方法2: 解析json结构
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	t.Logf("json.Unmarshal w.Body: %#v\n", res)
	assert.Equal(t, res.Code, CodeNeedLogin)
}
