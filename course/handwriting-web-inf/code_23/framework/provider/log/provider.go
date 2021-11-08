/**
 * @Author: vincent
 * @Description:
 * @File:  provider
 * @Version: 1.0.0
 * @Date: 2021/11/2 17:49
 */

package log

import (
	"go-examples/course/handwriting-web-inf/code_23/framework"
	"go-examples/course/handwriting-web-inf/code_23/framework/contract"
	"io"
	"strings"

	"go-examples/course/handwriting-web-inf/code_23/framework/provider/log/formatter"
	"go-examples/course/handwriting-web-inf/code_23/framework/provider/log/services"
)

type HadeLogServiceProvider struct {
	// 表示要实现的接口
	framework.ServiceProvider

	// Driver
	Driver string

	Level      contract.LogLevel   // 日志级别
	Formatter  contract.Formatter  // 输出格式
	CtxFielder contract.CtxFielder // context上下文
	Output     io.Writer           // 日志输出信息
}

// Register 注册一个服务实例
func (l *HadeLogServiceProvider) Register(c framework.Container) framework.NewInstance {
	if l.Driver == "" {
		tcs, err := c.Make(contract.ConfigKey)
		if err != nil {
			// 默认使用console
			return services.NewHadeConsoleLog
		}

		cs := tcs.(contract.Config)
		l.Driver = strings.ToLower(cs.GetString("log.Driver"))
	}

	// 根据driver的配置项确定
	switch l.Driver {
	case "single":
		return services.NewHadeSingleLog
	case "rotate":
		return services.NewHadeRotateLog
	case "console":
		return services.NewHadeConsoleLog
	case "custom":
		return services.NewHadeCustomLog
	default:
		return services.NewHadeConsoleLog
	}
}

// Boot 启动的时候注入
func (l *HadeLogServiceProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer 是否延迟加载
func (l *HadeLogServiceProvider) IsDefer() bool {
	return false
}

// Params 定义要传递给实例化方法的参数
func (l *HadeLogServiceProvider) Params(c framework.Container) []interface{} {
	// 获取configService
	configService := c.MustMake(contract.ConfigKey).(contract.Config)

	// 设置参数formatter
	if l.Formatter == nil {
		l.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			v := configService.GetString("log.formatter")
			if v == "json" {
				l.Formatter = formatter.JsonFormatter
			} else if v == "text" {
				l.Formatter = formatter.TextFormatter
			}
		}
	}

	if l.Level == contract.UnknownLevel {
		l.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			l.Level = logLevel(configService.GetString("log.level"))
		}
	}

	// 定义5个参数
	return []interface{}{c, l.Level, l.CtxFielder, l.Formatter, l.Output}
}

// Name 定义对应的服务字符串凭证
func (l *HadeLogServiceProvider) Name() string {
	return contract.LogKey
}

// logLevel get level from string
func logLevel(config string) contract.LogLevel {
	switch strings.ToLower(config) {
	case "panic":
		return contract.PanicLevel
	case "fatal":
		return contract.FatalLevel
	case "error":
		return contract.ErrorLevel
	case "warn":
		return contract.WarnLevel
	case "info":
		return contract.InfoLevel
	case "debug":
		return contract.DebugLevel
	case "trace":
		return contract.TraceLevel
	}
	return contract.UnknownLevel
}
