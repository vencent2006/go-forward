/**
 * @Author: vincent
 * @Description:
 * @File:  4_2_test
 * @Version: 1.0.0
 * @Date: 2022/9/21 09:44
 */

package __2

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

/*
// Variant with *byte pointer type for DWARF debugging.
// runtime/string.go
type stringStructDWARF struct {
	str *byte
	len int
}

// reflect/value.go
type StringHeader struct {
	Data uintptr
	Len  int
}
*/

func TestString(t *testing.T) {
	fmt.Println(unsafe.Sizeof("慕课网"))
	fmt.Println(unsafe.Sizeof("慕课网Moody老师"))

	getLen("慕课网")        // 9 (慕课网的unicode的utf-8）
	getLen("慕课网Moody老师") // 20: 英文5个字节（Moody），每个汉字是3个字节

	printString("慕课网")

	printString(sliceString("慕课网Moody老师", 3))

}

func getLen(s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sh.Len)
}

func printString(s string) {
	// 按照rune来遍历
	for _, char := range s {
		fmt.Println(char)
		//println函数可以打印字符，需要注意的是println()将字符转换成int32格式后输出并换行。
		fmt.Printf("%c", char)
		fmt.Println("\n******")
	}
}

// 取前i个rune后的字符串
func sliceString(s string, i int) string {
	// string->rune的切片->string
	return string([]rune(s)[:i])
}
