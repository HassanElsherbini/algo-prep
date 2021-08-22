package dp

/* MEDIUM

SOLUTION (dp bottom-up): Time O(N * S) | Space O(N * S) S = sum of all numbers
	- to have two subsets of equal sum, each subset sum must equal half the sum of all the numbers
	- this trasnforms the problem to: does a set of numbers that add up to S/2 exist?
	- we can find the solution by answering this question for every possible sum (0 --> S/2)
		and set variation ({}, {1}, {1, 2}, {1, 2, 3}, {1, 2, 3, 4})
		---------------------------------------------------------------------
	e.g. {1, 2, 3, 4} S/2= 5
	- We build the folliwng table by asking at each cell ([i][sum]) can we acheive such sum by either:
		including the number: num <= sum && [i-1][sum-num]? OR
		excluding it(were we able to achieve such sum with any of the previous sets): [i-1][sum]?

	i/sum           0   1   2   3   4   5
	--------------------------------------
	0: {}           T   F   F   F   F   F
	1: {1}          T   T   F   F   F   F
	2: {1,2}        T   T   T   T   F   F
	3: {1,2,3}      T   T   T   T   T   T
	4: {1,2,3,4}    T   T   T   T   T  |T| ---> our answer


	after building this table we can easily answer the question for {1,2,3,4} and sum 5:
	can it be done by inclusion?
		num <= sum && [i-1][sum-num] --> 4 <= 5 && [3][1] = True
	can it be done by exclusion?
		[i-1][sum] --> [3][5] = True

	so [4][5] = True OR True = True
*/

func EqualSumSubsets(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	// an odd sum can never yield two sets of equal sums
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	// setup set/sum matrix
	sumBySubset := make([][]bool, len(nums)+1)
	for i := range sumBySubset {
		sumBySubset[i] = make([]bool, sum+1)
	}
	// setup the first building block (an empty cell always yeilds a sum of zero)
	sumBySubset[0][0] = true

	for i := 1; i < len(sumBySubset); i++ {
		num := nums[i-1]
		for s := 0; s < sum+1; s++ {
			sumBySubset[i][s] = sumBySubset[i-1][s] || (num <= s && sumBySubset[i-1][s-num])
		}
	}

	return sumBySubset[len(nums)][sum]
}
