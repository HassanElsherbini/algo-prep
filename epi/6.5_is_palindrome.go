package epi

import (
	"unicode"
)

// 6.5 - Test Palindromicity

/* SOLUTION: Time O(N) | Space O(1)
- set left = 0, set right = index of last character
- while left < right:
	- find next alphanumeric character from left
	- find next alphanumeric character from right
	- if lowercase variant of both characters is not equal:
		- return false
	-increment left
	-decrement right

- return true
*/

func IsPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		for !isAlphaNumeric(rune(str[left])) && left < right {
			left++
		}

		for !isAlphaNumeric(rune(str[right])) && left < right {
			right--
		}

		if unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsNumber(char)
}
