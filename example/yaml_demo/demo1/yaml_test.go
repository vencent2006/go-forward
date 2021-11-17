/**
 * @Author: vincent
 * @Description:
 * @File:  yaml_test
 * @Version: 1.0.0
 * @Date: 2021/11/17 16:25
 */

package demo1

import (
	"testing"

	"github.com/mitchellh/mapstructure"
)

type DBConfig struct {
	// 以下配置关于dsn
	Timeout      string `yaml:"timeout"`      // 写超时时间
	WriteTimeout string `yaml:"writetimeout"` // 写超时时间
	Loc          string `yaml:"loc"`          // 时区
	Port         int    `yaml:"port"`         // 端口
	ReadTimeout  string `yaml:"read_timeout"` // 读超时时间
}

func TestYamlLoad(t *testing.T) {

	// map
	configMap := make(map[string]interface{})
	configMap["timeout"] = "3"
	configMap["writetimeout"] = "1"
	configMap["loc"] = "Loc"
	configMap["port"] = 1234
	configMap["read_timeout"] = "2"

	// structure
	config := &DBConfig{}

	t.Logf("configMap is %+v", configMap)

	// map -> struct
	err := mapstructure.Decode(configMap, config)
	if err != nil {
		t.Error(err)
	}

	t.Logf("config is %+v", config)
}
