/**
 * @Author: vincent
 * @Description:
 * @File:  cache
 * @Version: 1.0.0
 * @Date: 2021/11/21 16:14
 */

package services

import (
	"errors"
	"time"
)

const (
	NoneDuration = time.Duration(-1)
)

var ErrKeyNotFound = errors.New("key not found")
var ErrTypeNotOk = errors.New("val type not ok")
