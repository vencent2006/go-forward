package reflect

import (
	"errors"
	"example/daming/orm/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateFields(t *testing.T) {

	type User struct {
		Name string
		age  int
	}

	testCases := []struct {
		name    string
		entity  any
		wantErr error
		wantRes map[string]any
	}{
		{
			name: "struct",
			entity: User{
				Name: "Tom",
				age:  18,
			},
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name: "pointer",
			entity: &User{
				Name: "Tom",
				age:  18,
			},
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name:    "basic type",
			entity:  18,
			wantErr: errors.New("不支持类型"),
		},
		{
			name: "multiple pointer",
			entity: func() **User {
				res := &User{
					Name: "Tom",
					age:  18,
				}
				return &res
			}(),
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errors.New("不支持 nil"),
		},
		{
			name:    "user nil",
			entity:  (*User)(nil),
			wantErr: errors.New("不支持零值"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIterateFields_FromAnotherPkg(t *testing.T) {

	testCases := []struct {
		name    string
		entity  any
		wantErr error
		wantRes map[string]any
	}{
		{
			name: "struct",
			entity: user.User{
				Name: "Tom",
				//age:  18,
			},
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name: "pointer",
			entity: &user.User{
				Name: "Tom",
				//age:  18,
			},
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name:    "basic type",
			entity:  18,
			wantErr: errors.New("不支持类型"),
		},
		{
			name: "multiple pointer",
			entity: func() **user.User {
				res := &user.User{
					Name: "Tom",
					//age:  18,
				}
				return &res
			}(),
			wantRes: map[string]any{
				"Name": "Tom",
				"age":  0, // age私有的，拿不到，最终我们创建了0值来填充
			},
		},
		{
			name:    "nil",
			entity:  nil,
			wantErr: errors.New("不支持 nil"),
		},
		{
			name:    "user nil",
			entity:  (*user.User)(nil),
			wantErr: errors.New("不支持零值"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFields(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
