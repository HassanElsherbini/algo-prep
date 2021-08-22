package leetcode

/* EASY

Write a function that takes an unsigned integer and returns the number of '1' bits it has
(also known as the Hamming weight).

*The input must be a binary string of length 32.

Example 1:

Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

Example 2:

Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.


SOLUTION: Time O(Log(N)) | Space O(1) - N = input value
- to count the number of 1 bits, you need to check every bit to see if it's set to 1
	i.e 2^0, 2^1, 2^2, 2^3 etc..
	this can be done by shifting each bit of the number to the left by 1 (dividing by 2) till it reaches zero
	and incrementing the counter by 1 if the firtst bit is '1' at each shift.

- initilize a counter to 0
- for number > 0:
	- if first bit of number is 1:
		- increment counter
	- shift the number to the left by 1 (essentialy dividing it by 2)
- return counter
*/

func NumberOfOneBits(num uint32) int {
	var counter int
	for i := num; i > 0; i = i >> 1 {
		if (i & 1) == 1 {
			counter++
		}
	}
	return counter
}
