package twopointers_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/twopointers"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 3, 3, 3, 6, 9, 9}, 4},
		{[]int{2, 2, 2, 11}, 2},
	}

	for _, tt := range tests {
		assert.Equal(
			t,
			tt.expected,
			twopointers.RemoveDuplicates(tt.nums),
		)
	}
}

func TestSquareSortedArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{-2, -1, 0, 2, 3}, []int{0, 1, 4, 4, 9}},
		{[]int{-3, -1, 0, 1, 2}, []int{0, 1, 1, 4, 9}},
	}

	for _, tt := range tests {
		assert.Equal(
			t,
			tt.expected,
			twopointers.SquareSortedArray(tt.nums),
		)
	}
}
