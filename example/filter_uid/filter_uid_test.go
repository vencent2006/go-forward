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
}
