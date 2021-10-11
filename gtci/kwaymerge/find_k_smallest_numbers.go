package kwaymerge

import (
	"github.com/HassanElsherbini/algo-prep/common/binaryheap"
)

/* MEDIUM

SOLUTION: Time O(K * LOG(M)) | Space O(M)
	- add first number of all sorted lists into a min heap
	- while k > 1: (removes k-1 smallest numbers)
		- remove smallest num in min heap (the root)
		- if assosiated list of removed number is not empty:
			- add next number from that list to min heap
		- k--
	return curent root of minheap (the kth smallest number)
*/
func FindKSmallestNumber(nums [][]int, k int) int {
	type num struct {
		value  int
		index  int
		listId int
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
	minheap := binaryheap.New(comparator)

	for i, l := range nums {
		minheap.Insert(num{l[0], 0, i})
	}

	for k > 1 {
		element := minheap.Extract().(num)
		if element.index+1 < len(nums[element.listId]) {
			nextElement := num{
				nums[element.listId][element.index+1],
				element.index + 1,
				element.listId,
			}
			minheap.Insert(nextElement)
		}
		k--
	}

	return minheap.Extract().(num).value
}
