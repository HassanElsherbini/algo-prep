package modifiedbinarysearch

/* MEDIUM

SOLUTION: Time O(LOG(N)) | Space O(1)
- if key >= last letter:
	- return first letter

	- while start <= end:
		- if key < mid letter:
			- end = mid-1
		- else start = mid+1
	- return letter at start index
*/

func SearchNextLetter(letters []rune, key rune) rune {
	n := len(letters)
	if key >= letters[n-1] {
		return letters[0]
	}

	var start int
	end := n - 1

	for start <= end {
		mid := start + (end-start)/2
		if key < letters[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return letters[start]
}
