package leetcode

import "github.com/HassanElsherbini/algo-prep/common"

/* MEDIUM

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.


SOLUTION: Time O(N) | Space O(N)
	- create a sliding window that represents the evaluated substring
	- create hashtable to store indecies of currently evaluated substring
	- start = 0
	- for end = 0 < length of string:
	- if letter is not unique and existing index of letter is within range of the sliding window:
			- set start = existing index + 1
			- update its index value in hashtable
	- else:
		- store it in hashtable as key with its index as value or update
		- update longest length if sliding window length (end - start)+1 > longest so far

*/

func LongestSubstringWithOutRepeatingChars(str string) int {

	result := 0
	lastOccurenceOfChars := make(map[byte]int)

	for start, end := 0, 0; end < len(str); end++ {
		char := str[end]
		lastIndx, exists := lastOccurenceOfChars[char]

		if exists && lastIndx >= start {
			start = lastIndx + 1
		} else {
			result = common.Max(end-start+1, result)
		}

		lastOccurenceOfChars[char] = end
	}
	return result
}
