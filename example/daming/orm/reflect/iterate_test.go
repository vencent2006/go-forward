package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateArrayOrSlice(t *testing.T) {

	testCases := []struct {
		name     string
		entity   any
		wantVals []any
		wantErr  error
	}{
		{
			name:     "array",
			entity:   [3]int{1, 2, 3},
			wantVals: []any{1, 2, 3},
			wantErr:  nil,
		},
		{
			name:     "slice",
			entity:   []int{1, 2, 3},
			wantVals: []any{1, 2, 3},
			wantErr:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateArrayOrSlice(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVals, res)

		})
	}
}

func TestIterateMap(t *testing.T) {

	testCases := []struct {
		name     string
		entity   any
		wantKeys []any
		wantVals []any
		wantErr  error
	}{
		{
			name:     "map",
			entity:   map[string]int{"a": 1, "b": 2, "c": 3},
			wantKeys: []any{"a", "b", "c"},
			wantVals: []any{1, 2, 3},
			wantErr:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, vals, err := IterateMap(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantVals, vals)

		})
	}
}

func TestIterateMap2(t *testing.T) {

	testCases := []struct {
		name     string
		entity   any
		wantKeys []any
		wantVals []any
		wantErr  error
	}{
		{
			name:     "map",
			entity:   map[string]int{"a": 1, "b": 2, "c": 3},
			wantKeys: []any{"a", "b", "c"},
			wantVals: []any{1, 2, 3},
			wantErr:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, vals, err := IterateMap2(tc.entity)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantVals, vals)

		})
	}
}
