package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EncodeDecode(t *testing.T) {
	testCases := []struct {
		name string
		req  *Request
	}{
		{
			name: "normal",
			req: &Request{
				RequestID:   123,
				Version:     12,
				Compresser:  13,
				Serializer:  14,
				ServiceName: "user-service",
				MethodName:  "GetById",
				Meta: map[string]string{
					"trace-id": "123456",
					"a/b":      "a",
				},
				Data: []byte("hello, world"),
			},
		},
		{
			name: "no meta with data",
			req: &Request{
				RequestID:   123,
				Version:     12,
				Compresser:  13,
				Serializer:  14,
				ServiceName: "user-service",
				MethodName:  "GetById",
				Data:        []byte("hello, world"),
			},
		},
		{
			name: "no meta no data",
			req: &Request{
				RequestID:   123,
				Version:     12,
				Compresser:  13,
				Serializer:  14,
				ServiceName: "user-service",
				MethodName:  "GetById",
			},
		},
		{
			name: "data with \n",
			req: &Request{
				RequestID:   123,
				Version:     12,
				Compresser:  13,
				Serializer:  14,
				ServiceName: "user-service",
				MethodName:  "GetById",
				Data:        []byte("hello \n world"),
			},
		},
		// 你可以禁止开发者(框架使用者)在 meta 里面使用 \n 和 \r, 所以不会出现这种情况
		//{
		//	name: "meta with \n",
		//	req: &Request{
		//		RequestID:   123,
		//		Version:     12,
		//		Compresser:  13,
		//		Serializer:  14,
		//		ServiceName: "user-\nservice",
		//		MethodName:  "GetById",
		//		Data:        []byte("hello \n world"),
		//	},
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.req.calculateHeaderLength()
			tc.req.calculateBodyLength()
			data := EncodeReq(tc.req)
			decodeReq := DecodeReq(data)
			assert.Equal(t, tc.req, decodeReq)
		})

	}
}
