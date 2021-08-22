package leetcode

import "github.com/HassanElsherbini/algo-prep/common"

/* MEDIUM
Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product, and return the product.

It is guaranteed that the answer will fit in a 32-bit integer.

A subarray is a contiguous subsequence of the array.

Example 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:
	1 <= nums.length <= 2 * 10^4
	-10 <= nums[i] <= 10
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.



SOLUTION 1 (naive solution): Time O(N^2) | Space O(1):
- intiliaze variable max = -infinity
- for each number a:
	- for each number b in reamining numbers:
		- if a * b > max:
			- max = a * b

- return max
*/

/*
SOLUTION 2 (keeping track of min & max products): Time O(N) | Space O(1)
- at each number can be -/+ we need to evaluate:
	- if a previous negative product x number = new max
	- if a previous min product x number = new max  (covers the case of prev negative products x negative number)
	- if number alone = new max

- initialize a prevMax, prevMin, result to first number
- from 2nd number till the end:
	- potential max: Max(prevMax * number, prevMin * number, number)
	- potential min: Min(prevMax * number, prevMin * number, number)
		(keeps track of smallest negative product of current subarray to cover case of: negative product x negative number)

	- result = Max(result, potential max)
	- prev max= potential max
	- prev min = potential min

- return result
*/
func MaxProductSubArrayDP(nums []int) int {
	prevMax := nums[0]
	prevMin := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curMax := common.Max(common.Max(prevMax*num, prevMin*num), num)
		curMin := common.Min(common.Min(prevMax*num, prevMin*num), num)

		result = common.Max(result, curMax)
		prevMax = curMax
		prevMin = curMin
	}
	return result
}

/*
SOLUTION 3 (realizing mid subarrays can be ignored): Time O(N) | Space O(1)
- if we break the set of potential subarrays into:
	- left: a subarray that starts from the begining of the array
	- mid: a subarray that starts and ends some where in the middle of the array
	- right: a subarray that ends at the end of the array

  and look at all the possible scenarios given that break down:
	left   mid  right   solution       (+/- indicates a positive or negative subarray product)
	-----------------------------------
	  +     +     +     left x mid x right
	  +     +     -     left x mid
	  -     +     -     left x mid x right
	  -     +     +     mid x right

	  -     -     -     left x mid OR  mid x right
	  -     -     +     left x mid x right
	  +     -     +     left OR right
	  +     -     -     left x mid x right

	we realize that all solutions can be produced by left or right sub arrays

	given that intuition our solution simply becomes:
	- initialize "result" to first number
	- intialize "leftProduct" and "rightProduct" to 1
	- for i = 0 to n-1:
		- set leftProduct or rightProduct to 1 if their value is 0 (need to drop current sub array)
		- calculate left product: leftProduct * nums[i]
		- calculate right product: rightProduct * nums[n-1-i]
		- update result: Max(result, leftProduct, rightProduct)

	- return result
*/
func MaxProductSubArray(nums []int) int {
	result := nums[0]
	leftProduct := 1
	rightProduct := 1

	for i := range nums {
		leftProduct *= nums[i]
		rightProduct *= nums[len(nums)-1-i]
		result = common.Max(result, common.Max(leftProduct, rightProduct))
		if leftProduct == 0 {
			leftProduct = 1
		}
		if rightProduct == 0 {
			rightProduct = 1
		}
	}
	return result
}
