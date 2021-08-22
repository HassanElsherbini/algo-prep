package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMaxProductSubArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-3, 0, 1, -2}, 1},
		{[]int{2, 3, -2, -4, 1, 1, -2, -3, -10, 10}, 2400},
	}

	t.Log("Running left and right products solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.MaxProductSubArray(tt.nums))
	}

	t.Log("Running dp solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.MaxProductSubArrayDP(tt.nums))
	}
}
