package topk

import "github.com/HassanElsherbini/algo-prep/common"

/* EASY

SOLUTION: Time O(N * LOG(N)) | Space O(N)
- put all rope lengths in a min heap
- for each rope length in heap:
	- remove smallest two rope lengths from heap
	- create new rope by summing the two lengths
	- add new length to result
	- add new length to heap
- return result
*/

func MinCostToConnectRopes(ropeLengths []int) int {
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

	minHeap := common.NewBinaryHeap(comparator)
	for _, length := range ropeLengths {
		minHeap.Insert(length)
	}

	var result int
	for minHeap.Size() > 1 {
		newRopeLength := minHeap.Extract().(int) + minHeap.Extract().(int)
		result += newRopeLength
		minHeap.Insert(newRopeLength)
	}
	return result
}
