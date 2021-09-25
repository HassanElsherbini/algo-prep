package epi

// 5.5 Delete Duplicates From A Sorted Array

/* SOLUTION: Time O(N) | Space O(1)
- using two pointers, lastPlaced: pointing at last placed unique number. iterator: pointing at current number
- if current number != last Placed number:
	- place current number infront of last Placed number
	- increment lastPlaced
- increment iterator

*/
func DeleteDuplicatesFromSortedArray(nums []int) int {
	lastPlaced := 0
	for iterator := 1; iterator < len(nums); iterator++ {
		if nums[iterator] != nums[lastPlaced] {
			nums[lastPlaced+1] = nums[iterator]
			lastPlaced++
		}
	}

	return lastPlaced
}
