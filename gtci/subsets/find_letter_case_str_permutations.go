package subsets

import "unicode"

/* MEDIUM

SOLUTION BFS: Time O(N * 2^N) | Space O(N * 2^N)
	- create a result array with input string as our first permutation
	- for i=0 till end of input str:
		- create an array(newPerms) to hold all new permuations created at this iteration
		- for each permutation in result array:
			- if char at i is a letter:
				- create a new permutation by capitalizing the letter
				- add permutation to newPerms
		- add all new permutations in newPerms to result array

	return result


SOLUTION DFS: Time O(N * 2^N) | Space O(N * 2^N) + O(N - stack size at maximum depth)
	- add permuation to result
	for each char in permutation starting from startIndx:
		- if char is a letter:
			- create new permutation by capitalizing the letter
			- pass new permutation and startIndx+1 to recursive call
	return result


*/

// BFS approach
func FindLetterCaseStringPermutations(str string) []string {
	result := []string{str}
	for i, _ := range str {
		var newPerms []string
		for _, perm := range result {
			if unicode.IsLetter(rune(perm[i])) {
				newPerms = append(newPerms, changeLetterCaseAtPosition(perm, i))
			}
		}
		result = append(result, newPerms...)
	}
	return result
}

func changeLetterCaseAtPosition(str string, pos int) string {
	if pos > len(str)-1 {
		return str
	}

	char := rune(str[pos])
	if !unicode.IsLetter(char) {
		return str
	}

	if unicode.IsLower(char) {
		char = unicode.ToUpper(char)
	} else {
		char = unicode.ToLower(char)
	}

	return str[:pos] + string(char) + str[pos+1:]
}

// DFS approach
func findLetterCaseStringPermutations(str string, startIndx int, result []string) []string {
	result = append(result, str)
	for i := startIndx; i < len(str); i++ {
		char := rune(str[i])
		if unicode.IsLetter(char) {
			if unicode.IsLower(char) {
				char = unicode.ToUpper(char)
			} else {
				char = unicode.ToLower(char)
			}
			newPermutation := str[:i] + string(char) + str[i+1:]
			result = findLetterCaseStringPermutations(newPermutation, i+1, result)
		}
	}
	return result
}
