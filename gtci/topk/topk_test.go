package topk_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/common"
	"github.com/HassanElsherbini/algo-prep/gtci/topk"
	"github.com/stretchr/testify/assert"
)

func TestFindKLargestNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{3, 1, 5, 12, 2, 11}, 3, []int{5, 12, 11}},
		{[]int{5, 12, 11, -1, 12}, 3, []int{11, 12, 12}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindKLargestNumbers(tt.nums, tt.k))
	}
}

func TestFindKSmallestNumber(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 5, 12, 2, 11, 5}, 3, 5},
		{[]int{1, 5, 12, 2, 11, 5}, 4, 5},
		{[]int{5, 12, 11, -1, 12}, 3, 11},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindKSmallestNumber(tt.nums, tt.k))
	}
}

func TestFindKClosestPoints(t *testing.T) {
	tests := []struct {
		points   []common.Point2D
		k        int
		expected []common.Point2D
	}{
		{[]common.Point2D{{1, 2}, {1, 3}}, 1, []common.Point2D{{1, 2}}},
		{[]common.Point2D{{1, 3}, {3, 4}, {2, -1}}, 2, []common.Point2D{{1, 3}, {2, -1}}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindKClosestPoints(tt.points, tt.k))
	}
}

func TestMinCostToConnectRopes(t *testing.T) {
	tests := []struct {
		ropes    []int
		expected int
	}{
		{[]int{1, 3, 11, 5}, 33},
		{[]int{3, 4, 5, 6}, 36},
		{[]int{1, 3, 11, 5, 2}, 42},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.MinCostToConnectRopes(tt.ropes))
	}
}

func TestFindKFrequentNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 3, 5, 12, 11, 12, 11}, 2, []int{12, 11}},
		{[]int{5, 12, 11, 3, 11}, 2, []int{11, 3}},
		{[]int{5, 5, 5, 5, 12, 11, 30, 30, 11, 11}, 3, []int{5, 11, 30}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindKFrequentNumbers(tt.nums, tt.k))
	}
}

func TestSortByCharFrequency(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{"Programming", "rrggmmoaPin"},
		{"abcbab", "bbbaac"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.SortByCharFrequency(tt.str))
	}
}

func TestFindClosestNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		x        int
		expected []int
	}{
		{[]int{5, 6, 7, 8, 9}, 3, 3, []int{5, 6, 7}},
		{[]int{2, 4, 5, 6, 9}, 2, 5, []int{5, 6}},
		{[]int{2, 4, 5, 6, 9}, 3, 13, []int{5, 6, 9}},
		{[]int{2, 4, 5, 6, 9}, 8, 4, nil},
	}

	t.Logf("Running min-heap solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindClosestNumbers(tt.nums, tt.k, tt.x))
	}

	t.Logf("Running two-pointers solution tests")
	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindClosestNumbersUsingTwoPointers(tt.nums, tt.k, tt.x))
	}
}

func TestFindMaxDistinctNumbers(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{7, 3, 5, 8, 5, 3, 3}, 2, 3},
		{[]int{3, 5, 12, 11, 12}, 3, 2},
		{[]int{1, 2, 3, 3, 3, 3, 4, 4, 5, 5, 5}, 2, 3},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.FindMaxDistinctNumbers(tt.nums, tt.k))
	}
}

func TestRearrangeString(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{"aappp", "papap"},
		{"Programming", "rgmrgmoaPin"},
		{"aapa", ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, topk.RearrangeString(tt.str))
	}
}
