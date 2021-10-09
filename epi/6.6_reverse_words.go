package epi

// 6.6 - Reverse Words

/* SOLUTION: Time O(N) | Space O(N) strings are immutable in GO
- convert string to byte array
- reverse array
- reverse characters of each word
- return the string representation
*/
func ReversWords(str string) string {
	characters := []byte(str)
	reverse(characters[:])

	startOfWord := 0
	endOfWord := 0

	for endOfWord < len(characters) {
		for startOfWord < len(characters) && characters[startOfWord] == ' ' {
			startOfWord++
		}
		endOfWord = startOfWord
		for endOfWord < len(characters) && characters[endOfWord] != ' ' {
			endOfWord++
		}

		if startOfWord < endOfWord {
			reverse(characters[startOfWord:endOfWord])
		}

		startOfWord = endOfWord
	}

	return string(characters)
}

func reverse(characters []byte) {

	for left, right := 0, len(characters)-1; left < right; left, right = left+1, right-1 {
		characters[left], characters[right] = characters[right], characters[left]
	}
}
