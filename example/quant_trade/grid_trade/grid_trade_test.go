package grid_trade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalPercentLevels(t *testing.T) {
	gridNum := 1
	levels := CalPercentLevels(gridNum)
	for i, level := range levels {
		t.Logf("%d %.2f", i, level)
	}
}

func Test_CalCalPriceLevels(t *testing.T) {
	highest := float64(30000)
	lowest := float64(25000)
	gridNum := 1
	levels := CalPriceLevels(highest, lowest, gridNum)
	for i, level := range levels {
		t.Logf("%d %.2f", i, level)
	}
}

func Test_GetIndex(t *testing.T) {
	testCases := []struct {
		name        string
		priceLevels []float64
		close       float64
		wantRes     int
	}{
		{
			name:        "大于最大",
			priceLevels: []float64{30000, 29500, 29000, 28500, 28000, 27500, 27000, 26500, 26000, 25500, 25000},
			close:       30001,
			wantRes:     0,
		},
		{
			name:        "小于最小",
			priceLevels: []float64{30000, 29500, 29000, 28500, 28000, 27500, 27000, 26500, 26000, 25500, 25000},
			close:       24999,
			wantRes:     10,
		},
		{
			name:        "大于最大",
			priceLevels: []float64{30000, 29500, 29000, 28500, 28000, 27500, 27000, 26500, 26000, 25500, 25000},
			close:       30001,
			wantRes:     0,
		},
		{
			name:        "中间位置",
			priceLevels: []float64{30000, 29500, 29000, 28500, 28000, 27500, 27000, 26500, 26000, 25500, 25000},
			close:       26666,
			wantRes:     7,
		},
		{
			name:        "没有元素",
			priceLevels: []float64{},
			close:       24999,
			wantRes:     -1,
		},
		{
			name:        "不够一个网格，一个元素(小)",
			priceLevels: []float64{20000},
			close:       24999,
			wantRes:     -1,
		},
		{
			name:        "不够一个网格，比一个元素(大)",
			priceLevels: []float64{30000},
			close:       24999,
			wantRes:     -1,
		},
		{
			name:        "只有一个网格",
			priceLevels: []float64{30000, 20000},
			close:       24999,
			wantRes:     1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := GetIndexByPrice(tc.close, tc.priceLevels)
			target := float64(res) / float64(len(tc.priceLevels)-1)
			t.Logf("GetIndexByPrice(%.1f) gridNum:%d index:%d target:%.1f", tc.close, len(tc.priceLevels)-1, res, target)
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
