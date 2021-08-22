package cyclicsort_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/cyclicsort"
	"github.com/stretchr/testify/assert"
)

func TestCyclicSort(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{3, 1, 5, 4, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 6, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 5, 6, 4, 3, 2}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, cyclicsort.CyclicSort(tt.nums))
	}
}

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{4, 0, 3, 1}, 2},
		{[]int{8, 3, 5, 2, 4, 6, 0, 1}, 7},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, cyclicsort.FindMissingNumber(tt.nums))
	}
}

func TestFindAllMissingNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{2, 3, 1, 8, 2, 3, 5, 1}, []int{4, 6, 7}},
		{[]int{2, 4, 1, 2}, []int{3}},
		{[]int{2, 3, 2, 1}, []int{4}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, cyclicsort.FindAllMissingNumbers(tt.nums))
	}
}

func TestFindDuplicateNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 4, 4, 3, 2}, 4},
		{[]int{2, 1, 3, 3, 5, 4}, 3},
		{[]int{2, 4, 1, 4, 4}, 4},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, cyclicsort.FindDuplicateNumber(tt.nums))
	}
}
