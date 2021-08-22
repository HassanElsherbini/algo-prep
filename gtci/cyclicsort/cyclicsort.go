package cyclicsort

// EASY

// SOlUTION 1
/*
   for each number:
     if number is in right position (index == number):
       move to next number
     else
       swap it with the number currently occupying its position
*/

// SOLUTION 2 (no cyclic sort)
/*
   an array with values from 1 to N (length of array) is guaranteed to be sorted after N-1
	 values are placed in their corresponding index - (value - 1)
   while required number of placements > 0:
     if current value is in the right place:
       skip
     else
       place value in right position
       set next value to visit = to the value found at position

     decrement # of remaining placements to be done
*/

func CyclicSort(arr []int) []int {
	nums := make([]int, len(arr))
	copy(nums, arr)

	i := 0
	for i < len(nums) {
		value := nums[i]
		if value == i+1 {
			i++
		} else {
			nums[i], nums[value-1] = nums[value-1], nums[i]
		}
	}

	return nums
}
