package kwaymerge

import (
	"github.com/HassanElsherbini/algo-prep/common"
)

/* HARD

SOLUTION: Time O(K * LOG(N)) | Space O(N)
	- add first number of first min(k, N) lists to min heap
	- while k > 1: (removing k-1 smallest numbers)
		- remove smallest number (the root) from heap
		- from the array associated with the number removed from the heap,
			take the next number and add it to the heap.
		- decrement k

	- return the current head of the heap (the kth smallest)
*/

func FindKSmallestInSortedMatrix(matrix [][]int, k int) int {
	type num struct {
		value int
		row   int
		col   int
	}

	comparator := func(a, b interface{}) int {
		assertedA := a.(num)
		assertedB := b.(num)
		if assertedA.value < assertedB.value {
			return 1
		}
		if assertedA.value > assertedB.value {
			return -1
		}
		return 0
	}

	minHeap := common.NewBinaryHeap(comparator)
	for i := 0; i < common.Min(k, len(matrix)); i++ {
		minHeap.Insert(num{matrix[i][0], i, 0})
	}

	for i := k - 1; i > 0; i-- {
		removedNum := minHeap.Extract().(num)
		if removedNum.col+1 < len(matrix[removedNum.row]) {
			nextNum := num{matrix[removedNum.row][removedNum.col+1], removedNum.row, removedNum.col + 1}
			minHeap.Insert(nextNum)
		}
	}

	return minHeap.Extract().(num).value
}
