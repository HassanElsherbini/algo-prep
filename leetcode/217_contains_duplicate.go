package leetcode

/* EASY

Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true

Example 2:

Input: nums = [1,2,3,4]
Output: false

SOLUTION 1: Time O(N * LOG(N)) | Space O(1)
- sort array
- for each number in array:
	- if number is equal to the next number:
		- return true

- return false

SOLUTION 2: Time O(N) | Space O(N)
- create a hashmap for storing visited numbers
- for each number in array:
	- if number exists in hashmap:
		- return true
	- store number in hashmap

- return false
*/

// SOLUTION 2
func ContainsDuplicate(nums []int) bool {
	visitedNums := make(map[int]bool)
	for _, num := range nums {
		if visitedNums[num] {
			return true
		}
		visitedNums[num] = true
	}

	return false
}
