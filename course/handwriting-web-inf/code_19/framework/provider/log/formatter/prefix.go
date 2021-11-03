/**
 * @Author: vincent
 * @Description:
 * @File:  prefix
 * @Version: 1.0.0
 * @Date: 2021/11/2 18:13
 */

package formatter

import "go-examples/course/handwriting-web-inf/code_19/framework/contract"

func Prefix(level contract.LogLevel) string {
	prefix := ""
	switch level {
	case contract.PanicLevel:
		prefix = "[Panic]"
	case contract.FatalLevel:
		prefix = "[Fatal]"
	case contract.ErrorLevel:
		prefix = "[Error]"
	case contract.WarnLevel:
		prefix = "[Warn]"
	case contract.InfoLevel:
		prefix = "[Info]"
	case contract.DebugLevel:
		prefix = "[Debug]"
	case contract.TraceLevel:
		prefix = "[Trace]"
	}
	return prefix
}
