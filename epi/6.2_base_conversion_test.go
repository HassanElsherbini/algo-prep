package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestConvertBase(t *testing.T) {
	tests := []struct {
		num      string
		b1       int
		b2       int
		expected string
	}{
		{num: "615", b1: 7, b2: 13, expected: "1A7"},
		{num: "-615", b1: 7, b2: 13, expected: "-1A7"},
		{num: "1A7", b1: 13, b2: 7, expected: "615"},
		{num: "-1A7", b1: 13, b2: 7, expected: "-615"},

		{num: "1000000", b1: 10, b2: 2, expected: "11110100001001000000"},
		{num: "1000000", b1: 10, b2: 16, expected: "F4240"},

		{num: "15", b1: 7, b2: 7, expected: "15"},
		{num: "1a7", b1: 13, b2: 7, expected: "615"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.ConvertBase(test.num, test.b1, test.b2))
	}
}
