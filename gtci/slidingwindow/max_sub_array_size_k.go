package slidingwindow

import "math"

/* EASY

SOLUTION:
	-create a sliding window starting at index 0
	-for every number:
		add number at end to current window sum
		if window end >= k
			update max
			subtract num at start of window from sum
			move start of window by 1
*/

func MaxSumSubArrayOfSizeK(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	var start, windowSum int
	max := math.MinInt64

	for end, num := range nums {
		windowSum += num
		if end >= k-1 {
			if windowSum > max {
				max = windowSum
			}
			windowSum -= nums[start]
			start++
		}
	}

	return max
}
