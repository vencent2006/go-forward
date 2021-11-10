/**
 * @Author: vincent
 * @Description:
 * @File:  util
 * @Version: 1.0.0
 * @Date: 2021/11/9 16:00
 */

package blog_demo

import (
	"math/rand"
	"time"
)

// 返回截止到当前时间的随机日期
func randDateUpToNow() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	// max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
