package modifiedbinarysearch

/* MEDIUM

SOLUTION: Time O(LOG(N)) | Space O(1)
	- get earliest instance of key using binary search
	- get last instance of key using binary search
	- return range
*/

func FindNumberRange(nums []int, key int) [2]int {
	result := [2]int{-1, -1}

	result[0] = getFirstInstance(nums, key)
	if result[0] != -1 {
		result[1] = getLastInstance(nums, key)
	}

	return result
}

func getFirstInstance(nums []int, key int) int {
	result := -1
	n := len(nums)
	var start int
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2
		if key > nums[mid] {
			start = mid + 1
		} else {
			if key == nums[mid] {
				result = mid
			}
			end = mid - 1
		}
	}
	return result
}

func getLastInstance(nums []int, key int) int {
	result := -1
	n := len(nums)
	var start int
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2
		if key < nums[mid] {
			end = mid - 1
		} else {
			if key == nums[mid] {
				result = mid
			}
			start = mid + 1
		}
	}

	return result
}
