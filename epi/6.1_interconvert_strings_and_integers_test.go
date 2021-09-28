package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestInterconvertStringsAndIntegers(t *testing.T) {
	strToIntTests := []struct {
		num      string
		expected int
	}{
		{num: "314", expected: 314},
		{num: "-314", expected: -314},
		{num: "999999999999999999", expected: 999999999999999999},
		{num: "0", expected: 0},
	}

	t.Log("Running tests: string to integer")
	for _, test := range strToIntTests {
		assert.Equal(t, test.expected, epi.StringToInt(test.num))
	}

	intToStrTests := []struct {
		num      int
		expected string
	}{
		{num: 314, expected: "314"},
		{num: -314, expected: "-314"},
		{num: 999999999999999999, expected: "999999999999999999"},
		{num: 0, expected: "0"},
	}

	t.Log("Running tests: integer to string")
	for _, test := range intToStrTests {
		assert.Equal(t, test.expected, epi.IntegerToString(test.num))
	}
}
