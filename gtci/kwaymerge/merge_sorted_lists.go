package kwaymerge

import "github.com/HassanElsherbini/algo-prep/common"

/* MEDIUM

SOLUTION: Time O(N * LOG(K)) | Space O(k)
- add first number of every list to a min heap
- while heap is not empty:
	- remove smallest number from min heap (the root) and add it to result list
	- if the corresponding list of the removed number is not empty:
		- take the next number from that list and add it to the min heap
- return result list
*/

func MergeSortedLists(lists [][]int) []int {
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
	minheap := common.NewBinaryHeap(comparator)

	for i, l := range lists {
		minheap.Insert(num{l[0], 0, i})
	}

	result := []int{}
	for minheap.Size() > 0 {
		element := minheap.Extract().(num)
		result = append(result, element.value)
		if element.index+1 < len(lists[element.listId]) {
			nextElement := num{
				lists[element.listId][element.index+1],
				element.index + 1,
				element.listId,
			}
			minheap.Insert(nextElement)
		}
	}

	return result
}
