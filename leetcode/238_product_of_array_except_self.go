package leetcode

/* MEDIUM

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements
of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

Constraints:
	2 <= nums.length <= 10^5
	-30 <= nums[i] <= 30
	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.


Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


SOLUTION: Time O(N) | Space O(1) excluding space of result array
	- create a result array and set first element to 1
	- from number at i = 1 till the end:
		- set result[i] = result[i-1] * nums[i-1]

	- set a running product to 1
	- from number at i = n-1 till the first number:
		- result[i] = result[i] * running product
		- set running product = running product x nums[i]

	- return result
*/

func ProductOfArrayExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = 1

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	runningProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * runningProduct
		runningProduct = runningProduct * nums[i]
	}

	return result
}
