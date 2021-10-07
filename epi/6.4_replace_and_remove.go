package epi

// 6.4 - Replace and Remove

/* SOLUTION: Time O(N) | Space O(1)
- remove all "remove characters" by overwriting them with subsequent characters
- set a write index equal to size of expected result (index of last overwriten character + # of "replace characters" found)
- starting from index of lat overwritten character till 0:
	- if character is "character to replace":
		- overwrite two characters starting at write index with the "replace character"
		- decrement write index by 2
	- else:
		- overwrite the character at write index with the current character
		- decrement write index by 1
*/

func ReplaceAndRemove(chars []rune, size int, charToRemove, charToReplace, replaceChar rune) []rune {
	// remove instances of remove character and count instances of replace character
	charsToReplaceCount := 0
	writeIndx := 0

	for i := 0; i < size; i++ {
		if chars[i] != charToRemove {
			chars[writeIndx] = chars[i]
			writeIndx++
		}
		if chars[i] == charToReplace {
			charsToReplaceCount++
		}
	}

	// replace
	curIndx := writeIndx - 1
	writeIndx = writeIndx + charsToReplaceCount - 1
	finalSize := writeIndx + 1

	for curIndx >= 0 {
		if chars[curIndx] == charToReplace {
			chars[writeIndx] = replaceChar
			writeIndx--
			chars[writeIndx] = replaceChar
			writeIndx--
		} else {
			chars[writeIndx] = chars[curIndx]
			writeIndx--
		}
		curIndx--
	}
	return chars[:finalSize]
}
