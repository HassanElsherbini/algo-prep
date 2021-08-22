package subsets

/*EASY

SOLUTION: Time O(2^N * N) | Space O(2^N * N)
- create a 2d array and add one empty set
- for each number in array:
	- create new subsets by adding the number to each existing subset in the results array
	- add each new subet to the results array
*/
func DistinctSubsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		numOfSubsetsToCreate := len(result)
		for i := 0; i < numOfSubsetsToCreate; i++ {
			// newSubSet := make([]int, len(result[i]))
			// copy(newSubSet, result[i])
			// newSubSet = append(newSubSet, num)
			newSubSet := append([]int(nil), result[i]...)
			newSubSet = append(newSubSet, num)
			result = append(result, newSubSet)
		}
	}
	return result
}
