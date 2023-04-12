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
	// 都是99425，就对了
}
