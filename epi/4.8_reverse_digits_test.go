package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestReverseDigits(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{-314, -413},
		{42, 24},
		{9, 9},
		{-1, -1},
		{0, 0},
		{999999999999999919, 919999999999999999},
		{-999999999999999919, -919999999999999999},
	}

	t.Log("Running tests: multi-step solution")
	for _, test := range tests {
		assert.Equal(t, test.expected, epi.ReverseDigits(test.num))
	}

	t.Log("Running tests: simplified solution")
	for _, test := range tests {
		assert.Equal(t, test.expected, epi.ReverseDigitsSimplified(test.num))
	}
}
