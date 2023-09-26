package codec

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	//类比接收缓冲区 net.Conn
	bytesBuffer := bytes.NewBuffer([]byte{})
	_ = bytesBuffer

	//发送
	if err := Encode(bytesBuffer, "hello world 0!!!"); err != nil {
		panic(err)
	}
	if err := Encode(bytesBuffer, "hello world 1!!!"); err != nil {
		panic(err)
	}

	//读取
	for {
		if bt, err := Decode(bytesBuffer); err == nil {
			fmt.Println(string(bt))
			continue
		}
		break
	}
}
