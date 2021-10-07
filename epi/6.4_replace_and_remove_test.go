package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestReplaceAndRemove(t *testing.T) {
	tests := []struct {
		chars         []rune
		size          int
		charToRemove  rune
		charToReplace rune
		replaceChar   rune
		expected      []rune
	}{
		{
			chars:         []rune{'a', 'c', 'd', 'b', 'b', 'c', 'a', 0, 0, 0, 0},
			size:          7,
			charToReplace: 'a',
			charToRemove:  'b',
			replaceChar:   'd',
			expected:      []rune{'d', 'd', 'c', 'd', 'c', 'd', 'd'},
		},
		{
			chars:         []rune{'f', 'c', 'a', 'b', 'c', 'd', 'b', 'a', 'b', 0, 0, 0, 0},
			size:          9,
			charToRemove:  'b',
			charToReplace: 'a',
			replaceChar:   'd',
			expected:      []rune{'f', 'c', 'd', 'd', 'c', 'd', 'd', 'd'},
		},
		{
			chars:         []rune{'f', 'c', 'a', 'b', 'c', 'd', 'b', 'a', 'b', 0, 0, 0, 0},
			charToRemove:  'c',
			charToReplace: 'b',
			replaceChar:   'f',
			size:          9,
			expected:      []rune{'f', 'a', 'f', 'f', 'd', 'f', 'f', 'a', 'f', 'f'},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.ReplaceAndRemove(test.chars, test.size, test.charToRemove, test.charToReplace, test.replaceChar))
	}
}
