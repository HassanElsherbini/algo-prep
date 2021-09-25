package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{nums: []int{1}, expected: []int{1}},
		{nums: []int{1, 1, 2, 2, 3, 3}, expected: []int{1, 2, 3}},
		{nums: []int{2, 2, 2, 2, 2}, expected: []int{2}},
		{nums: []int{2, 3, 5, 5, 7, 11, 11, 11, 13}, expected: []int{2, 3, 5, 7, 11, 13}},
		{nums: []int{1, 3, 5, 7, 50, 12, 113, 10}, expected: []int{1, 3, 5, 7, 50, 12, 113, 10}},
	}

	for _, test := range tests {
		actual := epi.DeleteDuplicatesFromSortedArray(test.nums)
		assert.Equal(t, test.expected, test.nums[0:actual+1])
	}
}
