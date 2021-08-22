package slidingwindow

/* MEDIUM

SOLUTION:
	for each letter:
	store letter in hash or if it exists increment its count
	if hash size is bigger than k
		move through visited letters reducing their counts in the hash
		with the goal of restoring the hashmap with k distinct letters
	update length of max substring found so far
	return result
*/

func LongestSubStringWithKDistinct(str string, k int) int {
	if len(str) < k {
		return len(str)
	}
	var start, max int
	visited := make(map[rune]int)

	for end, char := range str {
		visited[char] += 1

		if len(visited) > k {
			for len(visited) > k {
				c := rune(str[start])
				visited[c] -= 1
				if visited[c] == 0 {
					delete(visited, c)
				}
				start++
			}
		}

		if (end - start + 1) > max {
			max = end - start + 1
		}
	}

	return max
}
