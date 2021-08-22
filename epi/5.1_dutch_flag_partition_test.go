package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestDutchFlagPartition(t *testing.T) {
	tests := []struct {
		nums     []int
		pivot    int
		expected []int
	}{
		{nums: []int{0, 1, 2, 0, 2, 1, 1}, pivot: 2, expected: []int{0, 1, 1, 0, 1, 2, 2}},
		{nums: []int{0, 1, 2, 0, 2, 1, 1}, pivot: 3, expected: []int{0, 0, 2, 1, 2, 1, 1}},
		{nums: []int{3, -2, 10, 7, 8, 2, 4, 4, 1, 9}, pivot: 6, expected: []int{3, -2, 1, 2, 4, 4, 7, 8, 10, 9}},
		{nums: []int{9, 9, 9, 9, 9, 9, 9}, pivot: 3, expected: []int{9, 9, 9, 9, 9, 9, 9}},
		{nums: []int{0, 1, 2, 3, 4, 5, 6}, pivot: 3, expected: []int{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, test := range tests {
		epi.DutchFlagPartition(test.nums, test.pivot)
		assert.Equal(t, test.expected, test.nums)
	}
}
