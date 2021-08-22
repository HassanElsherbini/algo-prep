package cyclicsort

/* EASY

SOLUTION:
	- using cyclic sort:
		- if current value has been placed in the right positon, and that position
		is not the current index then the current value is a duplicate.

*/

func FindDuplicateNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		value := nums[i]
		if nums[value-1] != value {
			nums[value-1], nums[i] = value, nums[value-1]
		} else if value-1 != i {
			return value
		} else {
			i++
		}
	}

	return -1
}
