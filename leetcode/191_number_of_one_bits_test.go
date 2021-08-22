package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestNumberOfOneBits(t *testing.T) {
	tests := []struct {
		num      uint32
		expected int
	}{
		{11, 3},
		{128, 1},
		{4294967293, 31},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.NumberOfOneBits(tt.num))
	}
}
