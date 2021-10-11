package topk

import "github.com/HassanElsherbini/algo-prep/common/binaryheap"

/*EASY

SOLUTION: Time O(K * LOG(K) + N-K * LOG(K)) | Space O(K)
- add first k numbers to a max heap
- for each num in remaining numbers:
	- if num < biggest number in heap (the root):
		- remove max number in heap
		- add new number to heap

- return root of heap
*/
func FindKSmallestNumber(nums []int, k int) int {
	comparator := func(a, b interface{}) int {
		assertedA := a.(int)
		assertedb := b.(int)
		if assertedA > assertedb {
			return 1
		}
		if assertedA < assertedb {
			return -1
		}

		return 0
	}

	maxHeap := binaryheap.New(comparator)

	for i := 0; i < k; i++ {
		maxHeap.Insert(nums[i])
	}

	for i := k; i < len(nums); i++ {
		maxNum := maxHeap.Peek().(int)
		if nums[i] < maxNum {
			maxHeap.Extract()
			maxHeap.Insert(nums[i])
		}
	}
	return maxHeap.Peek().(int)
}
