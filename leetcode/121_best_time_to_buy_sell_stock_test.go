package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuySellStock(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 11, 4, 7}, 10},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.BestTimeToBuySellStock(tt.prices))
	}
}
