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

func TestSerialize(t *testing.T) {
	var pkt = struct {
		Source   uint32
		Sequence uint64
		Data     []byte
	}{
		Source:   257,
		Sequence: 5,
		Data:     []byte("hello world"),
	}

	// 为了方便观看，使用大端序
	endian := binary.BigEndian

	buf := make([]byte, 1024) // buffer
	i := 0
	endian.PutUint32(buf[i:i+4], pkt.Source)
	t.Logf("buf[%d:%d]=%d", i, i+4, buf[i:i+4])
	i += 4
	endian.PutUint64(buf[i:i+8], pkt.Sequence)
	t.Logf("buf[%d:%d]=%d", i, i+8, buf[i:i+8])
	i += 4
	i += 8
	// 由于data长度不确定，必须先把长度写入buf, 这样在反序列化时就可以正确的解析出data
	dataLen := len(pkt.Data)
	endian.PutUint32(buf[i:i+4], uint32(dataLen))
	t.Logf("buf[%d:%d]=%d", i, i+4, buf[i:i+4])
	i += 4
	// 写入数据data
	copy(buf[i:i+dataLen], pkt.Data)
	t.Logf("buf[%d:%d]=%s", i, i+dataLen, buf[i:i+dataLen])
	i += dataLen
	t.Log(buf[0:i])
	t.Log("length", i)
}

func TestUnserialize(t *testing.T) {
	var pkt struct {
		Source   uint32
		Sequence uint64
		Data     []byte
	}

	recv := []byte{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 11, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	endian := binary.BigEndian
	i := 0
	pkt.Source = endian.Uint32(recv[i : i+4])
	i += 4
	pkt.Sequence = endian.Uint64(recv[i : i+8])
	i += 8
	dataLen := endian.Uint32(recv[i : i+4])
	i += 4
	pkt.Data = make([]byte, dataLen)
	copy(pkt.Data, recv[i:i+int(dataLen)])
	t.Logf("Src:%d Seq:%d Data:%s", pkt.Source, pkt.Sequence, pkt.Data)
}
