package slowfastpointer_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/slowfastpointer"
	"github.com/stretchr/testify/assert"
)

func TestFindHappyNumber(t *testing.T) {
	tests := []struct {
		num      int
		expected bool
	}{
		{23, true},
		{12, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, slowfastpointer.FindHappyNumber(tt.num))
	}

}
