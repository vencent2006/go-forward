package message

import (
	"bytes"
	"encoding/binary"
)

type Request struct {
	HeadLength uint32
	BodyLength uint32
	RequestID  uint32
	Version    uint8
	Compresser uint8
	Serializer uint8

	ServiceName string
	MethodName  string

	Meta map[string]string

	Data []byte
}

func (r *Request) CalculateHeaderLength() {
	// 不要忘了分隔符
	headerLength := 15 + len(r.ServiceName) + 1 + len(r.MethodName) + 1
	for key, value := range r.Meta {
		headerLength += len(key)
		headerLength++ // key 与 value 之间的分隔符
		headerLength += len(value)
		headerLength++ // key 与 key 之间的分隔符
	}
	r.HeadLength = uint32(headerLength)
}

func (r *Request) CalculateBodyLength() {
	r.BodyLength = uint32(len(r.Data))
}

func EncodeReq(req *Request) []byte {
	bs := make([]byte, req.HeadLength+req.BodyLength)

	// 1. 写入头部长度
	binary.BigEndian.PutUint32(bs[:4], req.HeadLength)
	// 2. 写入body长度
	binary.BigEndian.PutUint32(bs[4:8], req.BodyLength)
	// 3. 写入request id
	binary.BigEndian.PutUint32(bs[8:12], req.RequestID)
	// 4. version
	bs[12] = req.Version    // 1个字节不需要bigEndian了
	bs[13] = req.Compresser // 1个字节不需要bigEndian了
	bs[14] = req.Serializer // 1个字节不需要bigEndian了
	cur := bs[15:]
	copy(cur, req.ServiceName)
	cur = cur[len(req.ServiceName):]
	cur[0] = '\n'
	cur = cur[1:]
	copy(cur, req.MethodName)
	cur = cur[len(req.MethodName):]
	cur[0] = '\n'
	cur = cur[1:]

	// meta
	for key, value := range req.Meta {
		copy(cur, key)
		cur = cur[len(key):]
		cur[0] = '\r' // key和value之间的分隔符
		cur = cur[1:]
		copy(cur, value)
		cur = cur[len(value):]
		cur[0] = '\n' // key和key之间的分隔符
		cur = cur[1:]
	}

	copy(cur, req.Data)
	return bs
}

func DecodeReq(data []byte) *Request {
	req := &Request{}
	// 1. 头部长度
	req.HeadLength = binary.BigEndian.Uint32(data[:4])
	// 2. body长度
	req.BodyLength = binary.BigEndian.Uint32(data[4:8])
	// 3. request id
	req.RequestID = binary.BigEndian.Uint32(data[8:12])
	// 4. version
	req.Version = data[12]
	req.Compresser = data[13]
	req.Serializer = data[14]

	header := data[15:req.HeadLength]
	// 近似于
	// user-service
	// GetById
	index := bytes.IndexByte(header, '\n')
	// 引入分隔符，切分 service name 和 method name
	req.ServiceName = string(header[:index])
	header = header[index+1:] // index就是分隔符本身，所以你要加1
	// 切出来 methodName
	index = bytes.IndexByte(header, '\n')
	req.MethodName = string(header[:index])
	header = header[index+1:] // index就是分隔符本身，所以你要加1

	index = bytes.IndexByte(header, '\n')
	if index != -1 { // 有meta
		meta := make(map[string]string, 4)
		for index != -1 {
			pair := header[:index]
			// \r的位置
			pairIndex := bytes.IndexByte(pair, '\r')
			key := string(pair[:pairIndex])
			value := string(pair[pairIndex+1:])
			meta[key] = value
			header = header[index+1:]
			index = bytes.IndexByte(header, '\n')
		}
		req.Meta = meta
	}
	if req.BodyLength != 0 {
		req.Data = data[req.HeadLength:]
	}
	return req
}
