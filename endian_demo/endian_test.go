/**
 * @Author: vincent
 * @Description:
 * @File:  endian_test
 * @Version: 1.0.0
 * @Date: 2021/9/7 14:45
 */

package endian_demo

import (
	"encoding/binary"
	"testing"
)

func TestBigAndLittleEndian(t *testing.T) {
	a := uint32(0x01020304)
	arr := make([]byte, 4)

	// big endian
	binary.BigEndian.PutUint32(arr, a)
	t.Log(arr) // [1 2 3 4]

	// little endian
	binary.LittleEndian.PutUint32(arr, a)
	t.Log(arr) // [4 3 2 1]
}
