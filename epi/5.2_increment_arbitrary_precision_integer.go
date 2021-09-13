package epi

import "math"

// 5.2 Increment arbitrary precision integer

/* SOLUTION: Time O(N) | Space(1) (excluding space for result)
- set carry to 1
- while there is digits left && a carry:
	- calculate sum: current digit + carry
	- set current digit to sum % 10
	- set carry to sum / 10

- if carry > 0 (a carry remains):
	- append 0 to the num array
	- set first digit to 1

- return num array
*/
func IncrementInteger(num []int) []int {
	carry := 1
	for i := len(num) - 1; i >= 0 && carry > 0; i-- {
		sum := num[i] + carry
		carry = sum / 10
		num[i] = sum % 10
	}

	if carry > 0 {
		num = append(num, 0)
		num[0] = carry
	}

	return num
}

/* NAIVE SOLUTION (doesn't cover integer overflow edge case):
- recreate the number as an integer
- add one to the number
- write the new number into the num array starting from the least significant digit
- if the new number digits exceeds the array, then add 0 to the end of the array
	and set the first digit in the array to 1

- return num array
*/
func IncrementIntegerNaive(num []int) []int {
	value := 0
	decimalPlace := int(math.Pow(10, float64(len(num)-1)))

	for _, digit := range num {
		value += digit * decimalPlace
		decimalPlace /= 10
	}

	value += 1
	for i := len(num) - 1; i >= 0; i-- {
		num[i] = value % 10
		value /= 10
	}

	if value > 0 {
		num = append(num, 0)
		num[0] = value
	}

	return num
}
