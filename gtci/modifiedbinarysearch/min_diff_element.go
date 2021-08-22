package modifiedbinarysearch

/* MEDIUM

SOLUTION: Time O(LOG(N)) | Space O(1)
- if key < fisrt element: return first element
- if key > last element: return last element
- while start <= end:
	- calculate mid
	- if number at mid == key: return key
	- if number at mid > key: start = mid+1
	- else end = mid-1

- if element at start - key  > key - element at end: return element at start

- return element at end
*/

func SearchMindDiffElement(nums []int, key int) int {
	if key < nums[0] {
		return nums[0]
	}

	n := len(nums)
	if key > nums[n-1] {
		return nums[n-1]
	}

	var start int
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2
		if key == nums[mid] {
			return key
		}

		if key > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if (nums[start] - key) < (key - nums[end]) {
		return nums[start]
	}

	return nums[end]
}
