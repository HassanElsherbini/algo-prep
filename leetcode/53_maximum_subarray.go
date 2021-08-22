package leetcode

/* EASY

Given an integer array nums, find the contiguous subarray (containing at least one number)
which has the largest sum and return its sum.

Constraints:
	1 <= nums.length <= 3 * 10^4
	-10^5 <= nums[i] <= 10^5


Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:
Input: nums = [1]
Output: 1

SOLUTION: Time O(N) | Space O(1)
- create a "maxSum" to hold the maximum subaray sum found, and a running "sum" to hold the sum of the subarray
	currently being evaluated.
- intialize both to the first number
- for each number "num" starting at i = 1:
	- if sum + num is larger than num:
		- set sum = sum + num
	- else set sum = num
	- if sum > maxSum:
		- set sum as new maxSum

- return maxSum
*/

func MaximumSubArray(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if sum+nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
