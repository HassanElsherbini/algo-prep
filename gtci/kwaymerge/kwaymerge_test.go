package kwaymerge_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/kwaymerge"
	"github.com/stretchr/testify/assert"
)

func TestMergeSortedLists(t *testing.T) {
	tests := []struct {
		lists    [][]int
		expected []int
	}{
		{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, []int{1, 2, 3, 3, 4, 6, 6, 7, 8}},
		{[][]int{{5, 8, 9}, {1, 7}}, []int{1, 5, 7, 8, 9}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, kwaymerge.MergeSortedLists(tt.lists))
	}
}

func TestFindKSmallestNumber(t *testing.T) {
	tests := []struct {
		nums     [][]int
		k        int
		expected int
	}{
		{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, 5, 4},
		{[][]int{{5, 8, 9}, {1, 7}}, 3, 7},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, kwaymerge.FindKSmallestNumber(tt.nums, tt.k))
	}
}

func TestFindKSmallestInSortedMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		k        int
		expected int
	}{
		{[][]int{{2, 6, 8}, {3, 7, 10}, {5, 8, 11}}, 5, 7},
		{[][]int{{2, 3, 7}, {4, 7, 10}, {8, 8, 11}}, 4, 7},
		{[][]int{{2, 3, 7}, {4, 7, 10}, {8, 8, 11}}, 3, 4},
		{[][]int{{8, 8, 8}, {8, 9, 10}}, 4, 8},

		{[][]int{{2, 4}, {3, 5}}, 1, 2},
		{[][]int{{2, 4}, {3, 5}}, 4, 5},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, kwaymerge.FindKSmallestInSortedMatrix(tt.matrix, tt.k))
	}
}
