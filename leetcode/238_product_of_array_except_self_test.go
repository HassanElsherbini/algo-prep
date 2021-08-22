package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 2, 3, 4}, []int{24, 24, 16, 12}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{2, 4}, []int{4, 2}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.ProductOfArrayExceptSelf(tt.nums))
	}
}
