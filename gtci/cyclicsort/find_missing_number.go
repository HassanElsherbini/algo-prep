package cyclicsort

// EASY

/*
   SOLUTION 1:
     Given that all numbers in the array from 0 to N are distinct and there is only one missing number
     then the missing number must equal the diiference between the sum of numbers form (1 to N) and the sum of the numbers in the array
*/

/*
   SOLUTION 2:
     - use cyclic sort to sort the numbers
     - every value at this point is mapped to an index in the array: arr[index] = index
       the first index that doesn't hold a number equal to it is the missing number
*/

func FindMissingNumber(nums []int) int {
	var actualSum int
	expectedSum := (len(nums) + 1) * (len(nums) / 2)

	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}
