package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RespEncodeDecode(t *testing.T) {
	testCases := []struct {
		name string
		resp *Response
	}{
		{
			name: "normal",
			resp: &Response{
				RequestID:  123,
				Version:    12,
				Compresser: 13,
				Serializer: 14,
				Error:      []byte("this is error"),
				Data:       []byte("hello, world"),
			},
		},
		{
			name: "no error",
			resp: &Response{
				RequestID:  123,
				Version:    12,
				Compresser: 13,
				Serializer: 14,
				Data:       []byte("hello, world"),
			},
		},
		{
			name: "no data",
			resp: &Response{
				RequestID:  123,
				Version:    12,
				Compresser: 13,
				Serializer: 14,
			},
		},
		{
			name: "data with \n",
			resp: &Response{
				RequestID:  123,
				Version:    12,
				Compresser: 13,
				Serializer: 14,
				Data:       []byte("hello \n world"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.resp.calculateHeaderLength()
			tc.resp.calculateBodyLength()
			data := EncodeResp(tc.resp)
			resp := DecodeResp(data)
			assert.Equal(t, tc.resp, resp)
		})

	}
}
