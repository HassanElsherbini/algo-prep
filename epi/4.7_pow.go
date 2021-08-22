package epi

import "github.com/HassanElsherbini/algo-prep/common"

/* 4.7 Compute Pow(x,y) - not accounting for overflow and underflow

Examples:
x = 2
y = 8
output = 128

x = 3
y = -3
output = 1/27

x = 5
y = 0
output = 1
*/

/*

SOLUTION 1: Time O(Y) | Space O(1)
	- if y = 0:
			return 1

	- multiply x by itself abs(y) times
	- if y is negative:
		- return 1/result

	- else return result
*/

func Pow(x float64, y int) float64 {
	if y == 0 {
		if x < 0 {
			return -1
		}
		return 1
	}

	result := x
	for i := common.Abs(y); i > 1; i-- {
		result *= x
	}

	if y < 0 {
		return 1 / result
	}

	return result
}

/*
SOLUTION 2: Time O(log(Y)) | Space O(1)
	Using the law of raising exponents to a power: (x^n)^m = x^(n*m),
	we can reduce the # of multiplication operations y to log(y)
	eg.
	x^8 = (x*x)^4 -> (x^2 * x^2)^2 -> x^4 * x^4
	which is 3 operations! instead of 7 (x * x * x * x * x * x * x * x)
	-------------------------------------------------------------

	- while power is > 0:
		- if power is odd:
			- multiply result by x (since we are halfing the power we will miss a multiplication when power is odd)
		- multiply result by x
		- divide power by 2

	- return result
*/
func Pow2(x float64, y int) float64 {
	result := 1.0
	power := y

	if power == 0 {
		if x < 0 {
			return -1
		}
		return 1
	}

	if power < 0 {
		power *= -1
		x = 1 / x
	}

	for power > 0 {
		if power%2 == 1 {
			result *= x
		}
		x *= x
		power = power / 2
	}

	return result
}
