package epi_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/epi"
	"github.com/stretchr/testify/assert"
)

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{n: 0, expected: ""},
		{n: 1, expected: "1"},
		{n: 8, expected: "1113213211"},
		{
			n: 20,
			expected: "11131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112" +
				"22113111231133221121113122113121113221112131221123113111231121123222112132113213221133112132123123112" +
				"1113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, epi.LookAndSay(test.n))
	}
}
