package filter_uid

import (
	"runtime"
	"testing"
)

// awk -F, '{print $1}' order.csv | sort -n | uniq | wc -l
// wc -l huomian_uid.txt
func Test_doFilter(t *testing.T) {
	runtime.GOMAXPROCS(2)
	doFilter()
	//wc -l huomian_uid.txt
	// 99425
	// awk -F, '{print $1}' order.csv | sort -n | uniq | wc -l
	// 99425
	// 二者数据一样，就是对了，当前的执行就是把所有的uid去重，放入到huomian_uid.txt中，所以是一样的
}
