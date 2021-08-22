package dp_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/dp"
	"github.com/stretchr/testify/assert"
)

func TestKnapSack(t *testing.T) {
	tests := []struct {
		profits  []int
		weights  []int
		capacity int
		expected int
	}{
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 7, 22},
		{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6, 17},
		{[]int{4, 5, 3, 7}, []int{2, 3, 1, 4}, 5, 10},
		{[]int{4, 5, 3, 7}, []int{6, 7, 21, 9}, 5, 0},
	}

	t.Log("Running top-down solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, dp.KnapSackTopDown(tt.profits, tt.weights, tt.capacity))
	}

	t.Log("Running bottom-up solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, dp.KnapSackBottomUp(tt.profits, tt.weights, tt.capacity))
	}
}

func TestEqualSumSubsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, true},
		{[]int{1, 1, 3, 4, 7}, true},
		{[]int{2, 3, 4, 6}, false},
		{[]int{2, 4}, false},
		{[]int{2}, false},
		{[]int{}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, dp.EqualSumSubsets(tt.nums))
	}
}
