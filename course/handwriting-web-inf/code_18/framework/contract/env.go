/**
 * @Author: vincent
 * @Description:
 * @File:  env
 * @Version: 1.0.0
 * @Date: 2021/11/1 09:58
 */

package contract

const (
	EnvProduction  = "production"  // 生产环境
	EnvTesting     = "testing"     // 测试环境
	EnvDevelopment = "development" // 开发环境
	EnvKey         = "hade:env"    // 环境变量服务字符串凭证
)

// Env 定义环境变量服务
type Env interface {
	// AppEnv 获取当前的环境，建议分为 development/testing/production
	AppEnv() string
	// IsExist 判断一个环境变量是否有被设置
	IsExist(string) bool
	// Get 获取某个环境变量, 如果没有设置，返会""
	Get(string) string
	// All 获取所有的环境变量， .env 和运行环境变量融合后的结果
	All() map[string]string
}
