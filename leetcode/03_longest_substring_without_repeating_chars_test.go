package leetcode_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestLongestSubstringWithoutRepeatingChars(t *testing.T) {
	tests := []struct {
		str      string
		expected int
	}{
		{"abcabcbb", 3},
		{"abcadcbb", 4},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, leetcode.LongestSubstringWithOutRepeatingChars(tt.str))
	}
}
