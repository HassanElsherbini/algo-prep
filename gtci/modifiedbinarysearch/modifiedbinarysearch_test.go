package modifiedbinarysearch_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/modifiedbinarysearch"
	"github.com/stretchr/testify/assert"
)

func TestOrderAgnosticBFS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			"ascending",
			[]int{4, 6, 10, 11, 12, 20},
			6,
			1,
		},
		{
			"descending",
			[]int{10, 6, 4, 2, 1},
			2,
			3,
		},
		{
			"not found",
			[]int{1, 2, 3, 40, 43, 79, 81},
			80,
			-1,
		},
		{
			"some duplicates",
			[]int{1, 2, 4, 4, 6, 6, 7, 21, 32},
			7,
			6,
		},
		{
			"all duplicates",
			[]int{1, 1, 1, 1, 1, 1, 1},
			1,
			3,
		},
		{
			"all duplicates, not found",
			[]int{1, 1, 1, 1, 1, 1, 1},
			2,
			-1,
		},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, modifiedbinarysearch.OrderAgnosticBFS(tt.target, tt.nums))
	}
}

func TestFindNumberCeil(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		key      int
		expected int
	}{
		{"ceil of existing key is it self", []int{4, 6, 10}, 6, 1},
		{"ceil is largest number", []int{1, 3, 8, 10, 15}, 12, 4},
		{"ceil is smallest number", []int{4, 6, 10}, -1, 0},
		{"not found", []int{4, 6, 10}, 17, -1},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, modifiedbinarysearch.FindNumberCeil(tt.key, tt.nums))
	}
}

func TestSearchNextLetter(t *testing.T) {
	tests := []struct {
		name     string
		letters  []rune
		key      rune
		expected rune
	}{
		{"key is within letters range", []rune{'a', 'c', 'f', 'h'}, 'b', 'c'},
		{"key exists", []rune{'a', 'c', 'f', 'h'}, 'f', 'h'},
		{"key exceeds letters range", []rune{'a', 'c', 'f', 'h'}, 'm', 'a'},
		{"last letter is equal to key", []rune{'a', 'c', 'f', 'h'}, 'h', 'a'},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, modifiedbinarysearch.SearchNextLetter(tt.letters, tt.key))
	}
}

func TestFindNumberRange(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		key      int
		expected [2]int
	}{
		{"single instance", []int{1, 3, 8, 10, 15}, 10, [2]int{3, 3}},
		{"mid overlap", []int{4, 6, 6, 6, 9}, 6, [2]int{1, 3}},
		{"left overlap", []int{6, 6, 6, 4, 5}, 6, [2]int{0, 2}},
		{"right overlap", []int{4, 5, 6, 6, 6}, 6, [2]int{2, 4}},
		{"full overlap", []int{1, 1, 1, 1, 1}, 1, [2]int{0, 4}},
		{"not found", []int{1, 3, 8, 10, 15}, 12, [2]int{-1, -1}},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, modifiedbinarysearch.FindNumberRange(tt.nums, tt.key))
	}
}

func TestSearchMindDiffElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		key      int
		expected int
	}{
		{"key is with in range of two existing elements", []int{1, 3, 8, 10, 15}, 9, 8},
		{"key exists", []int{4, 6, 10}, 4, 4},
		{"key is larger than all existing elements", []int{4, 6, 10}, 12, 10},
		{"key is smaller than all existing elements", []int{4, 6, 10}, 2, 4},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, modifiedbinarysearch.SearchMindDiffElement(tt.nums, tt.key))
	}
}
