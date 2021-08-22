package topk

import (
	"github.com/HassanElsherbini/algo-prep/common"
)

/* MEDIUM

SOLUTION: Time O(N + N * LOG(K)) | Space O(N)
- add all the numbers to a hashmap: number -> frequency
- add all numbers in hashmap to a k-min-heap using the stored frequency as the priority.
- return all k numbers in the min heap
*/

func FindKFrequentNumbers(nums []int, k int) []int {
	numFrequency := make(map[int]int)
	uniqueNums := []int{} // maintains the insertion order as hashmaps don't.

	for _, num := range nums {
		if _, ok := numFrequency[num]; ok {
			numFrequency[num] += 1
		} else {
			numFrequency[num] = 1
			uniqueNums = append(uniqueNums, num)
		}
	}

	comparator := func(a, b interface{}) int {
		assertedA := a.(int)
		assertedb := b.(int)
		if numFrequency[assertedA] < numFrequency[assertedb] {
			return 1
		}
		if numFrequency[assertedA] > numFrequency[assertedb] {
			return -1
		}
		return 0
	}
	minHeap := common.NewBinaryHeap(comparator)

	for _, num := range uniqueNums {
		minHeap.Insert(num)
		if minHeap.Size() > k {
			minHeap.Extract()
		}
	}

	// returns most frequent numbers in sorted order
	result := make([]int, k)
	for i := k; i > 0; i-- {
		result[i-1] = minHeap.Extract().(int)
	}
	return result
}
