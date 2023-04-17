package message

import (
	"encoding/binary"
)

type Response struct {
	HeadLength uint32
	BodyLength uint32
	RequestID  uint32
	Version    uint8
	Compresser uint8
	Serializer uint8

	Error []byte

	Data []byte
}

func (r *Response) CalculateHeaderLength() {
	r.HeadLength = uint32(15 + len(r.Error))
}

func (r *Response) CalculateBodyLength() {
	r.BodyLength = uint32(len(r.Data))
}

func EncodeResp(resp *Response) []byte {
	bs := make([]byte, resp.HeadLength+resp.BodyLength)

	// 1. 写入头部长度
	binary.BigEndian.PutUint32(bs[:4], resp.HeadLength)
	// 2. 写入body长度
	binary.BigEndian.PutUint32(bs[4:8], resp.BodyLength)
	// 3. 写入request id
	binary.BigEndian.PutUint32(bs[8:12], resp.RequestID)
	// 4. version
	bs[12] = resp.Version    // 1个字节不需要bigEndian了
	bs[13] = resp.Compresser // 1个字节不需要bigEndian了
	bs[14] = resp.Serializer // 1个字节不需要bigEndian了
	cur := bs[15:]

	copy(cur, resp.Error)
	cur = cur[len(resp.Error):]
	copy(cur, resp.Data)
	return bs
}

func DecodeResp(data []byte) *Response {
	resp := &Response{}
	// 1. 头部长度
	resp.HeadLength = binary.BigEndian.Uint32(data[:4])
	// 2. body长度
	resp.BodyLength = binary.BigEndian.Uint32(data[4:8])
	// 3. request id
	resp.RequestID = binary.BigEndian.Uint32(data[8:12])
	// 4. version
	resp.Version = data[12]
	resp.Compresser = data[13]
	resp.Serializer = data[14]
	if resp.HeadLength > 15 {
		resp.Error = data[15:resp.HeadLength]
	}

	if resp.BodyLength != 0 {
		resp.Data = data[resp.HeadLength:]
	}
	return resp
}
