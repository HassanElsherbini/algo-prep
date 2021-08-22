package epi

/* 4.8 ReverseDigits

Solution 1: Time O(N)  N = number of digits | Space O(1)
	- get decimal place value of most significant digit
	- while num is > 0:
		- get last digit of num -> num % 10
		- multiply digit by decimal place value and add it to result
		- divide decimal place value by 10
		- divide num by 10
	- return result based on input sign
*/

func ReverseDigits(num int) int {
	absNum := abs(num)
	if absNum < 10 {
		return num
	}

	decimalPlaceValue := getDecimalPlaceValue(absNum)

	result := 0
	for absNum > 0 {
		result += (absNum % 10) * decimalPlaceValue
		absNum = absNum / 10
		decimalPlaceValue = decimalPlaceValue / 10
	}

	if num < 0 {
		return -result
	}

	return result
}

func getDecimalPlaceValue(num int) int {
	if num < 10 {
		return 1
	}

	decimalPlaceVal := 10

	for (num / decimalPlaceVal) > 9 {
		decimalPlaceVal *= 10
	}

	return decimalPlaceVal
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/* Solution 2: Time O(N) N = number of digits | Space O(1)
(simplified logic: no need to calculate decimal place value or abs value)
	- while num != 0:
		- multiply result by 10 then add last digit (num%10)
		- divide num by 10
	- return result
*/
func ReverseDigitsSimplified(num int) int {
	result := 0
	for num != 0 {
		result = result*10 + (num % 10)
		num /= 10
	}

	return result
}
