/**
 * @Author: vincent
 * @Description:
 * @File:  demo_test
 * @Version: 1.0.0
 * @Date: 2021/8/3 18:00
 */

package serialization

import (
	"encoding/binary"
	"encoding/json"
	"go-examples/demo/serialization/pkt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// 测试大端和小端的数据呈现
func TestEndian(t *testing.T) {
	// golang

	//小端序：低位字节在前，高位字节在后，计算机方便处理。
	//大端序：高位字节在前，低位字节在后，与人的使用习惯相同。
	a := uint32(0x01020304)
	arr := make([]byte, 4)
	binary.BigEndian.PutUint32(arr, a)
	t.Log(arr) //[1 2 3 4]

	binary.LittleEndian.PutUint32(arr, a)
	t.Log(arr) //[4 3 2 1]
}

// 测试序列化
func TestSerialization(t *testing.T) {
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
	i += 4
	endian.PutUint64(buf[i:i+8], pkt.Sequence)
	i += 8
	// 由于data长度不确定，必须先把长度写入buf, 这样在反序列化时就可以正确的解析出data
	dataLen := len(pkt.Data)
	endian.PutUint32(buf[i:i+4], uint32(dataLen))
	i += 4
	// 写入数据data
	copy(buf[i:i+dataLen], pkt.Data)
	i += dataLen
	t.Log(buf[0:i])
	t.Log("length", i)
}

// 测试反序列化
func TestDeserialization(t *testing.T) {
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

// 测试protobuf在小数据时的表现
func TestProtobufSmall(t *testing.T) {
	p := pkt.Pkt{
		Source:   257,
		Sequence: 5,
		Data:     []byte("hello world"),
	}

	bts, err := proto.Marshal(&p)
	assert.Nil(t, err)
	t.Log(bts)
	t.Log("length ", len(bts))
}

//测试protobuf在大数据时的表现
func TestProtobufBig(t *testing.T) {
	p := pkt.Pkt{
		Source:   10000000,
		Sequence: 2<<60 + 3, // 使最低位有值
		Data:     []byte("hello world"),
	}

	bts, err := proto.Marshal(&p)
	assert.Nil(t, err)
	t.Log(bts)
	t.Log("length ", len(bts))
}

func TestJson(t *testing.T) {
	p := pkt.Pkt{
		Source:   10000000,
		Sequence: 2<<60 + 3, // 使最低位有值
		Data:     []byte("hello world"),
	}

	bts, err := json.Marshal(&p)
	assert.Nil(t, err)
	t.Log(bts)
	t.Log("length ", len(bts))
}

// ********* benchmark: go test -bench=. -benchmem
func Benchmark_Protobuf(t *testing.B) {
	p := pkt.Pkt{
		Source:   10000000,
		Sequence: 2<<60 + 3,
		Data:     []byte("hello world"),
	}
	for i := 0; i < t.N; i++ {
		bts, err := proto.Marshal(&p)
		assert.Nil(t, err)
		assert.NotEmpty(t, bts)
	}
}

func Benchmark_Json(t *testing.B) {
	p := pkt.Pkt{
		Source:   10000000,
		Sequence: 2<<60 + 3,
		Data:     []byte("hello world"),
	}
	for i := 0; i < t.N; i++ {
		bts, err := json.Marshal(&p)
		assert.Nil(t, err)
		assert.NotEmpty(t, bts)
	}
}
