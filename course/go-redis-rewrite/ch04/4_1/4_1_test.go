/**
 * @Author: vincent
 * @Description:
 * @File:  4_1
 * @Version: 1.0.0
 * @Date: 2022/9/21 08:01
 */

package __1

import (
	"fmt"
	"testing"
	"unsafe"
)

type K struct {
}

type F struct {
	num1 K
	num2 int32
}

func TestEmptyStructure(t *testing.T) {
	a := F{}
	b := K{}
	c := K{}
	fmt.Printf("%p\n", &a.num1)   //	0xc0000a61f8
	fmt.Printf("%p\n", &b)        //0x1255790
	fmt.Printf("%p\n", &c)        //0x1255790 b和c都是zero_base
	fmt.Println(unsafe.Sizeof(a)) //4
	fmt.Println(unsafe.Sizeof(b)) //0
	fmt.Println(unsafe.Sizeof(c)) //0

}
