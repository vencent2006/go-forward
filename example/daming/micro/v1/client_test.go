package v1

import (
	"context"
	"errors"
	"example/daming/micro/v1/message"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_setFuncField(t *testing.T) {

	testCases := []struct {
		name    string
		mock    func(ctrl *gomock.Controller) Proxy
		service Service
		wantErr error
	}{
		{
			name:    "nil",
			service: nil,
			mock: func(ctrl *gomock.Controller) Proxy {
				return NewMockProxy(ctrl)
			},
			wantErr: errors.New("rpc: 不支持nil"),
		},
		{
			name:    "no pointer",
			service: UserService{},
			mock: func(ctrl *gomock.Controller) Proxy {
				return NewMockProxy(ctrl)
			},
			wantErr: errors.New("rpc: 只支持指向结构体的一级指针"),
		},
		{
			name:    "user service",
			service: &UserService{},
			mock: func(ctrl *gomock.Controller) Proxy {
				p := NewMockProxy(ctrl)
				p.EXPECT().Invoke(gomock.Any(), &message.Request{
					ServiceName: "user-service",
					MethodName:  "GetById",
					Data:        []byte(`{"Id":123}`),
				}).Return(&message.Response{}, nil)
				return p
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			err := setFuncField(tc.service, tc.mock(ctrl))
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			resp, err := tc.service.(*UserService).GetById(context.Background(), &GetByIdReq{Id: 123})
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			t.Log(resp)
		})
	}
}
