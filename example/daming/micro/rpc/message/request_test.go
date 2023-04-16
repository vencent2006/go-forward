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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.req.calculateHeaderLength()
			tc.req.calculateBodyLength()
			data := EncodeReq(tc.req)
			assert.Equal(t, tc.req, DecodeReq(data))
		})

	}
}
