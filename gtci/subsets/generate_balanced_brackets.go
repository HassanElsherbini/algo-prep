package subsets

/* HARD

SOLUTION: Time(N * 2^2N) | Space(N * 2^N)
	- if no "(" and ")" left to generate:
		- add permutation to result
		- return result

	-if # of remaining "(" to generate is bigger than 0:
		- add "(" to permutation
		- decrement # of remaining "("
		- recurse passing new permutation, # of remaining "(", # of remaining ")" and result array

	if # of remaining ")" to generate is greater than # of remaining "(" :
		- add ")" to permutation
		- decrement # of remaining ")"
		- recurse passing new permutation, # of remaining "(", # of remaining ")" and result array

	- return result
*/

func GenerateBalancedParanthesis(n int) []string {
	return generateBalancedParanthesis(n, n, "", []string{})
}

func generateBalancedParanthesis(remainingOpen, remainingClosed int, permutation string, result []string) []string {
	if remainingOpen == 0 && remainingClosed == 0 {
		result = append(result, permutation)
		return result
	}

	if remainingOpen > 0 {
		result = generateBalancedParanthesis(remainingOpen-1, remainingClosed, permutation+"(", result)
	}

	if remainingClosed > remainingOpen {
		result = generateBalancedParanthesis(remainingOpen, remainingClosed-1, permutation+")", result)
	}

	return result
}
