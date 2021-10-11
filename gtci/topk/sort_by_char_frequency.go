package topk

import (
	"strings"

	"github.com/HassanElsherbini/algo-prep/common/binaryheap"
)

/* MEDIUM

SOLUTION: Time O(N + N * LOG(N)) | Space O(N)
	- add all chars to a frequency hashmap where char --> frequency
	- add all chars to a max heap using frequency as the priority
	- untill heap is empty:
		- remove most frequent char in heap (the root)
		- add char to result string a number of times equal to its frequency
	- return result string
*/

func SortByCharFrequency(str string) string {
	charFrequency := make(map[rune]int)
	uniqueChars := []rune{}

	for _, char := range str {
		if _, ok := charFrequency[char]; ok {
			charFrequency[char] += 1
		} else {
			charFrequency[char] = 1
			uniqueChars = append(uniqueChars, char)
		}
	}

	comparator := func(a, b interface{}) int {
		assertedA := a.(rune)
		assertedB := b.(rune)
		if charFrequency[assertedA] > charFrequency[assertedB] {
			return 1
		}
		if charFrequency[assertedA] < charFrequency[assertedB] {
			return -1
		}
		return 0
	}

	maxHeap := binaryheap.New(comparator)
	for _, char := range uniqueChars {
		maxHeap.Insert(char)
	}

	var result strings.Builder
	for maxHeap.Size() > 0 {
		char := maxHeap.Extract().(rune)
		result.WriteString(strings.Repeat(string(char), charFrequency[char]))
	}

	return result.String()
}
