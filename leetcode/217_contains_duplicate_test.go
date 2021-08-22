package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.ContainsDuplicate(tt.nums))
	}
}
