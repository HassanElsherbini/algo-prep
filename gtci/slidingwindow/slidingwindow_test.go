package slidingwindow_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/slidingwindow"
	"github.com/stretchr/testify/assert"
)

func TestMaxSumSubArrayOfSizeK(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{2, 3, 4, 1, 5}, 2, 7},
	}

	for _, tt := range tests {
		assert.Equal(
			t,
			tt.expected,
			slidingwindow.MaxSumSubArrayOfSizeK(tt.nums, tt.k),
		)
	}
}

func TestLongestSubStringWithKDistinct(t *testing.T) {
	tests := []struct {
		str      string
		k        int
		expected int
	}{
		{"araaci", 2, 4},
		{"araaci", 1, 2},
		{"cbbebi", 3, 5},
	}

	for _, tt := range tests {
		assert.Equal(
			t,
			tt.expected,
			slidingwindow.LongestSubStringWithKDistinct(tt.str, tt.k),
		)
	}
}
