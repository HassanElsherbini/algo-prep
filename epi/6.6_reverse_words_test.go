package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{str: "Alice likes Bob", expected: "Bob likes Alice"},
		{str: "Memory isn't expensive anymore", expected: "anymore expensive isn't Memory"},
		{
			str: "When storage is allocated for a variable, either through a declaration or a call of new, or when a new value " +
				"is created, either through a composite literal or a call of make, and no explicit initialization is provided, " +
				"the variable or value is given a default value. Each element of such a variable or value is set to the zero " +
				"value for its type: false for booleans, 0 for numeric types, \"\" for strings, and nil for pointers, " +
				"functions, interfaces, slices, channels, and maps.",

			expected: "maps. and channels, slices, interfaces, functions, pointers, for nil and strings, for \"\" types, " +
				"numeric for 0 booleans, for false type: its for value zero the to set is value or variable a such of element Each " +
				"value. default a given is value or variable the provided, is initialization explicit no and make, of call a or " +
				"literal composite a through either created, is value new a when or new, of call a or declaration a through either " +
				"variable, a for allocated is storage When",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.ReversWords(test.str))
	}
}
