package epi

import (
	"math"
	"strings"
	"unicode"
)

// 6.2 - Base Conversion

/* SOLUTION: Time O(N * LOG_newbase(old base)) | Space O(1) excluding result
- convert input to base 10
- while number > 0:
	- get digit of new base number by taking the remainder of (number/new base)
	- add digit to result
	- divide number by new base

- if input is negative:
	- add '-' to result
-reverse result
-return result
*/
func ConvertBase(num string, b1, b2 int) string {
	if b1 == b2 {
		return num
	}
	num = strings.ToUpper(num)

	var result strings.Builder

	isNegative := num[0] == '-'
	if isNegative {
		num = num[1:]
	}

	numInBase10 := 0
	for i := len(num) - 1; i >= 0; i-- {
		multiplier := int(math.Pow(float64(b1), float64(len(num)-1-i)))
		if unicode.IsNumber(rune(num[i])) {
			numInBase10 += int(num[i]-'0') * multiplier
		} else {
			numInBase10 += (int(num[i]-'A') + 10) * multiplier
		}
	}

	for numInBase10 > 0 {
		if numInBase10%b2 >= 10 {
			result.WriteByte('A' + byte(numInBase10%b2-10))
		} else {
			result.WriteByte('0' + byte(numInBase10%b2))
		}
		numInBase10 /= b2
	}

	if isNegative {
		result.WriteByte('-')
	}

	return reverseString(result.String())
}

func reverseString(str string) string {
	var result strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return result.String()
}
