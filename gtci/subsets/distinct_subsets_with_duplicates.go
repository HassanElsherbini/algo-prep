package subsets

/*EASY

SOULTION: Time O(2^N * N) | Space O(2^N * N)
	- create a 2d array with one empty set
	- for each number in array:
		- if number was visited before (a duplicate):
			- start from the stored index, this index represents the beggining of the
				unique subsets last produced by this number in the result array.
				(this guarantess that we don't recreate subsets already created by this number)
			- else start from the beggining of the results array
		- create the new subsets and add them to the result array

*/

func DistinctSubsetsWithDuplicates(nums []int) [][]int {
	result := [][]int{{}} // [] [1] [3] [1,3] [3,3] [1,3,3]
	visitedNums := make(map[int]int)
	// 1 -> 1, 3 -> 4
	for _, num := range nums {
		startIndx := visitedNums[num]
		numOfSubsetsToCreate := len(result)

		for i := startIndx; i < numOfSubsetsToCreate; i++ {
			newSubSet := append([]int(nil), result[i]...)
			newSubSet = append(newSubSet, num)
			result = append(result, newSubSet)
		}
		visitedNums[num] = numOfSubsetsToCreate
	}

	return result
}
