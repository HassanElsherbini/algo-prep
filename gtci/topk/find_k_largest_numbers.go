package topk

import "github.com/HassanElsherbini/algo-prep/common/binaryheap"

/* EASY

SOLUTION: Time O(K * LOG(K) + N-K * LOG(K)) | Space O(K)
- insert the first K numbers in a min heap
- for each num in the remaining numbers:
	- if num is bigger than the smallest number in heap (the root):
		- remove smallest num from heap
		- insert new number
- return all K numbers in the heap
*/

func FindKLargestNumbers(nums []int, k int) []int {
	comparator := func(a, b interface{}) int {
		assertedA := a.(int)
		assertedb := b.(int)
		if assertedA < assertedb {
			return 1
		}
		if assertedA > assertedb {
			return -1
		}

		return 0
	}

	minHeap := binaryheap.New(comparator)
	for i := 0; i < k; i++ {
		minHeap.Insert(nums[i])
	}

	for i := k; i < len(nums); i++ {
		num := nums[i]
		smallest := minHeap.Peek().(int)
		if num > smallest {
			minHeap.Extract()
			minHeap.Insert(num)
		}
	}

	result := []int{}
	for _, item := range minHeap.Items {
		result = append(result, item.(int))
	}

	return result
}
