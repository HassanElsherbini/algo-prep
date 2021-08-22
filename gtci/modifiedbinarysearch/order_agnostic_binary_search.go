package modifiedbinarysearch

/* EASY

SOLUTION: Time O(Log(N)) | Space O(1)
	- determine if sort order by comparing first element to last element
	- while start <= end:
		- calculate mid = (end - start)/2 + start
		- if target = num at mid:
			- return mid

		- if ascending:
			- if target is < mid num:
				- set end = mid - 1
			- else set start = mid + 1
		- else:
			- if target is > mid num:
				- set end = mid - 1
			- else set start = mid + 1
*/

func OrderAgnosticBFS(target int, nums []int) int {
	var isAscending bool
	var start int
	end := len(nums) - 1

	if nums[start] < nums[end] {
		isAscending = true
	}

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if isAscending {
			if target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target > nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}
