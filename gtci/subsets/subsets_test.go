package subsets_test

import (
	"testing"

	"github.com/HassanElsherbini/algo-prep/gtci/subsets"
	"github.com/stretchr/testify/assert"
)

// func TestDistinctSubsets(t *testing.T) {
// 	tests := []struct {
// 		nums     []int
// 		expected [][]int
// 	}{
// 		{[]int{1, 3}, [][]int{{}, {1}, {3}, {1, 3}}},
// 		{[]int{1, 5, 3, 4}, [][]int{{}, {1}, {5}, {1, 5}, {3}, {1, 3}, {5, 3}, {1, 5, 3}}},
// 	}

// 	for _, tt := range tests {
// 		assert.Equal(t, tt.expected, subsets.DistinctSubsets(tt.nums))
// 	}
// }

func TestDistinctSubsetsWithDuplicates(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{1, 3, 3},
			[][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}},
		},
		{
			[]int{1, 5, 3, 3},
			[][]int{{}, {1}, {5}, {1, 5}, {3}, {1, 3}, {5, 3}, {1, 5, 3}, {3, 3}, {1, 3, 3}, {5, 3, 3}, {1, 5, 3, 3}},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, subsets.DistinctSubsetsWithDuplicates(tt.nums))
	}
}

func TestFindPermutations(t *testing.T) {
	testsDfs := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			"DFS approach",
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			"DFS approach",
			[]int{1, 3, 5},
			[][]int{{1, 3, 5}, {1, 5, 3}, {3, 1, 5}, {3, 5, 1}, {5, 1, 3}, {5, 3, 1}},
		},
	}

	for i, tt := range testsDfs {
		t.Logf("Running test: %s #%d", tt.name, i+1)
		assert.Equal(t, tt.expected, subsets.FindPermutationsDfs(tt.nums), tt.name)
	}

	testsBfs := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			"BFS approach",
			[]int{1, 2, 3},
			[][]int{{3, 2, 1}, {2, 3, 1}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {1, 2, 3}},
		},
		{
			"BFS approach",
			[]int{1, 3, 5},
			[][]int{{5, 3, 1}, {3, 5, 1}, {3, 1, 5}, {5, 1, 3}, {1, 5, 3}, {1, 3, 5}},
		},
	}

	for i, tt := range testsBfs {
		t.Logf("Running test: %s #%d", tt.name, i+1)
		assert.Equal(t, tt.expected, subsets.FindPermutationsBfs(tt.nums), tt.name)

	}
}

func TestFindLetterCaseStringPermutations(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
	}{
		{"lower case", "ad52", []string{"ad52", "Ad52", "aD52", "AD52"}},
		{"upper case", "ab7c", []string{"ab7c", "Ab7c", "aB7c", "AB7c", "ab7C", "Ab7C", "aB7C", "AB7C"}},
		{"mixed case", "aB7cD", []string{
			"aB7cD", "AB7cD", "ab7cD", "Ab7cD", "aB7CD", "AB7CD", "ab7CD", "Ab7CD",
			"aB7cd", "AB7cd", "ab7cd", "Ab7cd", "aB7Cd", "AB7Cd", "ab7Cd", "Ab7Cd",
		},
		},
	}

	for _, tt := range tests {
		t.Logf("Running test: %s", tt.name)
		assert.Equal(t, tt.expected, subsets.FindLetterCaseStringPermutations(tt.str))
	}
}

func TestGenerateBalancedParanthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, subsets.GenerateBalancedParanthesis(tt.n))
	}
}
