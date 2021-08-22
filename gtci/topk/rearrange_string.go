package topk

import (
	"strings"

	"github.com/HassanElsherbini/algo-prep/common"
)

/* HARD

SOLUTION: Time O(N * LOG(N)) | Space O(N)
- store char frequencies in hashmap
- put all char in a max-frequency heap
- while heap is not empty:
	- get most frequent char (heap root)
	- add 1 occurence of it to result string
	- decrement its frequency
	- add prev most frequent char to heap if its new frequency is > zero
	- set the currently removed most frequent char as the previous most frequent

- if result string length = input string length:
	- return result string
- else return empty string
*/

func RearrangeString(str string) string {
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

	maxHeap := common.NewBinaryHeap(comparator)
	for _, char := range uniqueChars {
		maxHeap.Insert(char)
	}

	var result strings.Builder
	var prevMostFrequentChar rune
	for maxHeap.Size() > 0 {
		mostFrequentChar := maxHeap.Extract().(rune)
		result.WriteRune(mostFrequentChar)
		charFrequency[mostFrequentChar]--

		if prevMostFrequentChar != 0 && charFrequency[prevMostFrequentChar] > 0 {
			maxHeap.Insert(prevMostFrequentChar)
		}
		prevMostFrequentChar = mostFrequentChar
	}

	if result.Len() == len(str) {
		return result.String()
	}

	return ""
}
