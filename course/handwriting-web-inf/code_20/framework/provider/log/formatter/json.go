/**
 * @Author: vincent
 * @Description:
 * @File:  json
 * @Version: 1.0.0
 * @Date: 2021/11/2 18:25
 */

package formatter

import (
	"bytes"
	"encoding/json"
	"go-examples/course/handwriting-web-inf/code_20/framework/contract"
	"time"

	"github.com/pkg/errors"
)

func JsonFormatter(level contract.LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	fields["msg"] = msg
	fields["level"] = level
	fields["timestamp"] = t.Format(time.RFC3339)
	c, err := json.Marshal(fields)
	if err != nil {
		// todo: 这里使用了 errors.Wrap
		return bf.Bytes(), errors.Wrap(err, "json format error")
	}

	bf.Write(c)

	return bf.Bytes(), nil
}
