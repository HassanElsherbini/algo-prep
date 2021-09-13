package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestIncrementInteger(t *testing.T) {
	tests := []struct {
		num      []int
		expected []int
	}{
		{num: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{num: []int{9, 9, 9, 9}, expected: []int{1, 0, 0, 0, 0}},
		{num: []int{1, 9, 9, 9}, expected: []int{2, 0, 0, 0}},
		{num: []int{1, 9, 8, 9}, expected: []int{1, 9, 9, 0}},
		{num: []int{0, 0, 0, 9}, expected: []int{0, 0, 1, 0}},
		{num: []int{0, 0, 2, 0}, expected: []int{0, 0, 2, 1}},

		{num: []int{0}, expected: []int{1}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.IncrementInteger(test.num))
	}
}
