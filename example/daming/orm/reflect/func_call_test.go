package reflect

import (
	"example/daming/orm/reflect/types"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIterateFunc(t *testing.T) {

	testCases := []struct {
		name    string
		entity  any
		wantErr error
		wantRes map[string]FuncInfo
	}{
		{
			name:    "struct",
			entity:  types.NewUser("Tom", 18),
			wantErr: nil,
			wantRes: map[string]FuncInfo{
				"GetAge": {
					Name: "GetAge",
					// 下标 0 的指向接收器
					InputTypes:  []reflect.Type{reflect.TypeOf(types.User{})},
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
				// 如果接收器是结构体，只能访问接收器是结构体的方法
				//"ChangeName": {
				//	Name:        "ChangeName",
				//	InputTypes:  []reflect.Type{reflect.TypeOf(types.User{}), reflect.TypeOf("")},
				//	OutputTypes: nil,
				//	Result:      nil,
				//},
			},
		},
		{
			name:    "pointer",
			entity:  types.NewUserPtr("Tom", 18),
			wantErr: nil,
			wantRes: map[string]FuncInfo{
				// 如果接收器是结构体指针，接收器是结构体和结构体指针的方法，都能访问
				"GetAge": {
					Name: "GetAge",
					// 下标 0 的指向接收器
					InputTypes:  []reflect.Type{reflect.TypeOf(&types.User{})}, //&types.User{}
					OutputTypes: []reflect.Type{reflect.TypeOf(0)},
					Result:      []any{18},
				},
				"ChangeName": {
					Name:        "ChangeName",
					InputTypes:  []reflect.Type{reflect.TypeOf(&types.User{}), reflect.TypeOf("")}, //&types.User{}
					OutputTypes: []reflect.Type{},
					Result:      []any{},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFunc(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
