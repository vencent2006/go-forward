/**
 * @Author: vincent
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2021/11/1 17:33
 */

package contract

import "time"

const (
	// ConfigKey is config key in container
	ConfigKey = "hade:config"
)

// Config 定义了配置文件服务，读取配置文件，支持点分割的路径读取
// 例如: .Get("app.name") 表示从 app 文件中读取 name 属性
// 建议使用 yaml 属性, https://yaml.org/spec/1.2/spec.html
type Config interface {
	// IsExist 检查一个属性是否存在
	IsExist(key string) bool
	// Get 获取一个属性值
	Get(key string) interface{}
	// GetBool 获取一个 bool 属性
	GetBool(key string) bool
	// GetInt 获取一个 int 属性
	GetInt(key string) int
	// GetFloat64 获取一个 float64 属性
	GetFloat64(key string) float64
	// GetTime 获取一个 time 属性
	GetTime(key string) time.Time
	// GetString 获取一个 string 属性
	GetString(key string) string
	// GetIntSlice 获取一个 int 数组属性
	GetIntSlice(key string) []int
	// GetStringSlice 获取一个 string 数组
	GetStringSlice(key string) []string
	// GetStringMap 获取一个 string 为 key，interface 为 val 的 map
	GetStringMap(key string) map[string]interface{}
	// GetStringMapString 获取一个 string 为 key，string 为 val 的 map
	GetStringMapString(key string) map[string]string
	// GetStringMapStringSlice 获取一个 string 为 key，数组 string 为 val 的 map
	GetStringMapStringSlice(key string) map[string][]string
	// Load 加载配置到某个对象
	Load(key string, val interface{}) error
}
