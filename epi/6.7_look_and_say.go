package epi

import "strings"

// 6.7 - Look and Say

/* SOLUTION: Time ~ O(N * 2^N) | Space O(1)
- create a result string and set it to "1"
- from 1 to n:
	- create a new sequence entry string
	- for each digit in result:
		- count its consecutive repetitions
		- write the digit + its count to new sequence entry string
	- set result to new sequence entry

- return result
*/
func LookAndSay(n int) string {
	if n < 1 {
		return ""
	}

	result := "1"
	var nextSequenceEntry strings.Builder

	for i := 1; i < n; i++ {
		for j := 0; j < len(result); {
			count := countConsecutiveRepetitions(result, j)
			nextSequenceEntry.WriteByte('0' + byte(count))
			nextSequenceEntry.WriteByte(result[j])
			j += count
		}
		result = nextSequenceEntry.String()
		nextSequenceEntry.Reset()
	}

	return result
}

func countConsecutiveRepetitions(str string, i int) int {
	count := 0
	for j := i; j < len(str) && str[j] == str[i]; j++ {
		count++
	}

	return count
}
