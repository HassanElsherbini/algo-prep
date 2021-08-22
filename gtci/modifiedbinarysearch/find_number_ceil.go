package modifiedbinarysearch

/* MEDIUM

SOLUTION: Time O(LOG(N)) | Space O(1)
- while start <= end:
	- calculate mid point
	- if target < num at midpoint:
			- set midpoint num as smallest num found so far that is bigger than target
			- set end = mid - 1
	- else set start = mid + 1

	return result
*/

func FindNumberCeil(key int, nums []int) int {
	if key > nums[len(nums)-1] {
		return -1
	}

	var start int
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)/2
		if key == nums[mid] {
			return mid
		}

		if key < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
