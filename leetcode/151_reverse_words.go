package leetcode

import "strings"

/* MEDIUM
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.



Example 1:
Input: s = "the sky is blue"
Output: "blue is sky the"

Example 2:
Input: s = "  hello world  "
Output: "world hello"
Explanation: Your reversed string should not contain leading or trailing spaces.


Solution 1: Time O(N) | Space O(N)
- trim leading preceeding and leading spaces
- parse string for words and place them in an array
- concatenate the words in reverse order to form the reversed string
- return the reversed string


test: " the sky is blue "
	"the sky is blue" 		-  remove trailing spaces
	[the, sky, is, blue] 	-	 parse str for words
	"blue is sky the" 		-	 reverse parsed words

*/
func ReverseWords(str string) string {
	words := strings.Fields(str)
	var result strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])
		if i > 0 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

/*Solution 2: Time O(N) | Space O(1) if string is mutable!
- reverse string
- reverse words
- remove trailing and extra spaces between words
*/
