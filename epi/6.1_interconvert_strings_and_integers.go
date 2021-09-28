package epi

import (
	"strings"

	"github.com/HassanElsherbini/algo-prep/common"
)

// 6.1 - Interconvert strings and integers

/* STRING TO INTEGER
----------------------------
SOLUTION: Time O(N) | Space O(1)
- Determine the index of the most significant digit: 0 if number is positive, 1 if negative
- from the the index of the most significant digit till the end:
	- convert the digit char to int: '0' - c
	- multiply result by 10 and add the converted digit to the result
- negate result if number is negative

- return result
*/
func StringToInt(num string) int {
	result := 0
	start := 0
	if num[0] == '-' {
		start = 1
	}

	for i := start; i < len(num); i++ {
		convertedDigit := num[i] - '0'
		result = result*10 + int(convertedDigit)
	}

	if num[0] == '-' {
		return result * -1
	}

	return result
}

/* INTEGER TO STRING
----------------------------
SOLUTION: Time O(N) | Space O(1) excluding space for result
- set number to its absolute value
- while number > 0:
	- extract least significant digit: num % 10
	- add digit to result string
	- shift num to the right: num / 10

- reverse result
- if number is negative prepend '-' to result
- return result
*/
func IntegerToString(num int) string {
	var result strings.Builder
	isNegative := (num < 0)
	num = common.Abs(num)

	for ok := true; ok; ok = num > 0 {
		result.WriteByte(byte(num%10) + '0')
		num /= 10
	}

	outPut := result.String()
	result.Reset()
	if isNegative {
		result.WriteByte('-')
	}

	for i := len(outPut) - 1; i >= 0; i-- {
		result.WriteByte(outPut[i])
	}

	return result.String()
}
