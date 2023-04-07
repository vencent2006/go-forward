package reflect

import (
	"errors"
	"example/daming/orm/reflect/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateFields(t *testing.T) {

	type User struct {
		Name string
		age  int
	}

	testCases := []struct {
		// test case name
		name string
		// parameters for IterateFields
		entity any
		// want
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
		// test case name
		name string
		// parameters for IterateFields
		entity any
		// want
		wantErr error
		wantRes map[string]any
	}{
		{
			name: "struct",
			entity: types.User{
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
			entity: &types.User{
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
			entity: func() **types.User {
				res := &types.User{
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
			entity:  (*types.User)(nil),
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

func TestSetField(t *testing.T) {
	type User struct {
		Name string
		age  int
	}

	testCases := []struct {
		// test case name
		name string
		// parameters for SetField
		entity any
		field  string
		newVal any
		// want
		wantErr    error
		wantEntity any // 修改后的entity
	}{
		{
			name: "struct",
			entity: User{
				Name: "Tom",
			},
			field:   "Name",
			newVal:  "Jerry",
			wantErr: errors.New("不可修改字段"),
			//wantEntity: User{Name: "Jerry"},
		},
		{
			name: "pointer",
			entity: &User{
				Name: "Tom",
			},
			field:      "Name",
			newVal:     "Jerry",
			wantErr:    nil,
			wantEntity: &User{Name: "Jerry"},
		},
		{
			name: "pointer unexported",
			entity: &User{
				age: 18,
			},
			field:      "age",
			newVal:     19,
			wantErr:    errors.New("不可修改字段"),
			wantEntity: &User{age: 19},
		},
		{
			name: "multiple pointer", // 多级指针的case
			entity: func() **User {
				res := &User{
					Name: "Tom",
				}
				return &res
			}(),
			field:   "Name",
			newVal:  "Jerry",
			wantErr: nil,
			wantEntity: func() **User {
				res := &User{
					Name: "Jerry",
				}
				return &res
			}(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantEntity, tc.entity)
		})
	}

	// unaddressable panic, 这个不通过
	//var i = 0
	//reflect.ValueOf(i).Set(reflect.ValueOf(12))
	//assert.Equal(t, 12, i)

	// 这个可以通过，指针的一般都可以改, 用CanSet可以检验
	//var i = 0
	//ptr := &i
	//reflect.ValueOf(ptr).Elem().Set(reflect.ValueOf(12))
	//assert.Equal(t, 12, i)
}
