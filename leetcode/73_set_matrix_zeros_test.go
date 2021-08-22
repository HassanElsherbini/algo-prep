package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestSetMatrixZeros(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{0}},
			[][]int{{0}},
		},
	}

	for _, tt := range tests {
		leetcode.SetMatrixZeros(tt.matrix)
		assert.Equal(t, tt.expected, tt.matrix)
	}
}
