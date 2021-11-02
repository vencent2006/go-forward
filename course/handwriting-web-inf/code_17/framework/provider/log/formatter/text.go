/**
 * @Author: vincent
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2021/11/2 18:08
 */

package formatter

import (
	"bytes"
	"fmt"
	"go-examples/course/handwriting-web-inf/code_17/framework/contract"
	"time"
)

// TextFormatter 表示文本格式输出
func TextFormatter(level contract.LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error) {
	// NewBuffer
	bf := bytes.NewBuffer([]byte{})
	separator := "\t"

	// 先输出日志级别
	prefix := Prefix(level)

	bf.WriteString(prefix)
	bf.WriteString(separator)

	// 输出时间
	ts := t.Format(time.RFC3339)
	bf.WriteString(ts)
	bf.WriteString(separator)

	// 输出msg
	bf.WriteString("\"")
	bf.WriteString(msg)
	bf.WriteString("\"")
	bf.WriteString(separator)

	// 输出map
	// todo: fmt.Sprint(fields) 挺有意思
	bf.WriteString(fmt.Sprint(fields))

	return bf.Bytes(), nil
}
