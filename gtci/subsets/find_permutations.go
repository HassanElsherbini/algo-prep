package subsets

/* MEDIUM

SOLUTION DFS: Time O(N * N!) | Spce O(N * N!)
 - if array is empty then add running permutation to result and return
 - for each num in array:
	- add it to permutation
	- create new array excluding that number and recurse using that new array


SOLUTION BFS: Time O(N * N!) | Spce O(N * N!)
	*this a bottom-up approach where all permutations are built starting with an empty array

 - create a result array with a single empty permutation array
 - for each number (num):
	- for each permutation built so far in result array:
		- for each position in permutation array:
			- create a new permutation out of inserting num in that position
	- set result = to all those newly created set of permutations

	for example: [1,3,5]
		[]
		[1]
		[3,1] [1,3]
		[5,3,1] [3,5,1] [3,1,5]  [5,1,3] [1,5,3] [1,3,5] --> final set of permutations
*/

func FindPermutationsDfs(nums []int) [][]int {
	return findPermutations(nums, []int{}, [][]int{})
}

func findPermutations(nums []int, permutation []int, result [][]int) [][]int {
	if len(nums) == 1 {
		permutation = append(permutation, nums[0])
		result = append(result, permutation)
		return result
	}

	for i, num := range nums {
		newPerm := append([]int(nil), permutation...)
		newPerm = append(newPerm, num)
		remainingNums := append([]int(nil), nums[:i]...)
		remainingNums = append(remainingNums, nums[i+1:]...)

		result = findPermutations(remainingNums, newPerm, result)
	}

	return result
}

func FindPermutationsBfs(nums []int) [][]int {
	result := [][]int{{}}
	for _, num := range nums {
		newPermutations := [][]int{}
		for i := 0; i < len(result); i++ {
			newPermutations = append(newPermutations, insertInEveryPosition(result[i], num)...)
		}
		result = newPermutations
	}
	return result
}

func insertInEveryPosition(nums []int, num int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		result = append(result, []int{num})
		return result
	}

	for i := 0; i <= len(nums); i++ {
		r := append([]int(nil), nums[:i]...)
		r = append(r, num)
		r = append(r, nums[i:]...)
		result = append(result, r)
	}

	return result
}
