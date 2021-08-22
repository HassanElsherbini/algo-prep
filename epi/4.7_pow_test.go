package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	tests := []struct {
		x        float64
		y        int
		expected float64
	}{
		{2, 8, 256},
		{5, 0, 1},
		{3, -3, 1.0 / 27.0},

		{-2, 8, 256},
		{-2, 9, -512},
		{-5, 0, -1},
		{-3, -3, 1.0 / -27.0},
		{-3, -4, 1.0 / 81.0},
	}

	t.Log("Running linear time solution tests")
	for _, test := range tests {
		assert.Equal(t, test.expected, epi.Pow(test.x, test.y))
	}

	t.Log("Running logarithmic time solution tests")
	for _, test := range tests {
		assert.Equal(t, test.expected, epi.Pow2(test.x, test.y))
	}
}
