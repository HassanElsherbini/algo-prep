package cyclicsort

/* EASY

SOLUTION:
	- for each number:
		- if value is at right position skip
			else swap
*/

func FindAllMissingNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		value := nums[i]
		if value != nums[value-1] {
			nums[i], nums[value-1] = nums[value-1], value
		} else {
			i++
		}
	}

	result := make([]int, 0)
	for i, num := range nums {
		if num != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}
