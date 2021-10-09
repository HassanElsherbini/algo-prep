package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		str      string
		expected bool
	}{
		{str: "A man, a plan, a canal, Panama.", expected: true},
		{str: "Able was I, ere I Saw Elba!", expected: true},
		{str: "Ray a Ray", expected: false},

		{str: "A ma ^^^^ <3 <3 **** n, a plan, a $     canal,$%^!@ Pan33ama.", expected: true},
		{str: "", expected: true},
		{str: "a", expected: true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.IsPalindrome(test.str))
	}
}
